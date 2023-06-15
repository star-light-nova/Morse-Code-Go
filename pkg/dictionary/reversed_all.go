package dictionary

func ReversedAll() map[string]rune {
    reversedAll := map[string]rune{}

    for key, val := range All() {
        reversedAll[val] = key
    }

    return reversedAll
}
