package cmd

import (
	"fmt"
	"os"

	"ghx/internal/config"
	"ghx/internal/i18n"
	"ghx/internal/ui"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <account>",
	Short: i18n.T().RemoveShort,
	Long:  i18n.T().RemoveLong,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			ui.PrintError(i18n.T().ErrConfigLoad)
			os.Exit(1)
		}

		if _, ok := cfg.Accounts[username]; !ok {
			ui.PrintError(fmt.Sprintf(i18n.T().RemoveAccountNotFound, username))
			os.Exit(1)
		}

		delete(cfg.Accounts, username)
		if cfg.CurrentAccount == username {
			cfg.CurrentAccount = ""
		}

		if err := config.SaveConfig(cfg); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		ui.PrintSuccess(fmt.Sprintf(i18n.T().RemoveSuccess, username))
	},
}
