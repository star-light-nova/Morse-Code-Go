package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/star-light-nova/Morse-Code-Go/pkg/morse"
)

func radioAndMultiLinesContainer(window fyne.Window) *fyne.Container {
    radio := radioWidget()
    copy := copyButton()
    inputLabel, inputMultiLineEntry, outputLabel, outputMultilineEntry := multilines()

    radio.OnChanged = func(option string) {
        if option == OPT_ENCODE {
            inputMultiLineEntry.SetPlaceHolder("Enter text in English.")
        } else if option == OPT_DECODE {
            inputMultiLineEntry.SetPlaceHolder("Enter text in Morse Code.")
        }
    }

    inputMultiLineEntry.OnChanged = func(text string) {
        if radio.Selected == OPT_ENCODE {
            outputMultilineEntry.SetText(morse.Encode(text))
        } else if radio.Selected == OPT_DECODE {
            outputMultilineEntry.SetText(morse.Decode(text))
        }
    }

    copy.OnTapped = func() {
        window.Clipboard().SetContent(outputMultilineEntry.Text)
    }

    inputBox := container.NewVBox(
        radio,
        inputLabel,
        inputMultiLineEntry,
    )

    outputBox := container.NewVBox(
        outputLabel,
        outputMultilineEntry,
        copy,
    )

    assembledGrid := container.NewGridWithColumns(1, inputBox, outputBox)

    return container.New(layout.NewPaddedLayout(), assembledGrid)
}
