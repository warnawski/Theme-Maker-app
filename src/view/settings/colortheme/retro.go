package colortheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type RetroTheme struct{}

func (l RetroTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {

	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{15, 15, 20, 255}
	case theme.ColorNameForeground:
		return color.RGBA{200, 255, 255, 255}
	case theme.ColorNameButton:
		return color.RGBA{20, 20, 40, 255}
	case theme.ColorNameFocus, theme.ColorNamePrimary, theme.ColorNameInputBorder:
		return color.RGBA{0, 255, 180, 255}
	case theme.ColorNameHover:
		return color.RGBA{40, 60, 80, 255}
	case theme.ColorNameSelection:
		return color.RGBA{0, 255, 255, 255}
	case theme.ColorNameScrollBar:
		return color.RGBA{30, 255, 200, 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{25, 25, 35, 255}
	case theme.ColorNamePlaceHolder:
		return color.RGBA{100, 150, 150, 255}
	case theme.ColorNameError:
		return color.RGBA{255, 60, 60, 255}
	case theme.ColorNameSuccess:
		return color.RGBA{60, 255, 120, 255}
	case theme.ColorNameWarning:
		return color.RGBA{255, 200, 60, 255}
	case theme.ColorNameDisabled:
		return color.RGBA{80, 80, 100, 255}
	case theme.ColorNameDisabledButton:
		return color.RGBA{40, 40, 50, 255}
	case theme.ColorNameSeparator:
		return color.RGBA{50, 80, 100, 255}
	case theme.ColorNameShadow:
		return color.RGBA{10, 10, 20, 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (l RetroTheme) Font(textStyle fyne.TextStyle) fyne.Resource {

	return theme.DefaultTheme().Font(textStyle)
}

func (l RetroTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (l RetroTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func NewRetroTheme() fyne.Theme {
	return &RetroTheme{}
}
