package dictionary

func newDigits() *map[rune]string {
    return &map[rune]string {
        '0': "-----",
        '1': ".----",
        '2': "..---",
        '3': "...--",
        '4': "....-",
        '5': ".....",
        '6': "-....",
        '7': "--...",
        '8': "---..",
        '9': "----.",
    }
}
