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
 * Complete the 'extraLongFactorials' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

type LL struct {
	List *longInt
}

type longInt struct {
	val  int32
	next *longInt
}

func createNode(n int32) *longInt {
	return &longInt{
		val:  n,
		next: nil,
	}
}

func (lst *LL) insertAtBeginning(data int32) {
	if nil == lst.List {
		lst.List = createNode(data)
		return
	}

	tempNode := createNode(data)
	head := lst.List
	tempNode.next = head
	lst.List = tempNode
}

func (lst *LL) deleteFromBigin() int32 {
	var val int32
	if nil != lst && nil != lst.List {
		head := lst.List
		val = head.val
		lst.List = head.next
		head = nil
	}
	return val
}

func (lst *LL) reverseList() {
	if nil != lst && nil != lst.List {
		head := lst.List
		firstPtr := head
		temp := new(longInt)
		temp = nil

		for head != nil && nil != firstPtr {
			firstPtr = head
			head = firstPtr.next
			firstPtr.next = nil
			firstPtr.next = temp
			temp = firstPtr
		}
		lst.List = temp
	}
}

func (lst *LL) readFromLast(data int32) {
	tempLst := new(LL)
	if nil != lst && nil != lst.List {
		lst.reverseList()
		var remain int32
		for nil != lst.List {
			val := lst.deleteFromBigin()
			//fmt.Println("Get Value from List: ", val)
			newMultVal := (val * data) + remain
			remain = 0
			if newMultVal <= 9 {
				tempLst.insertAtBeginning(newMultVal)
				continue
			}
			if newMultVal > 9 {
				tempLst.insertAtBeginning(newMultVal % 10)
				remain = newMultVal / 10
			}
		}
		if remain > 0 {
			for remain > 0 {
				tempLst.insertAtBeginning(remain % 10)
				remain /= 10
			}
		}
	} else {
		lst.insertAtBeginning(data)
		return
	}
	lst.List = tempLst.List
}

func (lst *LL) printList() {
	if nil != lst && nil != lst.List {
		head := lst.List
		for nil != head {
			fmt.Printf("%d", head.val)
			head = head.next
		}
	}
}

func (lst *LL) factMultiPly(val int32) {
	lst.readFromLast(val)
}

func extraLongFactorials(n int32) {
	// Write your code here
	bigLst := new(LL)
	for n > 0 {
		bigLst.factMultiPly(n)
		n--
	}
	bigLst.printList()

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	extraLongFactorials(n)
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
