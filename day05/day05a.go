package main

import (
	"math"
	"fmt"
	"strconv"
)

func loadBoardingPasses (filePath string) (bps []string) {
	bps = []string{"FBFBBFFRLR"}
	return bps 
}

func main() {

	bps := loadBoardingPasses("input.txt")
	//outer loop that reads all seats ... for 
		seatCode := bps[0]
		row := seatCode[0:7]
		fmt.Println("row:", row)
		//aisle := seatCode[7:10]
		var r uint8 = 127 
		for x := 0; x < len(row); x++ {
			b := (x - 7) * -1
			fmt.Println("b:", strconv.Itoa(int(b)))
			bit := math.Pow(2, float64(b))
			fmt.Println("bit:", strconv.Itoa(int(bit)))
			//if a 'B'...	
			if row[x] == 66 {
				r = r ^ uint8(bit - 1)
			}
			fmt.Println("r:", strconv.Itoa(int(r)))
		}
	fmt.Println("Final r:", strconv.Itoa(int(r)))
}
				

