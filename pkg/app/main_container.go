package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/StarLightNova/Morse-Code-Go/pkg/morse"
)

func radioAndMultiLinesContainer(window fyne.Window) *fyne.Container {
    radio := radioWidget()
    inputLabel, inputMultiLineEntry, outputLabel, outputMultilineEntry := multilines()
    copy := copyButton()

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

    vboxContainer := container.NewVBox(
        radio,
        inputLabel,
        inputMultiLineEntry,
        outputLabel,
        outputMultilineEntry,
        copy,
    )

    return vboxContainer
}
