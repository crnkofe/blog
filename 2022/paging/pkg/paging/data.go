package paging

import "fmt"

type Computer struct {
	ID   int    `sql:"id"`
	Name string `sql:"name"`
}

func (c Computer) String() string {
	return fmt.Sprintf("id: %3d | name: %s", c.ID, c.Name)
}