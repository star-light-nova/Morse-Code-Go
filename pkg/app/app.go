package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)


func Start() {
    _app := app.New()
    window := _app.NewWindow("Morse Code Application")

    window.SetContent(radioAndMultiLinesContainer())

    window.Resize(fyne.NewSize(DEFAULT_WIDTH, window.Content().MinSize().Width))
    window.ShowAndRun();
}
