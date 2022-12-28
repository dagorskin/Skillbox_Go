package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

func flags(stringForSearch *string, substringSearched *string) {
	flag.StringVar(stringForSearch, "str", "unidentified", "set str")
	flag.StringVar(substringSearched, "substr", "unidentified", "set substr")
	flag.Parse()
}

func main() {
	var stringForSearch, substringSearched string

	flags(&stringForSearch, &substringSearched)

	result := strings.Contains(stringForSearch, substringSearched)

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
