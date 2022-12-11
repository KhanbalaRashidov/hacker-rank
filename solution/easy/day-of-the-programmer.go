package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	// Write your code here
	var date string
	var programmerDay = 256
	var monthOneToEight = []int{31, 31, 30, 31, 30, 31, 31}
	if year > 1918 {
		if year%400 == 0 {
			monthOneToEight = append(monthOneToEight, 29)
		} else if year%100 == 0 {
			monthOneToEight = append(monthOneToEight, 28)
		} else if year%4 == 0 {
			monthOneToEight = append(monthOneToEight, 29)
		} else {
			monthOneToEight = append(monthOneToEight, 28)
		}
	} else {
		if year%4 == 0 {
			monthOneToEight = append(monthOneToEight, 29)
		} else {
			monthOneToEight = append(monthOneToEight, 28)
		}
	}
	var dayYear int
	for _, value := range monthOneToEight {
		dayYear += value
	}
	date = fmt.Sprintf("%d.09.%d", programmerDay-dayYear, year)

	if year == 1918 {
		date = fmt.Sprintf("26.09.%d", year)
	}

	return date

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	yearTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := dayOfProgrammer(year)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
