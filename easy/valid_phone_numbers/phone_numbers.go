package phone_numbers

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

func PhoneRegex() []string {
	phoneNumbers := readFile()

	r, _ := regexp.CompilePOSIX("^\\([[:digit:]]{3}\\)[[:space:]][[:digit:]]{3}-[[:digit:]]{4}$|^[[:digit:]]{3}-[[:digit:]]{3}-[[:digit:]]{4}$")
	var filteredNumbers []string
	for _, number := range phoneNumbers {
		if r.MatchString(number) {
			filteredNumbers = append(filteredNumbers, number)
		}
	}

	return filteredNumbers
}

func readFile() []string {
	file, err := os.Open("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var fileLines []string
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return fileLines
}
