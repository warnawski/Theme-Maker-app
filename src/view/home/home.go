package home

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	loader "github.com/warnawski/theme-maker/internal/loader"
)

func HomeView() fyne.CanvasObject {

	logo := loader.LoadFyneCanvasImage("resource/assets/theme-icon.png")
	logo.SetMinSize(fyne.NewSize(350, 350))

	view := container.NewVBox(
		container.NewCenter(logo),
		ListButton(),
		TopMargin(30),
		NewProjectButtons(),
	)

	return view
}

func ListButton() fyne.CanvasObject {

	btn1 := widget.NewButton("●  CCHHLENNNN...", func() {})
	btn1.Importance = widget.LowImportance

	btn2 := widget.NewButton("●  CCHHLENNNN...", func() {})
	btn2.Importance = widget.LowImportance

	btn3 := widget.NewButton("⠅  more", func() {})
	btn3.Importance = widget.LowImportance

	res := container.NewHBox(
		ButtonMargin(60, 20),
		btn1,
		ButtonGap(15, 20),
		btn2,
		ButtonGap(15, 20),
		btn3,
		ButtonMargin(60, 20),
	)

	return res
}

func NewProjectButtons() fyne.CanvasObject {

	btn1 := NewCButton("🎨 New Theme", func() {})

	btn2 := NewCButton("📌 New Icon", func() {})

	btns := container.NewHBox(
		ButtonMargin(80, 20),
		btn1,
		ButtonGap(20, 20),
		btn2,
		ButtonMargin(80, 20),
	)

	return btns
}
