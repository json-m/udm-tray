module udm-tray

go 1.19

replace jason.lv/UDM-API => ../UDM-API

require (
	fyne.io/systray v1.10.0
	gopkg.in/yaml.v3 v3.0.1
	jason.lv/UDM-API v0.0.0-00010101000000-000000000000
)

require (
	github.com/godbus/dbus/v5 v5.0.4 // indirect
	github.com/tevino/abool v1.2.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
)
