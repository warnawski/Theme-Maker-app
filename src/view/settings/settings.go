package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/warnawski/theme-maker/src/view/settings/colortheme"
)

type ThemeSwitcher interface {
	SetTheme(name string)
}

type ThemeManager struct {
	App   fyne.App
	Cache map[string]fyne.Theme
}

func NewThemeManager(app fyne.App) *ThemeManager {
	return &ThemeManager{
		App: app,
		Cache: map[string]fyne.Theme{
			"Dark":           theme.DarkTheme(),
			"Light":          theme.LightTheme(),
			"Clean-White":    colortheme.NewCleanWhiteTheme(),
			"Retro":          colortheme.NewRetroTheme(),
			"Dracula":        colortheme.NewDraculaTheme(),
			"Monokai":        colortheme.NewMonokaiTheme(),
			"Solarized-Dark": colortheme.NewSolarizedDarkTheme(),
		},
	}
}

func (tm ThemeManager) SetTheme(name string) {
	theme, ok := tm.Cache[name]
	if ok {
		tm.App.Settings().SetTheme(theme)
	}
}

func SettingsView(switcher ThemeSwitcher) fyne.CanvasObject {

	view := container.NewVBox(
		widget.NewLabel("App theme switcher"),
		widget.NewSelect([]string{"Dark", "Light", "Clean-White", "Retro", "Dracula", "Monokai", "Solarized-Dark"}, func(selected string) {
			switcher.SetTheme(selected)
		}),
	)

	return view
}
