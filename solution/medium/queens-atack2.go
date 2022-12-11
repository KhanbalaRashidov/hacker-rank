package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'queensAttack' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER k
 *  3. INTEGER r_q
 *  4. INTEGER c_q
 *  5. 2D_INTEGER_ARRAY obstacles
 */
func Max(num1, num2 int32) int32 {
	result := math.Max(float64(num1), float64(num2))
	return int32(result)
}
func Min(num1, num2 int32) int32 {
	result := math.Min(float64(num1), float64(num2))
	return int32(result)
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// Write your code here
	var row_left int32 = 0
	var row_right int32 = n + 1
	var col_up int32 = n + 1
	var col_down int32 = 0
	f := func(x int32) int32 {
		return n - x + 1
	}
	d1_up := (r_q + c_q + 2*(n-Max(r_q, c_q))) + 1
	d1_down := (r_q + c_q - 2*Min(r_q, c_q))
	// d1_q := r_q + c_q
	d2_up := (f(r_q) + c_q + 2*(n-Max(f(r_q), c_q)))
	d2_down := (f(r_q) + c_q + 2*(n-Max(f(r_q), c_q))) + 1
	//d2_q:=(f(r_q)+c_q)
	d_q := r_q - c_q
	s_q := r_q + c_q

	for _, r_v := range obstacles {
		for r_o, c_o := range r_v {
			if r_o == int(r_q) {
				if c_o > c_q {
					row_right = Min(row_right, c_o)
				} else {
					row_left = Max(row_left, c_o)
				}
			} else if c_o == c_q {
				if r_o > int(r_q) {
					col_up = Min(col_up, int32(r_o))
				} else {
					col_down = Max(col_down, int32(r_o))
				}
			} else if r_o-int(c_o) == int(d_q) {
				d1_o := (r_o + int(c_o))
				if r_o > int(r_q) {
					d1_up = Min(d1_up, int32(d1_o))
				} else {
					d1_down = Max(d1_down, int32(d1_o))
				}
			} else if r_o+int(c_o) == int(s_q) {
				d2_o := (f(int32(r_o)) + c_o)
				if c_o > c_q {
					d2_down = Min(d2_down, d2_o)
				} else {
					d2_up = Max(d2_up, d2_o)
				}
			}
		}
	}

	return (row_right - row_left + col_up - col_down + d1_up - d1_down + d2_down - d2_up - 8)

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

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
	checkError(err)
	r_q := int32(r_qTemp)

	c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
	checkError(err)
	c_q := int32(c_qTemp)

	var obstacles [][]int32
	for i := 0; i < int(k); i++ {
		obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var obstaclesRow []int32
		for _, obstaclesRowItem := range obstaclesRowTemp {
			obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
			checkError(err)
			obstaclesItem := int32(obstaclesItemTemp)
			obstaclesRow = append(obstaclesRow, obstaclesItem)
		}

		if len(obstaclesRow) != 2 {
			panic("Bad input")
		}

		obstacles = append(obstacles, obstaclesRow)
	}

	result := queensAttack(n, k, r_q, c_q, obstacles)

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
