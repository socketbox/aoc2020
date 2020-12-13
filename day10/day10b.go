package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

func main() {
	fileName := os.Args[1]
	file, fErr := os.Open(fileName) 
	if fErr != nil {
		//nop	
	}
	reader := bufio.NewReader(file)
	var rsErr error	
	var str string 
	var intArr [11]int
	for i:=0; i < len(intArr); i++ {
		str, rsErr = reader.ReadString(0x0a)
		if rsErr != nil && rsErr == io.EOF {
			fmt.Println(rsErr)
		} else { 
			conv, _ := strconv.Atoi(strings.TrimSpace(str))
			intArr[i] = conv
		}
	}
	sort.Ints(intArr[:])
	var cmbCt, prod int	
	for i := 0; i < len(intArr); i++ {
		if i == 0 {
			prod = 1
			cmbCt = -1
		} else if intArr[i] - 3 == intArr[i-1] {
		//is the current int three greater than its predecessor?
		prod *= (cmbCt * 2)	//can get away with this b/c base is never more than 2
			cmbCt = -1
		} else if intArr[i] - 1 == intArr[i-1] {
			cmbCt++
		}
	}
	fmt.Println("prod: ", prod)
}
