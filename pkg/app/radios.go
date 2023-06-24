package app

import (
	"log"

	"fyne.io/fyne/v2/widget"
)


func radioWidgets() *widget.RadioGroup {
    options := []string{OPT_ENCODE, OPT_DECODE}

    radioGroup := widget.NewRadioGroup(options, func(value string) {
        log.Println("Radio value set to", value)
    })

    radioGroup.SetSelected(OPT_ENCODE)
    radioGroup.Horizontal = true

    return radioGroup
}

