package icon

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	custom "github.com/warnawski/theme-maker/src/components/button"
	ht "github.com/warnawski/theme-maker/src/components/layout"
)

func NewIconView() fyne.CanvasObject {

	tabs := container.NewHBox(
		ht.ButtonMargin(80, 20),
		custom.NewCButton("IDE VSC", func() {}),
		ht.ButtonGap(20, 20),
		custom.NewCButton("IDE JetBrains", func() {}),
		ht.ButtonMargin(80, 20),
	)

	view := container.NewVBox(
		ht.TopMargin(50),
		tabs,
		layout.NewSpacer(),
	)

	return view
}
