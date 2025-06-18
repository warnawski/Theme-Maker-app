package custom

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type CustomButton struct {
	widget.BaseWidget
	Text            string
	Importance      widget.ButtonImportance
	OnTapped        func()
	hovered, tapped bool
	tapAnim         *fyne.Animation
}

func (cb *CustomButton) CreateRenderer() fyne.WidgetRenderer {
	label := canvas.NewText(cb.Text, color.Black)
	label.Alignment = fyne.TextAlignCenter
	label.TextSize = 14

	rect := canvas.NewRectangle(color.Transparent)
	rect.StrokeColor = color.Black
	rect.StrokeWidth = 2
	rect.CornerRadius = 8

	return &CustomButtonRenderer{
		btn:     cb,
		label:   label,
		rect:    rect,
		objects: []fyne.CanvasObject{rect, label},
	}
}

type CustomButtonRenderer struct {
	btn     *CustomButton
	label   *canvas.Text
	rect    *canvas.Rectangle
	objects []fyne.CanvasObject
}

func (r *CustomButtonRenderer) Layout(size fyne.Size) {
	r.rect.Resize(size)

	textSize := r.label.MinSize()
	r.label.Move(fyne.NewPos((size.Width-textSize.Width)/2, (size.Height-textSize.Height)/2))
	r.label.Resize(textSize)
}

func (r *CustomButtonRenderer) MinSize() fyne.Size {
	return fyne.NewSize(180, 80)
}

func (r *CustomButtonRenderer) Refresh() {
	r.label.Color = theme.Color(theme.ColorNameForeground)
	r.rect.StrokeColor = theme.Color(theme.ColorNamePrimary)
	r.rect.FillColor = color.Transparent
	r.label.TextSize = 22

	switch {
	case r.btn.tapped:
		r.rect.FillColor = theme.PressedColor()
	case r.btn.hovered:
		r.rect.FillColor = theme.HoverColor()
	case r.btn.Importance == widget.HighImportance:
		r.rect.FillColor = theme.PrimaryColor()
	default:
		r.rect.FillColor = color.Transparent
	}

	canvas.Refresh(r.btn)
}

func (r *CustomButtonRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *CustomButtonRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *CustomButtonRenderer) Destroy() {}

func NewCButton(text string, onClick func()) *CustomButton {

	btn := &CustomButton{Text: text, OnTapped: onClick}

	btn.tapAnim = fyne.NewAnimation(150*time.Millisecond, func(f float32) {
		if f > 0.5 {
			btn.tapped = false
		} else {
			btn.tapped = true
		}
		btn.Refresh()
	})
	btn.ExtendBaseWidget(btn)

	return btn
}

func (cb *CustomButton) Tapped(*fyne.PointEvent) {
	cb.tapAnim.Stop()
	cb.tapAnim.Start()
}

func (cb *CustomButton) MouseIn(*desktop.MouseEvent) {
	cb.hovered = true
	cb.Refresh()
}

func (cb *CustomButton) MouseOut() {
	cb.hovered = false
	cb.Refresh()
}

func (cb *CustomButton) MouseMoved(*desktop.MouseEvent) {}
