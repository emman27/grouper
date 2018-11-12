package grouper

import (
	"fmt"
	"math/rand"
	"time"
)

// Group is a grouping of people
type Group struct {
	Name    string
	Members []Person
	MaxSize int
}

// Groups is a collection of groups
type Groups []*Group

// NewGroup creates a new group
func NewGroup(name string, maxSize int) *Group {
	return &Group{
		Name:    name,
		Members: []Person{},
		MaxSize: maxSize,
	}
}

// Size returns the size of a group
func (g *Group) Size() int {
	return len(g.Members)
}

// IsFull tells you if a group is already full. A full group is not a valid group
func (g *Group) IsFull() bool {
	return g.Size() >= g.MaxSize
}

// AddMember adds a person to a group
func (g *Group) AddMember(p *Person) *Group {
	g.Members = append(g.Members, *p)
	return g
}

func (g *Group) String() string {
	baseText := fmt.Sprintf("%s with Group members:", g.Name)
	for _, member := range g.Members {
		baseText += fmt.Sprintf(" %s,", member.Name)
	}
	baseText += fmt.Sprintf("\nMax Size: %d, Current Size: %d", g.MaxSize, g.Size())
	return baseText
}

// GetValidGroups returns all groups in a set that can still have more people in them
func (g *Groups) GetValidGroups() Groups {
	valid := Groups{}
	for _, group := range *g {
		if !group.IsFull() {
			valid = append(valid, group)
		}
	}
	return valid
}

// Random selects a random group from the set of groups
func (g *Groups) Random() *Group {
	rand.Seed(time.Now().UnixNano())
	if len(*g) == 0 {
		panic("No groups to select from!")
	}
	i := rand.Intn(len(*g))
	return (*g)[i]
}
