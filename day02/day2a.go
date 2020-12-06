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
	l, h, letter, passwd := matches[1], matches[2], matches[3], matches[4]
	low, _ := strconv.Atoi(l) 
	high, _ := strconv.Atoi(h) 
	var count = strings.Count(passwd, letter)
	if low <= count && count <= high {
		valid = true
	}
	return err, valid
}

func main() {
	fileName := os.Args[1]
	bs, err := ioutil.ReadFile(fileName) 
	if err != nil {
		//pass	
	}
	passwdArr := strings.Split(string(bs), "\n")
	var matches []string = nil
	var val int;
	for i, p := range passwdArr {
		matches = passwdRegex.FindStringSubmatch(p)
		err, valid := IsValid(matches)
		if err != nil {
			//pass
		}
		if valid { val++ }
	}
	fmt.Println("Valid ", val)
}
		

