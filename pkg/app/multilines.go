package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)


func multilines() (*widget.Label, *widget.Entry, *widget.Label, *widget.Entry) {
    inputLabel, inputMultiline := inputLabelMultiline()
    outputLabel, outputMultiline := outputLabelMultiline()

    return inputLabel, inputMultiline, outputLabel, outputMultiline
}

func inputLabelMultiline() (*widget.Label, *widget.Entry){
    multiline := widget.NewMultiLineEntry()

    multiline.Wrapping = fyne.TextWrapWord

    // The default for the radios is "Encode"
    multiline.SetPlaceHolder("Enter text in English.")

    return widget.NewLabel("Input"), multiline
}

func outputLabelMultiline() (*widget.Label, *widget.Entry) {
    readOnlyOutput := widget.NewMultiLineEntry()

    readOnlyOutput.Wrapping = fyne.TextWrapWord

    // Function `Disable()` makes Entry read-only
    readOnlyOutput.Disable()

    return widget.NewLabel("Output"), readOnlyOutput
}
