package grouper

import (
	"reflect"
	"testing"
)

func TestGroup_Size(t *testing.T) {
	type fields struct {
		Name    string
		Members []Person
		MaxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{fields: fields{Name: "", Members: []Person{}, MaxSize: 1}, want: 0},
		{fields: fields{Name: "", Members: []Person{
			Person{Name: "Emmanuel"},
		}, MaxSize: 1}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Name:    tt.fields.Name,
				Members: tt.fields.Members,
				MaxSize: tt.fields.MaxSize,
			}
			if got := g.Size(); got != tt.want {
				t.Errorf("Group.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_IsFull(t *testing.T) {
	type fields struct {
		Name    string
		Members []Person
		MaxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{fields: fields{Name: "blah", Members: []Person{Person{Name: "Emmanuel"}}, MaxSize: 0}, want: true},
		{fields: fields{Name: "blah", Members: []Person{Person{Name: "Emmanuel"}}, MaxSize: 1}, want: true},
		{fields: fields{Name: "blah", Members: []Person{Person{Name: "Emmanuel"}}, MaxSize: 2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Name:    tt.fields.Name,
				Members: tt.fields.Members,
				MaxSize: tt.fields.MaxSize,
			}
			if got := g.IsFull(); got != tt.want {
				t.Errorf("Group.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_AddMember(t *testing.T) {
	type fields struct {
		Name    string
		Members []Person
		MaxSize int
	}
	type args struct {
		p *Person
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Group
	}{
		{
			fields: fields{Name: "My Group", Members: []Person{}},
			args:   args{p: &Person{Name: "Emmanuel"}},
			want: &Group{Name: "My Group", Members: []Person{
				Person{Name: "Emmanuel"},
			}},
		},
		{
			fields: fields{Name: "My Group", Members: []Person{
				Person{Name: "Emmanuel"},
			}},
			args: args{p: &Person{Name: "Goh"}},
			want: &Group{Name: "My Group", Members: []Person{
				Person{Name: "Emmanuel"},
				Person{Name: "Goh"},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Name:    tt.fields.Name,
				Members: tt.fields.Members,
				MaxSize: tt.fields.MaxSize,
			}
			if got := g.AddMember(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Group.AddMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_String(t *testing.T) {
	type fields struct {
		Name    string
		Members []Person
		MaxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			fields: fields{Name: "Group A", Members: []Person{}, MaxSize: 5},
			want:   "Group A with Group members:\nMax Size: 5, Current Size: 0",
		},
		{
			fields: fields{Name: "Group A", Members: []Person{Person{Name: "Emmanuel"}}, MaxSize: 5},
			want:   "Group A with Group members: Emmanuel,\nMax Size: 5, Current Size: 1",
		},
		{
			fields: fields{Name: "Group A", Members: []Person{Person{Name: "Emmanuel"}, Person{Name: "Goh"}}, MaxSize: 5},
			want:   "Group A with Group members: Emmanuel, Goh,\nMax Size: 5, Current Size: 2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Group{
				Name:    tt.fields.Name,
				Members: tt.fields.Members,
				MaxSize: tt.fields.MaxSize,
			}
			if got := g.String(); got != tt.want {
				t.Errorf("Group.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
