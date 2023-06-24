package app

import (
    "fyne.io/fyne/v2/widget"
)

func multilines() (*widget.Label, *widget.Entry, *widget.Label, *widget.Entry) {
    readOnlyOutput := widget.NewMultiLineEntry()
    readOnlyOutput.Disable()

    return widget.NewLabel("Input"), widget.NewMultiLineEntry(), widget.NewLabel("Output"), readOnlyOutput
}
