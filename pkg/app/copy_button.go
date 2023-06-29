package app

import (
	"log"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func copyButton() *widget.Button {
    button := widget.NewButtonWithIcon("Copy output", theme.ContentCopyIcon(), func() {
        log.Println("Copy button clicked.")
    })

    return button
}
