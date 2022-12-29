package main

import (
	"flag"
	"fmt"
	"time"
)

func flags(stringForSearch *string, substringSearched *string) {
	flag.StringVar(stringForSearch, "str", "unidentified", "set str")
	flag.StringVar(substringSearched, "substr", "unidentified", "set substr")
	flag.Parse()
}

func searching(stringInput string, substring string) (result bool) {

	runeStr := []rune(stringInput)
	runeSub := []rune(substring)

	for i, j := 0, len(runeSub); j < len(runeStr); i, j = i+1, j+1 {
		if string(runeStr[i:j]) == string(runeSub[:]) {
			return true
		}
	}
	return
}

func main() {
	var stringForSearch, substringSearched string

	flags(&stringForSearch, &substringSearched)

	//Простой способ поиска подстроки в строке.
	//result := strings.Contains(stringForSearch, substringSearched)

	result := searching(stringForSearch, substringSearched)

	if stringForSearch == "unidentified" || substringSearched == "unidentified" {
		fmt.Println("Один из аргументов команды пуст, поиск невозможен.")
	} else {
		fmt.Printf("Будет произведен поиск подстроки \"%v\" в строке \"%v\"...\n", substringSearched, stringForSearch)
		time.Sleep(time.Second * 3)
		fmt.Println("Еще немного...")
		time.Sleep(time.Second * 2)
		fmt.Println("Готово!")
		fmt.Printf("Результат поиска: %v", result)
	}
}
