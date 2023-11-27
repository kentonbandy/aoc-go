package helpers

import (
	"bufio"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
)

func GetLines() []string {
	// get the path of the file that called this function
	_, path, _, _ := runtime.Caller(1)
	// the input file should be in the same directory as the main file
	inputfilepath := strings.Replace(path, "main.go", "input.txt", -1)
	// attempt to open the input file
	file, err := os.Open(inputfilepath)
	if err != nil {
		// file doesn't exist - fetch data and create file
		pathslice := strings.Split(path, "/")
		day := StringToInt(pathslice[6])
		year := StringToInt(pathslice[5])
		fetchDataAndCreateInputFile(day, year)
		file, err = os.Open(inputfilepath)
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	lines := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func fetchDataAndCreateInputFile(day int, year int) {
	fmt.Printf("Attempting to create input file for year %d day %d\n", year, day)

	// get data
	content := getInputData(year, day)

	// create/write file
	filename := fmt.Sprintf("%d/%02d/input.txt", year, day)
	writeToFile(content, filename)

	if content != "" {
		fmt.Println("Success!")
	}
}

func getInputData(year int, day int) string {
	cookie := getSessionCookie()

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Unable to create request: %v", err)
	}

	req.AddCookie(&cookie)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Unable to fetch data: %v", err)
	}

	content, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unable to read data: %v", err)
	}

	return string(content)
}

func writeToFile(content string, filename string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func getSessionCookie() http.Cookie {
	cookie := os.Getenv("AOC_SESSION_COOKIE")
	if cookie == "" {
		envErr := godotenv.Load(".env")
		if envErr != nil {
			panic("Unable to load .env file")
		}
		cookie = os.Getenv("AOC_SESSION_COOKIE")
	}
	sessionCookie := http.Cookie{Name: "session", Value: cookie}
	return sessionCookie
}
