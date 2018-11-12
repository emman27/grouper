package grouper

import (
	"github.com/golang/glog"
)

// AssignGroups assigns groups to every person, respecting the rules of the group
func AssignGroups(people *[]Person, groups *Groups) {
	for _, person := range *people {
		availableGroups := groups.GetValidGroups()
		groupToAssignTo := availableGroups.Random()
		glog.Infof("Assigning %s to %s", person.Name, groupToAssignTo.Name)
		groupToAssignTo.AddMember(&person)
	}
}
