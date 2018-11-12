package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/emman27/grouper"
)

var numOfGroups = flag.Int("num-groups", 5, "The number of groups to split people into. Groups will be split somewhat evenly")

var peopleFilePath = flag.String("people-file", "people.txt", "The filepath to the newline separated list of names")

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	people := grouper.GetPeopleFromFile(peopleFilePath)
	groups := grouper.Groups{}
	groupMaxSize := int(math.Ceil(float64(len(people)) / float64(*numOfGroups)))

	for i := 0; i < *numOfGroups; i++ {
		groups = append(groups, grouper.NewGroup(fmt.Sprintf("Group %d", i), groupMaxSize))
	}

	grouper.AssignGroups(&people, &groups)

	for _, group := range groups {
		fmt.Println(group)
	}
}
