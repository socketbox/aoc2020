package main

import ( 
	"os"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
	"regexp"
	"errors"
)

var passwdRegex = regexp.MustCompile(`^([0-9]{1,3})-([0-9]{1,3})\s{1}([a-z]{1}):\s{1}([a-z]*)$`)

func IsValid(matches []string) (err error, valid bool) {
	if len(matches) < 1 {
		return errors.New("No matches in argument"), false
	}
	f, l, letter, passwd := matches[1], matches[2], matches[3], matches[4]
	first, _ := strconv.Atoi(f)
	second, _ := strconv.Atoi(l)
	var fBool, sBool bool
	/* while letter is still found in passwd, index under passwd length, and letter not at either
	index */
	for idx := 0; idx > -1 && idx <= len(passwd) && !valid; 
	{
		idx = strings.Index(passwd, letter)
		if idx == -1 { 	
			break;	
		}
		if idx + 1 == first {//  &! (idx + 1 == first == second) {
			fBool = true
		} else if idx + 1 == second {
			sBool = true
		}
		passwd = strings.Replace(passwd, letter, "_", 1)	
	}
	return err, (fBool || sBool) && !(fBool == sBool)
}

func main() {
	fileName := os.Args[1]
	bs, err := ioutil.ReadFile(fileName) 
	if err != nil {
		//nop	
	}
	passwdArr := strings.Split(string(bs), "\n")
	var matches []string = nil
	var val int;
	for _, p := range passwdArr {
		matches = passwdRegex.FindStringSubmatch(p)
		err, valid := IsValid(matches)
		if err != nil {
			//nop
		}
		if valid { 
			fmt.Println("valid passwd: ", p)
			val++ 
		}
	}
	fmt.Println("Valid ", val)
}
		

