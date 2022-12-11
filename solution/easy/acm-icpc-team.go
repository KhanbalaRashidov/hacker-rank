package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'acmTeam' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts STRING_ARRAY topic as parameter.
 */

func acmTeam(topic []string) []int32 {
	// Write your code here
	var temp []int32
	var result = make([]int32, 2)
	for i := 0; i < len(topic); i++ {
		for j := i + 1; j < len(topic); j++ {
			var tempValue int32
			for k := 0; k < len(topic[i]); k++ {
				if string(topic[i][k]) == "1" || string(topic[j][k]) == "1" {
					tempValue += 1
				}
			}
			temp = append(temp, tempValue)
			tempValue = 0
		}
	}
	sort.Slice(temp, func(i, j int) bool {
		return temp[i] < temp[j]
	})
	for i := 0; i < len(temp); i++ {
		if temp[i] == temp[len(temp)-1] {
			result[1] += 1
		}
	}
	result[0] = temp[len(temp)-1]
	return result
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	// mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	// checkError(err)
	// m := int32(mTemp)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
