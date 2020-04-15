package main

import (
	"fmt"
	"github.com/aurora/autodeploy/models/hmc"
	"github.com/aurora/autodeploy/pkg/setting"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	"strconv"
)

type Slide func() (title string, content tview.Primitive)

func main() {
	hmc.Logon(setting.HmcUserName, setting.HmcPassWord)
	hmc.GetCPCId()
	hmc.App = tview.NewApplication()
	slides := []Slide{
		hmc.Help,
		hmc.Partition,
		hmc.SG,
		hmc.OS,
		hmc.OsStandardization,
		//hmc.Database,
		//hmc.AutoDeploy,
	}

	pages := tview.NewPages()

	info := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(false).
		SetHighlightedFunc(func(added, removed, remaining []string) {
			pages.SwitchToPage(added[0])
		})

	previousSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide - 1 + len(slides)) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	nextSlide := func() {
		slide, _ := strconv.Atoi(info.GetHighlights()[0])
		slide = (slide + 1) % len(slides)
		info.Highlight(strconv.Itoa(slide)).
			ScrollToHighlight()
	}
	for index, slide := range slides {
		title, primitive := slide()
		pages.AddPage(strconv.Itoa(index), primitive, true, index == 0)
		_, _ = fmt.Fprintf(info, `  ["%d"][darkcyan]%s[white][""]  `, index, title)
	}
	info.Highlight("0")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 1, true).
		AddItem(info, 1, 1, false)

	hmc.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyCtrlN {
			nextSlide()
		} else if event.Key() == tcell.KeyCtrlP {
			previousSlide()
		}
		return event
	})

	if err := hmc.App.SetRoot(layout, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}
