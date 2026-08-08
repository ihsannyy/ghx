package cmd

import (
	"os"
	"sort"

	"ghx/internal/config"
	"ghx/internal/i18n"
	"ghx/internal/ui"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: i18n.T().ListShort,
	Long:  i18n.T().ListLong,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			ui.PrintError(i18n.T().ErrConfigLoad)
			os.Exit(1)
		}

		if len(cfg.Accounts) == 0 {
			ui.PrintInfo(i18n.T().ListNoAccounts)
			return
		}

		headers := []string{
			i18n.T().ListHeaderStatus,
			i18n.T().ListHeaderUsername,
			i18n.T().ListHeaderName,
			i18n.T().ListHeaderEmail,
		}

		var keys []string
		for u := range cfg.Accounts {
			keys = append(keys, u)
		}
		sort.Strings(keys)

		var rows []ui.TableRow
		for _, u := range keys {
			acc := cfg.Accounts[u]
			activeMark := " "
			if u == cfg.CurrentAccount {
				activeMark = "*"
			}
			rows = append(rows, ui.TableRow{
				Active:   activeMark,
				Username: acc.Username,
				Name:     acc.Name,
				Email:    acc.Email,
			})
		}

		ui.PrintAccountTable(os.Stdout, headers, rows)
	},
}
