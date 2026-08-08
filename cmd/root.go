package cmd

import (
	"os"

	"ghx/internal/config"
	"ghx/internal/i18n"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ghx",
	Short: i18n.T().RootShort,
	Long:  i18n.T().RootLong,
}

func Execute() {
	cfg, err := config.LoadConfig()
	if err == nil && cfg != nil && cfg.Language != "" {
		i18n.SetLanguage(cfg.Language)
	}

	updateCommandTexts()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func updateCommandTexts() {
	m := i18n.T()
	rootCmd.Short = m.RootShort
	rootCmd.Long = m.RootLong

	for _, c := range rootCmd.Commands() {
		switch c.Name() {
		case "login":
			c.Short = m.LoginShort
			c.Long = m.LoginLong
		case "add":
			c.Short = m.AddShort
			c.Long = m.AddLong
		case "switch":
			c.Short = m.SwitchShort
			c.Long = m.SwitchLong
		case "list":
			c.Short = m.ListShort
			c.Long = m.ListLong
		case "remove":
			c.Short = m.RemoveShort
			c.Long = m.RemoveLong
		case "email":
			c.Short = m.EmailShort
			c.Long = m.EmailLong
		case "current":
			c.Short = m.CurrentShort
			c.Long = m.CurrentLong
		case "lang":
			c.Short = m.LangShort
			c.Long = m.LangLong
		case "doctor":
			c.Short = m.DoctorShort
			c.Long = m.DoctorLong
		}
	}
}

func init() {
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(switchCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(emailCmd)
	rootCmd.AddCommand(currentCmd)
	rootCmd.AddCommand(langCmd)
	rootCmd.AddCommand(doctorCmd)
}
