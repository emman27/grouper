package grouper

import (
	"reflect"
	"testing"
)

func TestGetPeopleFromFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want []Person
	}{
		{args: args{filepath: "./people.txt"}, want: []Person{
			{Name: "Emmanuel Goh"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPeopleFromFile(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPeopleFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
