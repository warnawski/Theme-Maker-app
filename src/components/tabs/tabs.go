package tabs

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"github.com/warnawski/theme-maker/src/view/home"
	"github.com/warnawski/theme-maker/src/view/settings"
	void "github.com/warnawski/theme-maker/src/view/work/content"
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
			home.HomeView(t.tabs),
		),

		container.NewTabItemWithIcon(
			"Work Area                                                            ",
			theme.WarningIcon(),
			void.VoidContent(),
		),

		container.NewTabItemWithIcon("Docs", theme.HelpIcon(), void.VoidContent()),
		container.NewTabItemWithIcon("List", theme.ListIcon(), void.VoidContent()),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), settings.SettingsView(t.switcher)),
	)

	return tabs
}
