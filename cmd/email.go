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

var emailCmd = &cobra.Command{
	Use:   "email [email]",
	Short: i18n.T().EmailShort,
	Long:  i18n.T().EmailLong,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			ui.PrintError(i18n.T().ErrConfigLoad)
			os.Exit(1)
		}

		currentUsername := cfg.CurrentAccount

		if len(args) == 0 {
			_, gitEmail, _ := ghcli.GetGitUser()
			if currentUsername != "" {
				if acc, ok := cfg.Accounts[currentUsername]; ok {
					ui.PrintInfo(fmt.Sprintf(i18n.T().EmailCurrent, acc.Username, acc.Email))
					return
				}
			}
			ui.PrintInfo(fmt.Sprintf("Git global email: %s", gitEmail))
			return
		}

		newEmail := args[0]
		if err := ghcli.SetGitUser("", newEmail); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		if currentUsername != "" {
			if acc, ok := cfg.Accounts[currentUsername]; ok {
				acc.Email = newEmail
				cfg.Accounts[currentUsername] = acc
				_ = config.SaveConfig(cfg)
			}
		}

		ui.PrintSuccess(fmt.Sprintf(i18n.T().EmailUpdated, currentUsername, newEmail))
	},
}
