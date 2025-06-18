package theme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type SolarizedDarkTheme struct{}

func (t SolarizedDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{0, 43, 54, 255}
	case theme.ColorNameForeground:
		return color.RGBA{131, 148, 150, 255}
	case theme.ColorNameButton:
		return color.RGBA{7, 54, 66, 255}
	case theme.ColorNamePrimary:
		return color.RGBA{38, 139, 210, 255}
	case theme.ColorNameHover:
		return color.RGBA{88, 110, 117, 255}
	case theme.ColorNameSelection:
		return color.RGBA{42, 161, 152, 255}
	case theme.ColorNameInputBackground:
		return color.RGBA{0, 43, 54, 255}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (t SolarizedDarkTheme) Font(ts fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(ts)
}
func (t SolarizedDarkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (t SolarizedDarkTheme) Size(s fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(s)
}

func NewSolarizedDarkTheme() fyne.Theme {
	return &SolarizedDarkTheme{}
}
