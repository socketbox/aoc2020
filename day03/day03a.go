package main

import (
	"fmt"
	"os"
	"io/ioutil"
	//"strconv"
)

func main() {
	fileName := os.Args[1]
	bs, err := ioutil.ReadFile(fileName) 
	if err != nil {
		//nop	
	}
	var trees, move int	
	var crBool = false
	bytesLen := len(bs)	
	for i := 0; i < bytesLen; {
		//reset end-of-line flag
		crBool = false
		for j := 1; j < 4 && i < bytesLen; j++ {
			i++
			//we're crossing a line feed
			if i < bytesLen && bs[i] == 10 {
				crBool = true
			}
		}
		//if we didn't cross a line boundary
		if !crBool {
			i += 32
		}
		if i < bytesLen && bs[i] == 35 {
			trees++
		}
		if i < bytesLen { fmt.Println("move: ", move, "; i: ", i, "; bs[i]: ", fmt.Sprintf("%c", bs[i])) }
	}
	fmt.Println("Trees: ", trees)
}
