package cliui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
	"github.com/gvcgo/version-manager/internal/download"
	"github.com/gvcgo/version-manager/internal/installer/install"
	"github.com/gvcgo/version-manager/internal/terminal"
	"github.com/gvcgo/version-manager/internal/tui/table"
	"github.com/gvcgo/version-manager/internal/utils"
)

type SDKSearcher struct {
	SdkList download.SDKList
}

func NewSDKSearcher() *SDKSearcher {
	return &SDKSearcher{
		SdkList: make(download.SDKList),
	}
}

func (v *SDKSearcher) GetShaBySDKName(sdkName string) (ss string) {
	if s, ok := v.SdkList[sdkName]; ok {
		ss = s.Sha256
	}
	return
}

func (v *SDKSearcher) GetSDKItemByName(sdkName string) (item download.SDK) {
	item = v.SdkList[sdkName]
	return
}

func (v *SDKSearcher) Show() (nextEvent, selectedItem string) {
	v.SdkList = download.GetSDKList()

	ll := table.NewList()
	ll.SetListType(table.SDKList)
	v.RegisterKeyEvents(ll)

	_, w, _ := terminal.GetTerminalSize()
	if w > 30 {
		w -= 30
	} else {
		w = 120
	}
	ll.SetHeader([]table.Column{
		{Title: "sdkname", Width: 20},
		{Title: "homepage", Width: w},
	})
	rows := download.GetSDKSortedRows(v.SdkList)
	if len(rows) == 0 {
		gprint.PrintWarning("No sdk found!")
		gprint.PrintWarning("Please check if you have a proxy or reverse proxy available.")
		return
	}
	ll.SetRows(rows)
	ll.Run()

	selectedItem = ll.GetSelected()
	nextEvent = ll.NextEvent
	return
}

func (v *SDKSearcher) ShowInstalledOnly() (nextEvent, selectedItem string) {
	ll := table.NewList()
	ll.SetListType(table.SDKList)
	v.RegisterKeyEvents(ll)

	_, w, _ := terminal.GetTerminalSize()
	if w > 30 {
		w -= 30
	} else {
		w = 120
	}
	ll.SetHeader([]table.Column{
		{Title: "installed sdk", Width: 20},
		{Title: "homepage", Width: w},
	})
	rows := download.GetSDKSortedRows(v.SdkList)

	installedRows := []table.Row{}
	for _, r := range rows {
		if install.IsSDKInstalledByVMR(r[0]) {
			installedRows = append(installedRows, r)
		}
	}
	if len(installedRows) == 0 {
		gprint.PrintWarning("no installed sdk found!")
		return
	}
	ll.SetRows(installedRows)
	ll.Run()

	selectedItem = ll.GetSelected()
	nextEvent = ll.NextEvent
	return
}

func (v *SDKSearcher) PrintInstalledSDKs() {
	v.SdkList = download.GetSDKList()
	rows := download.GetSDKSortedRows(v.SdkList)

	installedRows := []table.Row{}
	for _, r := range rows {
		if install.IsSDKInstalledByVMR(r[0]) {
			installedRows = append(installedRows, r)
		}
	}
	if len(installedRows) == 0 {
		gprint.PrintWarning("no installed sdk found!")
		return
	}

	for _, r := range installedRows {
		gprint.PrintInfo(r[0])
	}
}

func (v *SDKSearcher) RegisterKeyEvents(ll *table.List) {
	// Open homepage.
	ll.SetKeyEventForTable("o", table.KeyEvent{
		Event: func(key string, l *table.List) tea.Cmd {
			sr := l.Table.SelectedRow()
			if len(sr) > 1 {
				utils.OpenURL(sr[1])
			}
			l.NextEvent = ""
			return nil
		},
		HelpInfo: "open homepage of the selected sdk",
	})
}
