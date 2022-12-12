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
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	// Write your code here
	var typeBird = make(map[string]int32)
	// Write your code here
	for i := 0; i < len(arr); i++ {
		key := strconv.Itoa(int(arr[i]))
		_, exist := typeBird[key]
		if exist {
			typeBird[key] += 1
		} else {
			typeBird[key] = 1
		}
	}
	var keys = make([]string, 0, len(typeBird))
	var values = make([]int, 0, len(typeBird))
	for key, value := range typeBird {
		keys = append(keys, key)
		values = append(values, int(value))
	}
	sort.Strings(keys)
	sort.Ints(values)
	var result, temp int32
	for i, key := range keys {
		atoi, _ := strconv.Atoi(key)
		if i == 0 {
			fmt.Println(key, typeBird[key], "ASO", i)
			result = int32(atoi)
			temp = typeBird[key]
		} else if typeBird[key] > temp {
			if result < int32(atoi) {
				result = int32(atoi)
				temp = typeBird[key]
				fmt.Println(key, "HMMM")
			}
		}
	}
	fmt.Println(typeBird)
	return result

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := migratoryBirds(arr)

	fmt.Fprintf(writer, "%d\n", result)

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
