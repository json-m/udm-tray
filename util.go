package main

// short helper function to handle debug messages differently
func l(msg string) {
	if debugMode {
		debug.Println(msg)
		return
	}
}
