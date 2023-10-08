build:
	gox -osarch="windows/amd64" -output="udm-tray"

release:
	gox -osarch="windows/amd64" -ldflags -H=windowsgui -output="udm-tray"
	zip udm-tray-release.zip udm-tray.exe config.yml.example
	rm udm-tray.exe

clean:
	rm udm-tray*
