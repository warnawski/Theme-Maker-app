package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/warnawski/theme-maker/src/components"
)

func MainWindow() {

	app := app.New()
	app.Settings().SetTheme(theme.DarkTheme())
	window := app.NewWindow("Theme-Maker")

	window.SetContent(withThemeBackground(components.NewTabs(app)))

	window.Resize(fyne.NewSize(550, 700))
	window.SetFixedSize(true)
	window.ShowAndRun()
}

func withThemeBackground(content fyne.CanvasObject) fyne.CanvasObject {

	bg := canvas.NewRectangle(theme.Current().Color(theme.ColorNameBackground, theme.VariantLight))
	return container.NewMax(bg, content)
}
