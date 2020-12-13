package main

import (
	"os"
	"io"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"sort"
	"math"
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
	//arrSz == 2 + number of ints in file
	const arrSz int = 33 
	var intArr [arrSz] int
	for i:=1; i < len(intArr); i++ {
		str, rsErr = reader.ReadString(0x0a)
		if rsErr != nil && rsErr == io.EOF {
			//prints EOF
			//fmt.Println(rsErr)
		} else { 
			conv, _ := strconv.Atoi(strings.TrimSpace(str))
			intArr[i] = conv
		}
	}
	sort.Ints(intArr[1:arrSz-1])
	intArr[0]=0
	intArr[arrSz-1] = intArr[arrSz-2]+3
	fmt.Println(intArr)
	var cmbCt, prod float64 
	for i := 0; i < len(intArr); i++ {
		if i == 0 {
			fmt.Println("zero")
			prod = 1
			cmbCt = -1
		} else if intArr[i] - intArr[i-1] == 3 { //is the current int three greater than its predecessor?
			fmt.Println("i:", i, "arr[i]:", intArr[i], "cmbCt: ", cmbCt, "two anchors")	
			if cmbCt > 0 {
				prod *= math.Pow(2, cmbCt)	//can get away with this b/c base is never more than 2
			}
			cmbCt = -1 
		//} else if 0 <= intArr[i] - intArr[i-1] && intArr[i] - intArr[i-1] <= 2 {
		} else if intArr[i] - intArr[i-1] == 1 || intArr[i] - intArr[i-1] == 2 {
			fmt.Println("i:", i, "arr[i]:", intArr[i], "cmbCt: ", cmbCt, "incrementing cmbCt")	
			cmbCt++
		
			/*if i > 2 {
				if intArr[i] - intArr[i-3] == 3 && cmbCt >= 0 {
					fmt.Println("i:", i, "arr[i]:", intArr[i], "cmbCt: ", cmbCt, "two anchors [i-3]")	
					if cmbCt > 0 {
						prod *= (cmbCt * 2)	//can get away with this b/c base is never more than 2
					}
					cmbCt = -1
				}
			}*/
		}
		fmt.Println("prod:", prod, "cmbCt", cmbCt)	
	}
	fmt.Println("prod: ", prod)
}
