package cmd

import (
	"fmt"
	"os"

	"ghx/internal/config"
	"ghx/internal/i18n"
	"ghx/internal/ui"

	"github.com/spf13/cobra"
)

var currentCmd = &cobra.Command{
	Use:   "current",
	Short: i18n.T().CurrentShort,
	Long:  i18n.T().CurrentLong,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			ui.PrintError(i18n.T().ErrConfigLoad)
			os.Exit(1)
		}

		if cfg.CurrentAccount == "" {
			ui.PrintInfo(i18n.T().CurrentNoActive)
			return
		}

		acc, ok := cfg.Accounts[cfg.CurrentAccount]
		if !ok {
			ui.PrintInfo(i18n.T().CurrentNoActive)
			return
		}

		ui.PrintInfo(fmt.Sprintf(i18n.T().CurrentActive, acc.Username, acc.Name, acc.Email))
	},
}
