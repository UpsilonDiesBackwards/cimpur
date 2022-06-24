package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func CreateMainContextWindow(wTitle string, wWidth, wHeight float32) fyne.App {
	// Create a new application and set its window size to that of the func's input parameters
	Application := app.New()
	window := Application.NewWindow(wTitle)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(wWidth, wHeight))

	// Show the window then check for cleanup
	window.ShowAndRun()
	cleanUp()

	// Return as singleton if not already created
	if Application != nil {
		_ = fmt.Errorf("there is already an application main context window opened")
		return Application
	}
	return Application
}
