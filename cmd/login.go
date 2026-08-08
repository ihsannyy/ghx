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
	loginToken    string
	loginHostname string
	loginName     string
	loginEmail    string
)

var loginCmd = &cobra.Command{
	Use:   "login [token]",
	Short: i18n.T().LoginShort,
	Long:  i18n.T().LoginLong,
	Run: func(cmd *cobra.Command, args []string) {
		if !ghcli.IsGHInstalled() {
			ui.PrintError(i18n.T().ErrGHNotInstalled)
			os.Exit(1)
		}
		if !ghcli.IsGitInstalled() {
			ui.PrintError(i18n.T().ErrGitNotInstalled)
			os.Exit(1)
		}

		token := loginToken
		if token == "" && len(args) > 0 {
			token = args[0]
		}

		if loginHostname == "" {
			loginHostname = "github.com"
		}

		if token == "" {
			activeToken, err := ghcli.GetActiveGHToken(loginHostname)
			if err == nil && activeToken != "" {
				token = activeToken
			}
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

		username, fetchedName, fetchedEmail, err := ghcli.FetchUserInfo(token, loginHostname)
		if err != nil {
			ui.PrintError(fmt.Sprintf("%s (%v)", i18n.T().LoginErrorFetch, err))
			os.Exit(1)
		}

		name := fetchedName
		if loginName != "" {
			name = loginName
		}
		email := fetchedEmail
		if loginEmail != "" {
			email = loginEmail
		}

		if err := ghcli.LoginWithToken(token, loginHostname); err != nil {
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
			Host:     loginHostname,
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
	loginCmd.Flags().StringVarP(&loginToken, "token", "t", "", "GitHub Personal Access Token")
	loginCmd.Flags().StringVar(&loginHostname, "hostname", "github.com", "GitHub Hostname")
	loginCmd.Flags().StringVar(&loginName, "name", "", "Custom git user.name")
	loginCmd.Flags().StringVar(&loginEmail, "email", "", "Custom git user.email")
}
