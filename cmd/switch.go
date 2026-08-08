package cmd

import (
	"fmt"
	"os"

	"ghx/internal/config"
	"ghx/internal/ghcli"
	"ghx/internal/i18n"
	"ghx/internal/ui"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch <account>",
	Short: i18n.T().SwitchShort,
	Long:  i18n.T().SwitchLong,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]

		cfg, err := config.LoadConfig()
		if err != nil {
			ui.PrintError(i18n.T().ErrConfigLoad)
			os.Exit(1)
		}

		acc, ok := cfg.Accounts[username]
		if !ok {
			ui.PrintError(fmt.Sprintf(i18n.T().SwitchAccountNotFound, username))
			os.Exit(1)
		}

		if err := ghcli.LoginWithToken(acc.Token, acc.Host); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		if err := ghcli.SetGitUser(acc.Name, acc.Email); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		cfg.CurrentAccount = username
		if err := config.SaveConfig(cfg); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		ui.PrintSuccess(fmt.Sprintf(i18n.T().SwitchSuccess, username, acc.Name))
	},
}
