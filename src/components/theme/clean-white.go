package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CleanWhiteTheme struct{}

func (t CleanWhiteTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{245, 245, 245, 255}
	case theme.ColorNameForeground:
		return color.RGBA{20, 20, 20, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{0, 122, 255, 255}
	case theme.ColorNameHover:
		return color.RGBA{200, 200, 255, 255}
	case theme.ColorNameSelection:
		return color.RGBA{0, 100, 255, 255}
	case theme.ColorNameButton:
		return color.RGBA{255, 255, 255, 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{255, 255, 255, 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t CleanWhiteTheme) Font(ts fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(ts)
}

func (t CleanWhiteTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t CleanWhiteTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func NewCleanWhiteTheme() fyne.Theme {
	return &CleanWhiteTheme{}
}
