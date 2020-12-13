package main

import (
	"os"
	"io"	
	"io/ioutil"
	"fmt"
	"bufio"
	//"bytes"
)

type Passport struct {
	byr		int
	iyr		int
	eyr		int
	hgt		string
	hcl		string
	ecl		string
	pid		string
	cid		string
}

func parsePassports(bs []byte) (passPorts []Passport) {
	var adv int
	var textBs []byte
	var err error
	var rawPport [][]byte
	//var rawPport = make( [][]byte, 4, 8) 
	for i := 0; err == nil && i < 250 && i < len(bs); {
		adv, textBs, err = bufio.ScanLines(bs[i:], true)
		if len(textBs) > 0 { 
			rawPport = append(rawPport, textBs)	
			fmt.Println("rawPport", i, ":", rawPport)	
		} else {
			fmt.Println("to makePassport")	
			var pPort Passport = makePassport(rawPport)
			passPorts = append(passPorts, pPort)	
			rawPport = nil
		}
		i += adv
	}
	return passPorts
}

func makePassport(rawPp [][]byte) (pp Passport) {
	var rd *Reader = new Reader()
	var err error 
	for i := 0; i < len(rawPp); i++ {
		fmt.Println("rawPp len:", len(rawPp))
		rz := bufio.NewReaderSize(rd, len(rawPp[i]))
		err = nil
		var field []byte
		for ; err == nil; {
			field, err = rz.ReadBytes(' ')
			fmt.Println(string(field))
		}
	}
	return pp 
}

func main() {
	fileName := os.Args[1]
	bs, err := ioutil.ReadFile(fileName) 
	if err != nil {
		//nop	
	}
	var passPorts []Passport = parsePassports(bs)
	fmt.Println(&passPorts)
}
