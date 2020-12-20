package main

import (
	"bufio"
	"strings"	
	"fmt"
	"log"	
	"os"
)

type PassportMask uint8 

const (
	byr	= 1 << iota	
	iyr		
	eyr		
	hgt		
	hcl		
	ecl		
	pid		
	cid //128	
)


func checkPassp(rawPp []string) (valid bool) {
	var ppm PassportMask	
	for _, str := range rawPp {
		strArr := strings.Split(str, " ")
		for _, ppField := range strArr {
			fields := strings.Split(ppField, ":")
			switch fields[0] {
				case "byr": ppm = ppm | byr
				case "iyr": ppm = ppm | iyr 
				case "eyr": ppm = ppm | eyr 
				case "hgt": ppm = ppm | hgt
				case "hcl": ppm = ppm | hcl
				case "ecl": ppm = ppm | ecl
				case "pid": ppm = ppm | pid 
				case "cid": ppm = ppm | cid
			}
			println("fields[0]", fields[0], "ppm:", ppm)
		}
	}
	return (ppm == 255 || ppm == 127)
}

func main() {
	r, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Can't open file arg")	
	}	
	scnr := bufio.NewScanner(r)
	//scnr.Whitespace ^= 1 << "\r" | 1 << " "	
	var rawPassp []string	
	var valid int	
	for scnr.Scan() {
		ln := scnr.Text()
		if len(ln) > 0 {
			rawPassp = append(rawPassp, ln)
		} else {
			if checkPassp(rawPassp) {
				valid++
			}
			rawPassp = nil
		}
	}
	if rawPassp != nil {
		if checkPassp(rawPassp) {
			valid++
		}
	}
	fmt.Println("Valid:", valid)
}

