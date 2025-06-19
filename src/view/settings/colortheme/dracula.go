package colortheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type DraculaTheme struct{}

func (t DraculaTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{40, 42, 54, 255}
	case theme.ColorNameForeground:
		return color.RGBA{248, 248, 242, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{255, 121, 198, 255}
	case theme.ColorNameHover:
		return color.RGBA{68, 71, 90, 255}
	case theme.ColorNameSelection:
		return color.RGBA{189, 147, 249, 255}
	case theme.ColorNameButton:
		return color.RGBA{68, 71, 90, 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{50, 52, 70, 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t DraculaTheme) Font(ts fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(ts)
}

func (t DraculaTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t DraculaTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func NewDraculaTheme() fyne.Theme {
	return &DraculaTheme{}
}
