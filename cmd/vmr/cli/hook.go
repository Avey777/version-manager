package cli

import (
	"github.com/gvcgo/version-manager/cmd/vmr/cli/vcli"
	"github.com/gvcgo/version-manager/internal/installer"
	"github.com/spf13/cobra"
)

/*
subcommand for cd hook.
*/
var useHookCmd = &cobra.Command{
	Use:     "use",
	Aliases: []string{"u", "h"},
	GroupID: vcli.GroupID,
	Short:   "This is the hook for command cd.",
	Long:    "Users do not use this command directly.",
	Run: func(cmd *cobra.Command, args []string) {
		// enable locked version.
		if ok, _ := cmd.Flags().GetBool("enable-locked-version"); ok {
			l := installer.NewVLocker()
			l.HookForCdCommand()
		}

		// use a version.

		// lock a version.

		// use a version only for current session.
	},
}

func init() {
	useHookCmd.Flags().BoolP("enable-locked-version", "E", false, "To enable the locked version for current project.")
}
