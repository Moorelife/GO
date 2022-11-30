// system
package system

import (
	"fmt"
	"os"
)

type System struct {
	Name string
}

func NewSystem(name string) *System {
	newSystem := System{name}
	return &newSystem
}

func (s *System) GetName() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Println("hostname:", name)
}
