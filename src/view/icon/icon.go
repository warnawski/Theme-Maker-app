package icon

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/warnawski/theme-maker/src/components/custom"
	"github.com/warnawski/theme-maker/src/components/indents"
)

func IconView() fyne.CanvasObject {

	view := container.NewHBox(
		indents.TopMargin(50),
		chooseButtons(),
	)

	return view
}

func chooseButtons() fyne.CanvasObject {

	btn1 := custom.NewCButton("VSC", func() {})

	btn2 := custom.NewCButton("JetBrains", func() {})

	btns := container.NewHBox(
		btn1,
		indents.ButtonGap(10, 20),
		btn2,
	)

	return btns
}
