package main

import (
	"os"
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
	var intArr [103]int
	for i:=0; i < len(intArr); i++ {
		str, rsErr = reader.ReadString(0x0a)
		if rsErr != nil { 
			panic("error while reading line") 
		} else { 
			conv, _ := strconv.Atoi(strings.TrimSpace(str))
			intArr[i] = conv
		}
	}
	sort.Ints(intArr[:])
	var last, diff, j1, j3 int 
	for j:=0;j<len(intArr);j++ {
		diff = intArr[j]-last
		if diff == 1 {
			j1++
		} else if diff == 3 {
			j3++
		} else {
			panic("diff is neither 1 nor 3")
		}
		last = intArr[j]
	}
	//built-in adapter
	j3++
	fmt.Println("j1:",j1,"j3:",j3,"j1*j3:",j1*j3)

}
