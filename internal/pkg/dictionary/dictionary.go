package dictionary

/*
* NOTE: P35S Represents - Punctuation Marks And Miscellaneous Signs
*/

type Dictionary struct {
    All map[rune]string
    Alphabet map[rune]string
    Digits map[rune]string
    P35S map[rune]string
}

func New() Dictionary {
    return Dictionary{
        Alphabet: *newAlphabet(),
        Digits: *newDigits(),
        P35S: *newP35S(),
    }
}
