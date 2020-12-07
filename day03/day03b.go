package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

//nb: each line in input is 32 chars with line feed
func GetTrees( moves [2]int, treesArr *[5]int, ix int, byteSlice []byte) {
	right := moves[0]
	down := moves[1]
	var trees int	
	var lfFlag = false
	fileBytes := len(byteSlice)

	//walk the array, simulating "downward" movement by jumping a line's worth of bytes ahead	
	for i := 0; i < fileBytes; {
		//reset end-of-line flag
		lfFlag = false
		//all of this is to track whether or not we should wrap an entire line
		for j := 1; j <= right && i < fileBytes; {
			i++
			j++
			//we're crossing a line feed
			if i < fileBytes && byteSlice[i] == 10 {
				lfFlag = true
				//lf isn't on the "map", so make an extra move
				i++
			}
		}
		//if we didn't cross a line boundary
		//we did cross a line boundary, but b/c of multiple downward moves, we still must increment
		//by a factor of 32
		if !lfFlag {
			i += (32 * down) 
		} else if lfFlag && down > 1 {
			i += ( 32 * (down - 1) )
		}

		if i < fileBytes && byteSlice[i] == 35 {
			trees++
		}
	}
	(*treesArr)[ix] = trees
	return
}

func main() {
	fileName := os.Args[1]
	bs, err := ioutil.ReadFile(fileName) 
	if err != nil {
		//nop	
	}

	var movesArr = [5][2]int{ {1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2} }
	var treesArr [5]int

	for i, t := range movesArr {
		GetTrees(t, &treesArr, i, bs)
		fmt.Println("Trees", i, ":", treesArr[i])
	}

	var treeProduct = func(treesArr *[5]int) int {
		prod := 1	
		for i:=0; i<len(treesArr); i++ {
			prod *= treesArr[i]	
		}
		return prod
	}

	fmt.Println("Trees: ", treeProduct(&treesArr))
}
