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
 * Complete the 'bonAppetit' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY bill
 *  2. INTEGER k
 *  3. INTEGER b
 */

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32
	for i, val := range bill {
		if int32(i) == k {
			continue
		}
		sum += val
	}
	var check int32 = b - sum/2
	if check != 0 {
		fmt.Println(check)
	} else {
		fmt.Println("Bon Appetit")
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	billTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var bill []int32

	for i := 0; i < int(n); i++ {
		billItemTemp, err := strconv.ParseInt(billTemp[i], 10, 64)
		checkError(err)
		billItem := int32(billItemTemp)
		bill = append(bill, billItem)
	}

	bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	b := int32(bTemp)

	bonAppetit(bill, k, b)
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
