package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	view "github.com/warnawski/theme-maker/src/view/home"
	"github.com/warnawski/theme-maker/src/view/settings"
)

func NewTabs(a fyne.App) *container.AppTabs {

	tabs := container.NewAppTabs(

		container.NewTabItemWithIcon(
			"Home                           ",
			theme.HomeIcon(),
			view.HomeView()),

		container.NewTabItemWithIcon(
			"Work Area                                                            ",
			theme.WarningIcon(),
			widget.NewLabel("")),

		container.NewTabItemWithIcon(
			"Docs",
			theme.HelpIcon(),
			widget.NewLabel("")),

		container.NewTabItemWithIcon(
			"Settings",
			theme.SettingsIcon(),
			settings.SettingsView(a)),
	)

	return tabs

}
