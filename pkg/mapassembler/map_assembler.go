package mapassembler

func Assemble(maps ...map[rune]string) map[rune]string {
    asselmbedMap := map[rune]string{}

    for _, currentMap := range maps {
        for key, value := range currentMap {
            asselmbedMap[key] = value
        }
    }

    return asselmbedMap
}
