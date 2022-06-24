package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func CreateUAuthWindow(wTitle string, wWidth, wHeight float32) fyne.App {
	// Create a new application and set its window size to that of the func's input parameters
	Application := app.New()
	window := Application.NewWindow(wTitle)
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(wWidth, wHeight))

	// Create User Sign in form
	userNameEntry := widget.NewEntry()
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Username: ", Widget: userNameEntry},
			{Text: "Password: ", Widget: widget.NewPasswordEntry()}},
		OnSubmit: func() {
			fmt.Println("Sumbitted")
			window.Close()
		},
	}
	window.SetContent(form)

	// Show the window then check for cleanup
	window.ShowAndRun()
	cleanUp()

	// Return as singleton if not already created
	if Application != nil {
		_ = fmt.Errorf("there is already an user auth window opened")
		return Application
	}
	return Application
}

func cleanUp() {

	fmt.Println("Exited Cimpur user auth.")

}
