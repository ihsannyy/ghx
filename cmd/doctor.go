package cmd

import (
	"fmt"

	"ghx/internal/config"
	"ghx/internal/ghcli"
	"ghx/internal/i18n"
	"ghx/internal/ui"

	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: i18n.T().DoctorShort,
	Long:  i18n.T().DoctorLong,
	Run: func(cmd *cobra.Command, args []string) {
		m := i18n.T()
		ui.PrintInfo(m.DoctorHeader)
		fmt.Println()

		mismatches := 0

		if ghcli.IsGHInstalled() {
			ui.PrintItemPass(m.DoctorGHInstalled)
		} else {
			ui.PrintItemFail(m.DoctorGHNotFound)
			mismatches++
		}

		if ghcli.IsGitInstalled() {
			ui.PrintItemPass(m.DoctorGitInstalled)
		} else {
			ui.PrintItemFail(m.DoctorGitNotFound)
			mismatches++
		}

		cfg, err := config.LoadConfig()
		if err != nil || cfg == nil {
			ui.PrintItemFail(m.DoctorConfigError)
			mismatches++
			fmt.Println()
			ui.PrintError(fmt.Sprintf(m.DoctorIssues, mismatches))
			return
		}
		ui.PrintItemPass(m.DoctorConfigValid)

		if cfg.CurrentAccount == "" {
			ui.PrintItemWarn(m.DoctorNoActive)
			mismatches++
			fmt.Println()
			ui.PrintError(fmt.Sprintf(m.DoctorIssues, mismatches))
			return
		}

		activeAcc, exists := cfg.Accounts[cfg.CurrentAccount]
		if !exists {
			ui.PrintItemFail(fmt.Sprintf("Active account '%s' not found in configuration", cfg.CurrentAccount))
			mismatches++
			fmt.Println()
			ui.PrintError(fmt.Sprintf(m.DoctorIssues, mismatches))
			return
		}

		ui.PrintItemPass(fmt.Sprintf(m.DoctorActiveAccount, activeAcc.Username))

		ghUser, err := ghcli.GetActiveGHUsername(activeAcc.Host)
		if err != nil || ghUser == "" {
			ui.PrintItemWarn(fmt.Sprintf(m.DoctorGHMismatch, "none", activeAcc.Username, activeAcc.Username))
			mismatches++
		} else if ghUser == activeAcc.Username {
			ui.PrintItemPass(fmt.Sprintf(m.DoctorGHMatch, ghUser))
		} else {
			ui.PrintItemWarn(fmt.Sprintf(m.DoctorGHMismatch, ghUser, activeAcc.Username, activeAcc.Username))
			mismatches++
		}

		gitName, gitEmail, _ := ghcli.GetGitUser()
		if gitName == activeAcc.Name && gitEmail == activeAcc.Email {
			ui.PrintItemPass(fmt.Sprintf(m.DoctorGitMatch, gitName, gitEmail))
		} else {
			ui.PrintItemWarn(fmt.Sprintf(m.DoctorGitMismatch, gitName, gitEmail, activeAcc.Name, activeAcc.Email, activeAcc.Username))
			mismatches++
		}

		sshUser, err := ghcli.TestSSHConnection(activeAcc.Host)
		if err != nil || sshUser == "" {
			ui.PrintItemInfo(m.DoctorSSHNotConfig)
		} else if sshUser == activeAcc.Username {
			ui.PrintItemPass(fmt.Sprintf(m.DoctorSSHMatch, sshUser))
		} else {
			ui.PrintItemWarn(fmt.Sprintf(m.DoctorSSHMismatch, sshUser, activeAcc.Username))
			mismatches++
		}

		fmt.Println()
		if mismatches == 0 {
			ui.PrintSuccess(m.DoctorPassed)
		} else {
			ui.PrintError(fmt.Sprintf(m.DoctorIssues, mismatches))
		}
	},
}
