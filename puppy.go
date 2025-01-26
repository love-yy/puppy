package puppy

import (
	"fmt"

	"github.com/love-yy/dog"
)

func Bark() string {
	return "Woof!🐕‍🦺 "
}

func Barks() string {
	return "Woof! Woof! Woof!🐩 "
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func From11() {
	fmt.Println("I'm from version v1.1.0")
}

func From12() {
	fmt.Println("I'm from version v1.2.0")
}
