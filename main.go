package main

import (
	"flag"
	"fmt"
	"fyne.io/systray"
	udm "jason.lv/UDM-API"
	"log"
	"os"
	"udm-tray/icon"
)

// create a debug mode logger
var debug = log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile)
var debugMode bool

func init() {
	// create a debug mode logger
	debugPtr := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if *debugPtr {
		debugMode = true
	}
	err := readConfig()
	if err != nil {
		log.Fatalf("Error reading config: %s\n", err)
	}

}

var Routes []udm.TrafficRouteStruct

func main() {
	var err error
	l(fmt.Sprintf("Connecting to %s as %s...\n", config.Host, config.User))

	// login
	config.Client, err = udm.CreateClient(config.User, config.Pass, config.Host, 5, config.SkipInsecure)
	if err != nil {
		log.Fatalf("Error creating client: %s\n", err)
	}

	// reset and update traffic rules with what is on the remote side
	tr, err := config.Client.TrafficRouteController_GetRoutes(config.Site)
	if err != nil {
		log.Fatalf("Error getting traffic rules: %s\n", err)
	}
	Routes = []udm.TrafficRouteStruct{}
	Routes = append(Routes, tr...)

	// print out traffic rules
	for _, rules := range Routes {
		l(fmt.Sprintf("Rule ID/Desc: %s - %s\n", rules.ID, rules.Description))
	}

	// prevent app from exiting
	systray.Run(onReady, onExit)
	done := make(chan bool)
	<-done
}

// systray stuff
func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("UDM Traffic Toggle")
	systray.SetTooltip("UDM Traffic Toggle")

	// add all Traffic Routes as menu items, and toggle them when clicked with addmenucheckbox
	for _, rule := range Routes {
		r := rule
		m := systray.AddMenuItemCheckbox(r.Description, "Toggle this traffic rule", r.Enabled)
		go func() {
			for {
				<-m.ClickedCh

				// toggle enabled field in Routes for r.ID
				var idx int
				for i, rr := range Routes {
					if rr.ID == r.ID {
						Routes[i].Enabled = !Routes[i].Enabled
						idx = i
					}
				}

				// make request to update rule
				err := config.Client.TrafficRouteController_UpdateRoute(config.Site, r.ID, Routes[idx])
				if err != nil {
					l(fmt.Sprintf("Error updating traffic rule: %s\n", err))
					continue
				}

				checked := m.Checked()
				if checked {
					m.Uncheck()
				} else {
					m.Check()
				}
			}
		}()

	}

	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

}

// will probably never use this
func onExit() {
	// clean up anything here???
}
