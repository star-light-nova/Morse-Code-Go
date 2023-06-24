package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)


func Start() {
    _app := app.New()
    window := _app.NewWindow("Morse Code Application")

    inputLabel, inputMultiLineEntry, outputLabel, outputMultilineEntry := multilines()

    vboxContainer := container.NewVBox(
        radioWidgets(),
        inputLabel,
        inputMultiLineEntry,
        outputLabel,
        outputMultilineEntry,
    )

    window.SetContent(vboxContainer)
    window.Resize(fyne.NewSize(DEFAULT_WIDTH, window.Content().MinSize().Width))
    window.ShowAndRun();
}
