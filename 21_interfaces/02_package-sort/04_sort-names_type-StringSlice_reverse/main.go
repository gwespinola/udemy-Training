package main

import (
	"fmt"
	"sort"
)

type people struct {
	Names []string
}

func (studyGroup people) members() []string {
	return fmt.Println(studyGroup)

}

type gang interface {
	members() []string
}

func reverse (studyGroup people) {
	fmt.Println(studyGroup)
	fmt.Println(studyGroup.members())
	sort.Reverse(gang(studyGroup))
}

func main() {
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
}
