package morse

import (
    "github.com/StarLightNova/Morse-Code-Go/pkg/dictionary"
    "strings"
)

func Translate(text string) string {
    textRunes := []rune(strings.ToUpper(text))
    result := ""

    for _, runeValue := range textRunes {
        if runeValue == ' ' {
            result += "/"
        } else {
            result += dictionary.All()[runeValue]
        }

        result += " "
    }

    return strings.TrimRight(result, " ")
}
