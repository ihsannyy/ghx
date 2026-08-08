package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ghx/internal/config"
	"ghx/internal/ghcli"
	"ghx/internal/i18n"
	"ghx/internal/ui"

	"github.com/spf13/cobra"
)

var (
	addHostname string
	addName     string
	addEmail    string
)

var addCmd = &cobra.Command{
	Use:   "add <token>",
	Short: i18n.T().AddShort,
	Long:  i18n.T().AddLong,
	Run: func(cmd *cobra.Command, args []string) {
		if !ghcli.IsGHInstalled() {
			ui.PrintError(i18n.T().ErrGHNotInstalled)
			os.Exit(1)
		}
		if !ghcli.IsGitInstalled() {
			ui.PrintError(i18n.T().ErrGitNotInstalled)
			os.Exit(1)
		}

		var token string
		if len(args) > 0 {
			token = args[0]
		}

		if addHostname == "" {
			addHostname = "github.com"
		}

		if token == "" {
			fmt.Print(i18n.T().LoginPrompt)
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				token = strings.TrimSpace(scanner.Text())
			}
		}

		if token == "" {
			ui.PrintError(i18n.T().LoginErrorFetch)
			os.Exit(1)
		}

		username, fetchedName, fetchedEmail, err := ghcli.FetchUserInfo(token, addHostname)
		if err != nil {
			ui.PrintError(fmt.Sprintf("%s (%v)", i18n.T().LoginErrorFetch, err))
			os.Exit(1)
		}

		name := fetchedName
		if addName != "" {
			name = addName
		}
		email := fetchedEmail
		if addEmail != "" {
			email = addEmail
		}

		if err := ghcli.LoginWithToken(token, addHostname); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		if err := ghcli.SetGitUser(name, email); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		cfg, err := config.LoadConfig()
		if err != nil {
			ui.PrintError(i18n.T().ErrConfigLoad)
			os.Exit(1)
		}

		cfg.Accounts[username] = config.Account{
			Username: username,
			Name:     name,
			Email:    email,
			Token:    token,
			Host:     addHostname,
		}
		cfg.CurrentAccount = username

		if err := config.SaveConfig(cfg); err != nil {
			ui.PrintError(i18n.T().LoginErrorSave)
			os.Exit(1)
		}

		ui.PrintSuccess(fmt.Sprintf(i18n.T().LoginSuccess, username, name))
	},
}

func init() {
	addCmd.Flags().StringVar(&addHostname, "hostname", "github.com", "GitHub Hostname")
	addCmd.Flags().StringVar(&addName, "name", "", "Custom git user.name")
	addCmd.Flags().StringVar(&addEmail, "email", "", "Custom git user.email")
}
