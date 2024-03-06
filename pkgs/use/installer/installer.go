package installer

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gvcgo/goutils/pkgs/archiver"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
	"github.com/gvcgo/goutils/pkgs/gutils"
	"github.com/gvcgo/goutils/pkgs/request"
	"github.com/gvcgo/version-manager/pkgs/conf"
	"github.com/gvcgo/version-manager/pkgs/search"
	"github.com/gvcgo/version-manager/pkgs/versions"
)

type Env struct {
	Name  string
	Value string
}

type Installer struct {
	AppName            string
	Version            string
	Searcher           *search.Searcher
	Fetcher            *request.Fetcher
	V                  *versions.VersionItem
	IsZipFile          bool
	BinaryRenameTo     string
	BinDirGetter       func(version string) [][]string
	BinListGetter      func() []string
	FlagFileGetter     func() []string
	EnvGetter          func(appName, version string) []Env
	StoreMultiVersions bool
}

func NewInstaller(appName, version string) (i *Installer) {
	i = &Installer{
		AppName:  appName,
		Version:  version,
		Searcher: search.NewSearcher(),
		Fetcher:  conf.GetFetcher(),
	}
	return
}

// Searches version files for an application.
func (i *Installer) SearchVersion() {
	if i.Searcher == nil {
		i.Searcher = search.NewSearcher()
	}
	vf := i.Searcher.GetVersions(i.AppName)
	vs := make([]string, 0)
	for key := range vf {
		if strings.Contains(key, i.Version) {
			vs = append(vs, key)
		}
	}

	if len(vs) == 0 {
		i.V = nil
		gprint.PrintError("Cannot find version: %s", i.Version)
	} else if len(vs) == 1 {
		i.Version = vs[0]
		i.V = &vf[i.Version][0]
	} else {
		i.V = nil
		gprint.PrintError("Found multiple versions: \n%v", strings.Join(vs, "\n"))
	}
}

func (i *Installer) Download() (zipFilePath string) {
	i.SearchVersion()
	if i.V == nil {
		return
	}
	zipDir := conf.GetZipFileDir(i.AppName)
	if ok, _ := gutils.PathIsExist(zipDir); !ok {
		if err := os.MkdirAll(zipDir, os.ModePerm); err != nil {
			gprint.PrintError("Failed to create directory: %s", zipDir)
			return
		}
	}

	i.Fetcher.SetUrl(conf.DecorateUrl(i.V.Url))
	zipFilePath = filepath.Join(zipDir, filepath.Base(i.V.Url))
	i.Fetcher.GetAndSaveFile(zipFilePath)

	// checksum
	if i.V.Sum != "" && i.V.SumType != "" {
		if ok := gutils.CheckSum(zipFilePath, i.V.SumType, i.V.Sum); !ok {
			zipFilePath = ""
			os.RemoveAll(zipFilePath) // checksum failed.
		}
	}
	return
}

func handleUnzipFailedError(zipFilePath string, err error) {
	gprint.PrintError("Failed to unzip file: %s, %+v", zipFilePath, err)
	os.RemoveAll(zipFilePath)
}

func (i *Installer) Unzip(zipFilePath string) {
	if i.IsZipFile {
		tempDir := conf.GetVMTempDir()
		if arch, err := archiver.NewArchiver(zipFilePath, tempDir, false); err == nil {
			_, err = arch.UnArchive()
			if err != nil {
				handleUnzipFailedError(zipFilePath, err)
				return
			}
		} else {
			handleUnzipFailedError(zipFilePath, err)
		}
	}
}

func (i *Installer) Copy() {
	// find directory to copy.
	if i.FlagFileGetter != nil {
		f := NewFinder(i.FlagFileGetter()...)
		f.Find(conf.GetVMTempDir())
		if f.Home != "" {
			err := gutils.CopyDirectory(f.Home, filepath.Join(conf.GetVMVersionsDir(i.AppName), i.Version), true)
			if err != nil {
				gprint.PrintError("Failed to copy directory: %s, %+v", f.Home, err)
			}
		}
	}
	os.RemoveAll(conf.GetVMTempDir())
}

func (i *Installer) CreateVersionSymbol() {
	versionPath := filepath.Join(conf.GetVMVersionsDir(i.AppName), i.Version)
	symbolPath := filepath.Join(conf.GetVMVersionsDir(i.AppName), i.AppName)

	if ok, _ := gutils.PathIsExist(versionPath); ok {
		// remove old symbol
		if ok, _ := gutils.PathIsExist(symbolPath); ok {
			os.RemoveAll(symbolPath)
		}
		// create symbolic
		os.Symlink(versionPath, symbolPath)
	}
}

func (i *Installer) CreateBinarySymbol() {
	if i.BinDirGetter != nil {
		symbolPath := filepath.Join(conf.GetVMVersionsDir(i.AppName), i.AppName)
		if ok, _ := gutils.PathIsExist(symbolPath); ok {
			for _, bDir := range i.BinDirGetter(i.Version) {
				d := filepath.Join(symbolPath, filepath.Join(bDir...))
				if dList, err := os.ReadDir(d); err == nil {
					for _, dd := range dList {
						if !dd.IsDir() {
							fPath := filepath.Join(d, dd.Name())
							symPath := filepath.Join(conf.GetAppBinDir(), dd.Name())
							// if ok, _ := gutils.PathIsExist(symPath); ok {
							// 	os.RemoveAll(symPath)
							// }
							os.Symlink(fPath, symPath)
						}
					}
				}
			}
		}
	}
}

func (i *Installer) SaveLinksInfo() {

}
