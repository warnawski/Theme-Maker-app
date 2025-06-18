package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type MonokaiTheme struct{}

func (t MonokaiTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{39, 40, 34, 255}
	case theme.ColorNameForeground:
		return color.RGBA{248, 248, 242, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{255, 85, 85, 255}
	case theme.ColorNameHover:
		return color.RGBA{102, 217, 239, 255}
	case theme.ColorNameSelection:
		return color.RGBA{166, 226, 46, 255}
	case theme.ColorNameButton:
		return color.RGBA{73, 72, 62, 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{44, 44, 38, 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t MonokaiTheme) Font(ts fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(ts)
}

func (t MonokaiTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t MonokaiTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func NewMonokaiTheme() fyne.Theme {
	return &MonokaiTheme{}
}
