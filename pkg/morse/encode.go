package morse

import (
    "github.com/star-light-nova/morse-code-go/pkg/dictionary"
    "strings"
)

func Encode(text string) string {
    textRunes := []rune(strings.ToUpper(text))
    result := ""

    for _, runeValue := range textRunes {
        if runeValue == ' ' {
            result += WORD_SEPARATOR
        } else {
            result += dictionary.All()[runeValue]
        }

        result += LETTER_SEPARATOR
    }

    return strings.TrimRight(result, " ")
}
