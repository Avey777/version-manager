package vcli

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/gvcgo/version-manager/internal/download"
	"github.com/gvcgo/version-manager/internal/installer"
	"github.com/spf13/cobra"
)

/*
Uninstalls a version for an SDK.
*/
var UninstallVersionCmd = &cobra.Command{
	Use:     "uninstall",
	Aliases: []string{"uni", "r"},
	GroupID: GroupID,
	Short:   "Uninstall versions for an SDK.",
	Long:    "Example: vmr uninstall sdkname@version or sdkname@all.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			return
		}
		versionInfo := args[0]
		if !strings.Contains(versionInfo, "@") {
			cmd.Help()
			return
		}
		sList := strings.Split(versionInfo, "@")
		if len(sList) != 2 || sList[1] == "" {
			cmd.Help()
			return
		}
		sdkName := sList[0]
		version := sList[1]

		if version == "all" {
			lif := installer.NewIVFinder(sdkName)
			lif.UninstallAllVersions()
		} else {
			versionFilePath := download.GetVersionFilePath(sdkName)
			content, _ := os.ReadFile(versionFilePath)
			rawVersionList := make(download.VersionList)
			json.Unmarshal(content, &rawVersionList)
			installerType := "unarchiver"
			for _, vl := range rawVersionList {
				if len(vl) > 0 {
					installerType = vl[0].Installer
					break
				}
			}
			ins := installer.NewInstaller(sdkName, version, "", download.Item{Installer: installerType})
			ins.Uninstall()
		}
	},
}
