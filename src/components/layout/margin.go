package layout

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func ButtonMargin(w float32, h float32) fyne.CanvasObject {

	margin := canvas.NewRectangle(color.Transparent)
	margin.SetMinSize(fyne.NewSize(w, h))
	return margin
}

func TopMargin(h float32) fyne.CanvasObject {
	top := canvas.NewRectangle(color.Transparent)
	top.SetMinSize(fyne.NewSize(550, h))

	return top
}
