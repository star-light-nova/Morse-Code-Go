package reversemap

func Reverse(originalMap map[string]string) map[string]string {
    result := map[string]string{}

    for key, val := range originalMap {
        result[val] = key
    }

    return result
}
