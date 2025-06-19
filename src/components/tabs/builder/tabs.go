package builder

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/warnawski/theme-maker/src/view/home"
	"github.com/warnawski/theme-maker/src/view/settings"
	"github.com/warnawski/theme-maker/src/view/work/content"
)

type TabBuilder struct {
	switcher settings.ThemeSwitcher
	tabs     *container.AppTabs
}

func NewTabBuilder(switcher settings.ThemeSwitcher) TabBuilder {
	builder := TabBuilder{switcher: switcher}
	builder.tabs = builder.BuildTab()
	return builder
}

func (t TabBuilder) BuildTab() *container.AppTabs {

	tabs := container.NewAppTabs(

		container.NewTabItemWithIcon(
			"Home                           ",
			theme.HomeIcon(),
			container.NewWithoutLayout(),
		),
		container.NewTabItemWithIcon(
			"Work Area                                                            ",
			theme.WarningIcon(),
			container.NewWithoutLayout(),
		),
		container.NewTabItemWithIcon("Docs", theme.HelpIcon(), container.NewWithoutLayout()),
		container.NewTabItemWithIcon("List", theme.ListIcon(), container.NewWithoutLayout()),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), container.NewWithoutLayout()),
	)

	t.tabs = tabs

	tabs.Items[0].Content = home.HomeView(t, tabs)
	tabs.Items[1].Content = content.VoidContent()
	tabs.Items[4].Content = settings.SettingsView(t.switcher)

	return tabs
}

func (t TabBuilder) AskIndex(title string) {

	for index, tab := range t.tabs.Items {
		if tab.Text == title {
			t.tabs.SelectIndex(index)
			break
		}
	}
}
