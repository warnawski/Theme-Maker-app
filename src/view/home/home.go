package home

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	loader "github.com/warnawski/theme-maker/internal/loader"
	custom "github.com/warnawski/theme-maker/src/components/button"
	"github.com/warnawski/theme-maker/src/components/layout"
	"github.com/warnawski/theme-maker/src/components/tabs/common"
	"github.com/warnawski/theme-maker/src/view/icon"
)

func HomeView(tabmngr common.TapManager, tabs *container.AppTabs) fyne.CanvasObject {

	logo := loader.LoadFyneCanvasImage("resource/assets/theme-icon.png")
	logo.SetMinSize(fyne.NewSize(350, 350))

	view := container.NewVBox(
		container.NewCenter(logo),
		ListButton(),
		layout.TopMargin(30),
		NewProjectButtons(tabmngr, tabs),
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
		layout.ButtonMargin(60, 20),
		btn1,
		layout.ButtonGap(15, 20),
		btn2,
		layout.ButtonGap(15, 20),
		btn3,
		layout.ButtonMargin(60, 20),
	)

	return res
}

func NewProjectButtons(tabmngr common.TapManager, tabs *container.AppTabs) fyne.CanvasObject {

	btn1 := custom.NewCButton("🎨 New Theme", func() {})

	btn2 := custom.NewCButton("📌 New Icon", func() {

		tabs.SelectIndex(1)
		tabs.Items[1].Content = icon.NewIconView()
		log.Println("CKICK")

	})

	btns := container.NewHBox(
		layout.ButtonMargin(80, 20),
		btn1,
		layout.ButtonGap(20, 20),
		btn2,
		layout.ButtonMargin(80, 20),
	)

	return btns
}
