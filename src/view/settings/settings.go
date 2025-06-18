package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	t "github.com/warnawski/theme-maker/src/components/theme"
)

func SettingsView(a fyne.App) fyne.CanvasObject {

	view := container.NewVBox(
		widget.NewLabel("App theme switcher"),
		widget.NewSelect([]string{"Dark", "Retro"}, func(selected string) {
			themeSwitcher(a, selected)
		}),
	)

	return view
}

func themeSwitcher(a fyne.App, selected string) {

	switch selected {

	case "Dark":
		a.Settings().SetTheme(theme.DarkTheme())
	case "Retro":
		a.Settings().SetTheme(t.NewRetroTheme())
	}
}
