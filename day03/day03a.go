package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	fileName := os.Args[1]
	bs, err := ioutil.ReadFile(fileName) 
	if err != nil {
		//nop	
	}
	var trees int	
	var crBool = false
	fileBytes := len(bs)	
	//walk the array, simulating "downward" movement by jumping a line's worth of bytes ahead	
	for i := 0; i < fileBytes; {
		//reset end-of-line flag
		crBool = false
		//all of this is to track whether or not we should wrap an entire line
		for j := 1; j < 4 && i < fileBytes; {
			i++
			j++
			//we're crossing a line feed
			if i < fileBytes && bs[i] == 10 {
				crBool = true
				//lf isn't on the "map", so make an extra move
				i++
			}
		}
		//if we didn't cross a line boundary
		if !crBool {
			i += 32
		}
		if i < fileBytes && bs[i] == 35 {
			trees++
		}
	}
	fmt.Println("Trees: ", trees)
}
