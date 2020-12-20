package main

import (
	"bufio"
	"bytes"	
	"fmt"
	//"io"
	"io/ioutil"
	"os"
	"strings"
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
/*
type Passport struct {
	byr		int
	iyr		int
	eyr		int
	hgt		string
	hcl		string
	ecl		string
	pid		string
	cid		string
}*/

func parsePassports(bs []byte) (passPorts []PassportMask) {
	var adv int
	var textBs []byte
	var err error
	var rawPport [][]byte
	//var rawPport = make( [][]byte, 4, 8) 
	var ppm PassportMask 
	for i := 0; err == nil && i < len(bs); {
		adv, textBs, err = bufio.ScanLines(bs[i:], true)
		if len(textBs) > 0 { 
			rawPport = append(rawPport, textBs)	
			fmt.Println("rawPport:", rawPport)	
		} else {
			fmt.Println("to makePassport")	
			ppm = makePassport(rawPport)
			fmt.Println("ppm:", ppm)
			passPorts = append(passPorts, ppm)
			fmt.Println("passPorts:", len(passPorts))
			rawPport = nil
		}
		i += adv
	}
	return passPorts
}

func makePassport(rawPp [][]byte) (pp PassportMask) {
	var err error 
	var rd *bytes.Reader
	var rz *bufio.Reader
	for i := 0; i < len(rawPp); i++ {
		fmt.Println("rawPp len:", len(rawPp[i]))
		rd = bytes.NewReader(rawPp[i])	
		rz = bufio.NewReaderSize(rd, len(rawPp[i]))
		err = nil
		var field []byte
		for err == nil {
			field, err = rz.ReadBytes(' ')
			fields := strings.Split(string(field), ":")
			fmt.Println(fields[0])
			switch fields[0] {
				case "byr": pp = pp | byr
				case "iyr": pp = pp | iyr 
				case "eyr": pp = pp | eyr 
				case "hgt": pp = pp | hgt
				case "hcl": pp = pp | hcl
				case "ecl": pp = pp | ecl
				case "pid": pp = pp | pid 
				case "cid": pp = pp | cid
			}
			println("fields[0]", fields[0], "pp:", pp)
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
	
	var passPorts []PassportMask = parsePassports(bs)
	
	val := 0
	for _, p := range passPorts {
		//fmt.Println("p:", p, "i:", i)	
		if p == 255 || p == 127 {
			val++
		}
	}
	fmt.Println("valid:", val)
}
