package morse

import (
	"strings"
	"github.com/StarLightNova/Morse-Code-Go/pkg/dictionary"
)

func Decode(morseText string) string {
    splitMorse := strings.Split(morseText, LETTER_SEPARATOR)
    result := ""

    for _, morseCharacter := range splitMorse {
        if morseCharacter == WORD_SEPARATOR {
            result += LETTER_SEPARATOR
        } else if morseCharacter != LETTER_SEPARATOR {
            result += string(dictionary.ReversedAll()[morseCharacter])
        }
    }

    return result
}
