package main

import (
	"os"
	"runtime"
	"fmt"

	"github.com/getlantern/systray"
	"github.com/pocke/lemonade/lemon"
)

func tray() {
	systray.Run(func() {
		systray.SetTitle("Lemonade")
		systray.SetTooltip("Lemon")

		if runtime.GOOS == "windows" {
			data, _ := Asset("icon/icon.ico")
			systray.SetIcon(data)
		} else {
			data, _ := Asset("icon/icon.png")
			systray.SetIcon(data)
		}
    mVersion := systray.AddMenuItem(fmt.Sprintf("Lemonade version: %s", lemon.Version), "")
    mVersion.Disable()
		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
		go func() {
			<-mQuit.ClickedCh
			systray.Quit()
			os.Exit(lemon.Success)
		}()
	})
}
