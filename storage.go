package main

var store = make(map[string]string)

func SaveURL(short, original string) {
    store[short] = original
}

func GetURL(short string) (string, bool) {
    val, ok := store[short]
    return val, ok
}