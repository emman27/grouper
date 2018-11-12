package grouper

import (
	"io/ioutil"
	"strings"
)

type Person struct {
	Name string
}

// GetPeopleFromFile reads a file from the filesystem.
// The file should contain a list of names, one on each line
func GetPeopleFromFile(filepath string) []Person {
	people := []Person{}
	dat, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic("The specified file does not exist")
	}
	commaSeparated := strings.Split(string(dat), "\n")
	for _, name := range commaSeparated {
		if name != "" {
			people = append(people, Person{Name: name})
		}
	}
	return people
}
