package paging

import "fmt"

type Computer struct {
	ID   int
	Name string
}

func (c Computer) String() string {
	return fmt.Sprintf("id: %3d | name: %s", c.ID, c.Name)
}