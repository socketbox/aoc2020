package main

import (
	//"reflect"
	"bufio"
	"strings"	
	"fmt"
	"log"	
	"os"
	"strconv"
	"regexp"
)

var yrRgx = regexp.MustCompile(`^[0-9]{4}$`)
var hgtRgx = regexp.MustCompile(`([0-9]{2,3})(cm|in)`)
var hclRgx = regexp.MustCompile(`#[0-9a-f]{6}`)
var eclRgx = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
var pidRgx = regexp.MustCompile(`^[0-9]{9}$`)

type PassportMask uint8 

type Passport struct {
	byr		int
	iyr		int
	eyr		int
	hgt		string
	hcl		string
	ecl		string
	pid		string
	cid		string
	mask	PassportMask
}

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

func checkPassp(rawPp []string) (valid bool, pp Passport) {
	var err error = nil	
	for _, str := range rawPp {
		strArr := strings.Split(str, " ")
		for _, ppField := range strArr {
			fields := strings.Split(ppField, ":")
			switch fields[0] {
				case "byr": {
					pp.mask = pp.mask | byr
					pp.byr, err = strconv.Atoi(fields[1])
				}
				case "iyr": {
					pp.mask = pp.mask | iyr 
					pp.iyr, err = strconv.Atoi(fields[1])
				}
				case "eyr": {
					pp.mask = pp.mask | eyr
					pp.eyr, err = strconv.Atoi(fields[1])
				} 
				case "hgt": {
					pp.mask = pp.mask | hgt
					pp.hgt = fields[1]
				}
				case "hcl": {
					pp.mask = pp.mask | hcl
					pp.hcl = fields[1]
				}
				case "ecl": {
					pp.mask = pp.mask | ecl
					pp.ecl = fields[1]
				}
				case "pid": {
					if strings.Index(fields[1], "175631967") != -1 {
						fmt.Println("Found:", fields[1])
					}
					pp.mask = pp.mask | pid
					pp.pid = fields[1]
				} 
				case "cid": {
					pp.mask = pp.mask | cid
					pp.cid = fields[1]
				}
				default: panic("Invalid passport field.")
			}
			if err != nil {
				log.Output(1, "Error in str conversion")
			}
		}
	}
	return (pp.mask == 255 || pp.mask == 127), pp
}

func heightIsValid(hgt string) (valid bool) {
	valid = true	
	matches := hgtRgx.FindStringSubmatch(hgt)
	if len(matches) != 3 {
		valid = false
	} else if matches[2] == "cm" {
		h, err := strconv.Atoi(matches[1])
		if h < 150 || h > 193 {
			valid = false
		}
		if err != nil {
			log.Output(1, "Bad conv to int")
		}
	} else if matches[2] == "in" {
		h, err := strconv.Atoi(matches[1])
		if h < 59 || h > 76 {
			valid = false
		}
		if err != nil {
			log.Output(1, "Bad conv to int")
		}
	} else {
		valid = false
	}
	return valid
}

func validatePassp(pp Passport) (valid bool) {
	valid = true
	if pp.byr < 1920 || pp.byr > 2002 || !yrRgx.MatchString( strconv.Itoa(pp.byr)) {
		valid = false
	} else if pp.iyr < 2010 || pp.iyr > 2020 || !yrRgx.MatchString( strconv.Itoa(pp.iyr)) { 
		valid = false
	} else if pp.eyr < 2020 || pp.eyr > 2030 || !yrRgx.MatchString( strconv.Itoa(pp.eyr)) { 
		valid = false
	} else if !heightIsValid(pp.hgt) {
		valid = false
	} else if !hclRgx.MatchString(pp.hcl) {
		valid = false
	} else if !eclRgx.MatchString(pp.ecl) {
		valid = false
	} else if !pidRgx.MatchString(pp.pid) {
		valid = false
	}
	if valid {
		fmt.Println(pp)
	}
	return valid
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
			v, pp := checkPassp(rawPassp)
			if v {
				if validatePassp(pp) {
					valid++
				}
			}
			rawPassp = nil
		}
	}
	if rawPassp != nil {
		v, pp := checkPassp(rawPassp)
		if v {
			if validatePassp(pp) {
				valid++
			}
		}
	}
	fmt.Println("Valid:", valid)
}

