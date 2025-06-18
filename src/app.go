package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/warnawski/theme-maker/src/components"
)

func StartApp() {

	app := app.New()
	app.Settings().SetTheme(theme.DarkTheme())
	window := app.NewWindow("Theme-Maker")

	window.SetContent(withThemeBackground(components.NewTabs(app), app))

	window.Resize(fyne.NewSize(550, 700))
	window.SetFixedSize(true)
	window.ShowAndRun()
}

func withThemeBackground(content fyne.CanvasObject, a fyne.App) fyne.CanvasObject {
	bg := canvas.NewRectangle(
		theme.Current().Color(theme.ColorNameBackground, a.Settings().ThemeVariant()),
	)

	themeChanges := make(chan fyne.Settings)
	a.Settings().AddChangeListener(themeChanges)

	go func() {
		for range themeChanges {
			newVariant := a.Settings().ThemeVariant()
			newColor := theme.Current().Color(theme.ColorNameBackground, newVariant)

			fyne.Do(func() {
				bg.FillColor = newColor
				bg.Refresh()
			})
		}
	}()

	return container.NewMax(bg, content)
}
