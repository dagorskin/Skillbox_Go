package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Создание флагов.
func flags(fileOneName *string, fileOneContent *string, fileTwoName *string, fileTwoContent *string) {
	flag.StringVar(fileOneName, "fname", "unidentified1", "set first")
	flag.StringVar(fileOneContent, "fcontent", "unidentified2", "set fcontent")
	flag.StringVar(fileTwoName, "sname", "unidentified3", "set second")
	flag.StringVar(fileTwoContent, "scontent", "unidentified4", "set scontent")
	flag.Parse()
}

// Создание файла и запись в него.
func creationAndWrite(fileName, fileContent string) {
	file, err := os.Create(fileName + ".txt")
	if err != nil {
		fmt.Println("Error at the file creation stage:", err)
		os.Exit(1)
	}

	_, err = file.WriteString(fileContent)
	if err != nil {
		fmt.Println("Error at the stage of writing to the file:", err)
		os.Exit(2)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("Error at the file closing stage:", err)
			os.Exit(3)
		}
	}(file)
}

// Чтение файла.
func readFile(fileName string) []byte {
	fileContent, err := os.ReadFile(fileName + ".txt")
	if err != nil {
		fmt.Println("Error at the file reading stage:", err)
		os.Exit(4)
	}
	return fileContent
}

func main() {
	var fileOneName, fileOneContent, fileTwoName, fileTwoContent string

	flags(&fileOneName, &fileOneContent, &fileTwoName, &fileTwoContent)

	switch {
	case fileOneName != "unidentified" && fileTwoName == "unidentified":
		creationAndWrite(fileOneName, fileOneContent)
		fmt.Println("Только контент первого файла:", string(readFile(fileOneName)))

	case fileOneName == "unidentified" && fileTwoName != "unidentified":
		creationAndWrite(fileTwoName, fileTwoContent)
		fmt.Println("Только контент второго файла:", string(readFile(fileTwoName)))

	default:
		creationAndWrite(fileOneName, fileOneContent)
		creationAndWrite(fileTwoName, fileTwoContent)

		readOneFile, err := os.ReadFile(fileOneName + ".txt")
		if err != nil {
			fmt.Println("Error at the file reading stage:", err)
			os.Exit(4)
		}
		readTwoFile, err := os.ReadFile(fileTwoName + ".txt")
		if err != nil {
			fmt.Println("Error at the file reading stage:", err)
			os.Exit(4)
		}

		filesContent := []string{string(readOneFile), string(readTwoFile)}
		content := strings.Join(filesContent, "\n")
		creationAndWrite("result", content)

		fmt.Print("Контент двух файлов конкатенированный в один файл:\n", string(readFile("result")))
	}
}
