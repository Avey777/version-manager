package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
	"github.com/gvcgo/goutils/pkgs/gutils"
	"github.com/gvcgo/version-manager/pkgs/conf"
	"github.com/gvcgo/version-manager/pkgs/envs"
	"github.com/gvcgo/version-manager/pkgs/installer"
	"github.com/gvcgo/version-manager/pkgs/register"
	"github.com/spf13/cobra"
)

const (
	GroupID string = "vm"
)

/*
CLIs
*/
type Cli struct {
	rootCmd *cobra.Command
	groupID string
}

func New() (c *Cli) {
	c = &Cli{
		rootCmd: &cobra.Command{
			Short: "version manager",
			Long:  "vm <Command> <SubCommand> --flags args...",
		},
		groupID: GroupID,
	}
	c.rootCmd.AddGroup(&cobra.Group{ID: c.groupID, Title: "Command list: "})
	c.initiate()
	return
}

func (c *Cli) initiate() {
	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "search",
		Aliases: []string{"s"},
		GroupID: GroupID,
		Short:   "Shows the available versions of an application.",
		Long:    "Example: vm search go.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			sch := installer.NewSearcher()
			sch.Search(args[0])
		},
	})

	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "show",
		Aliases: []string{"sh", "S"},
		GroupID: GroupID,
		Short:   "Shows the supported applications.",
		Run: func(cmd *cobra.Command, args []string) {
			register.ShowAppList()
		},
	})

	useCmd := &cobra.Command{
		Use:     "use",
		Aliases: []string{"u"},
		GroupID: GroupID,
		Short:   "Installs and switches to specified version.",
		Long:    "Example: vm use go@1.22.1",
		Run: func(cmd *cobra.Command, args []string) {
			mirrorInChina, _ := cmd.Flags().GetBool("mirror_in_china")
			rds, _ := cmd.Flags().GetBool("rustup-default-stable")
			if rds {
				// only for rustup default.
				if mirrorInChina {
					os.Setenv("RUSTUP_DIST_SERVER", "https://mirrors.ustc.edu.cn/rust-static")
					os.Setenv("RUSTUP_UPDATE_ROOT", "https://mirrors.ustc.edu.cn/rust-static/rustup")
				}
				gutils.ExecuteSysCommand(false, "",
					"rustup", "default", "stable")
				return
			}
			if len(args) == 0 || !strings.Contains(args[0], "@") {
				cmd.Help()
				return
			}
			sList := strings.Split(args[0], "@")
			if len(sList) != 2 {
				cmd.Help()
				return
			}
			threads, _ := cmd.Flags().GetInt("threads")
			os.Setenv(conf.VMDownloadThreadsEnvName, gconv.String(threads))

			if mirrorInChina {
				os.Setenv(conf.VMUseMirrorInChinaEnvName, "true")
			} else {
				os.Setenv(conf.VMUseMirrorInChinaEnvName, "false")
			}

			if ins, ok := register.VersionKeeper[sList[0]]; ok {
				ins.SetVersion(sList[1])
				register.RunInstaller(ins)
			} else {
				gprint.PrintError("Unsupported app: %s.", sList[0])
			}
		},
	}
	useCmd.Flags().IntP("threads", "t", 1, "Number of threads to use for downloading.")
	useCmd.Flags().BoolP("mirror_in_china", "c", false, "Downlowd from mirror sites in China.")
	useCmd.Flags().BoolP("rustup-default-stable", "r", false, "Set rustup default stable.")
	c.rootCmd.AddCommand(useCmd)

	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "uninstall",
		Aliases: []string{"U"},
		GroupID: GroupID,
		Short:   "Uninstalls a version or an app.",
		Long:    "Example: 1. vm U go@all; 2. vm U go@1.22.1",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 || !strings.Contains(args[0], "@") {
				cmd.Help()
				return
			}
			sList := strings.Split(args[0], "@")
			if len(sList) != 2 {
				cmd.Help()
				return
			}
			if ins, ok := register.VersionKeeper[sList[0]]; ok {
				ins.SetVersion(sList[1])
				register.RunUnInstaller(ins)
			} else {
				gprint.PrintError("Unsupported app: %s.", sList[0])
			}
		},
	})

	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "local",
		Aliases: []string{"l"},
		GroupID: GroupID,
		Short:   "Shows installed versions for an app.",
		Long:    "Example: vm l go.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			register.ShowInstalled(args[0])
		},
	})

	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "set-proxy",
		Aliases: []string{"sp"},
		GroupID: GroupID,
		Short:   "Sets proxy for version manager.",
		Long:    "Example: vm sp http://127.0.0.1:2023",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			pUrl := args[0]
			conf.SaveConfigFile(&conf.Config{ProxyURI: pUrl})
		},
	})

	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "set-reverse-proxy",
		Aliases: []string{"sr", "srp"},
		GroupID: GroupID,
		Short:   "Sets reverse proxy for version manager.",
		Long:    "Example: vm sr https://gvc.1710717.xyz/proxy/",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			rPxy := args[0]
			conf.SaveConfigFile(&conf.Config{ReverseProxy: rPxy})
		},
	})

	// c.rootCmd.AddCommand(&cobra.Command{
	// 	Use:     "set-app-dir",
	// 	Aliases: []string{"sd", "sad"},
	// 	GroupID: GroupID,
	// 	Short:   "Sets installation dir for apps.",
	// 	Long:    "Example: vm sd <where-to-install-apps>.",
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		if len(args) == 0 {
	// 			cmd.Help()
	// 			return
	// 		}
	// 		appDir := args[0]
	// 		conf.SaveConfigFile(&conf.Config{AppInstallationDir: appDir})
	// 	},
	// })

	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "install-self",
		Aliases: []string{"i", "is"},
		GroupID: GroupID,
		Short:   "Installs version manager.",
		Run: func(cmd *cobra.Command, args []string) {
			vmBinName := "vm"
			if runtime.GOOS == gutils.Windows {
				vmBinName = "vm.exe"
			}
			binPath := filepath.Join(conf.GetManagerDir(), vmBinName)
			if ok, _ := gutils.PathIsExist(binPath); ok {
				os.RemoveAll(binPath)
			}
			currentBinPath, _ := os.Executable()
			if strings.HasSuffix(currentBinPath, vmBinName) {
				gutils.CopyFile(currentBinPath, binPath)
			}
			em := envs.NewEnvManager()
			defer em.CloseKey()
			em.AddToPath(conf.GetManagerDir())

			if ok, _ := gutils.PathIsExist(conf.GetConfPath()); ok {
				return
			}
			// Sets app installation Dir.
			fmt.Println(gprint.CyanStr(`Enter App Installation Dir["$Home/.vm/" by default]:`))
			var appDir string
			fmt.Scanln(&appDir)
			if appDir == "" {
				appDir = conf.GetManagerDir()
			}
			conf.SaveConfigFile(&conf.Config{AppInstallationDir: appDir})
		},
	})

	c.rootCmd.AddCommand(&cobra.Command{
		Use:     "clear-cache",
		Aliases: []string{"c", "cc"},
		GroupID: GroupID,
		Short:   "Clear cached zip files for an app.",
		Long:    "Example: vm c go",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			appName := args[0]
			if ins, ok := register.VersionKeeper[appName]; ok {
				register.RunClearCache(ins)
			}
		},
	})
}

func (that *Cli) Run() {
	if err := that.rootCmd.Execute(); err != nil {
		gprint.PrintError("%+v", err)
	}
}
