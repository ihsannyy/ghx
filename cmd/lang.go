package cmd

import (
	"fmt"
	"os"

	"ghx/internal/config"
	"ghx/internal/i18n"
	"ghx/internal/ui"

	"github.com/spf13/cobra"
)

var langCmd = &cobra.Command{
	Use:   "lang [en|id]",
	Short: i18n.T().LangShort,
	Long:  i18n.T().LangLong,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil {
			ui.PrintError(i18n.T().ErrConfigLoad)
			os.Exit(1)
		}

		if len(args) == 0 {
			ui.PrintInfo(fmt.Sprintf(i18n.T().LangCurrent, cfg.Language))
			return
		}

		newLang := args[0]
		if newLang != "en" && newLang != "id" {
			ui.PrintError(fmt.Sprintf(i18n.T().LangUnsupported, newLang))
			os.Exit(1)
		}

		cfg.Language = newLang
		if err := config.SaveConfig(cfg); err != nil {
			ui.PrintError(err.Error())
			os.Exit(1)
		}

		i18n.SetLanguage(newLang)
		updateCommandTexts()

		ui.PrintSuccess(fmt.Sprintf(i18n.T().LangUpdated, newLang))
	},
}
