package stablemarriage

import (
	"encoding/json"
	"os"
	"strings"
)

type Person struct {
	Name        string   `json:"name"`
	Male        bool     `json:"male"`
	Preferences []string `json:"preferences"`

	EngagedTo *Person
}

// Perfers returns
//  n > 1 if person1 a is prefered
//  n = 0 if there is no preference
//  n < 1 if person2 a is prefered
func (person Person) Perfers(person1, person2 *Person) int {
	p1Index := len(person.Preferences)
	p2Index := len(person.Preferences)

	for i, prefName := range person.Preferences {
		if prefName == person1.Name {
			p1Index = i
		}
		if prefName == person2.Name {
			p2Index = i
		}
	}

	return p2Index - p1Index
}

func SinglesExist(people []Person) bool {
	for i := range people {
		if people[i].EngagedTo == nil {
			return true
		}
	}
	return false
}

// func HaveUntestedOptions(people []Person) bool {
// 	for i := range people {
// 		if len(people[i].Preferences) > 0 {
// 			return true
// 		}
// 	}
// 	return false
// }

func Propose(person1, person2 *Person) {
	person1.EngagedTo, person2.EngagedTo = person2, person1
}

func FindIndex(people []Person, name string) int {
	for i := range people {
		if people[i].Name == name {
			return i
		}
	}
	panic(name + " not found")
}

func SortPeopleByGender(people []Person) (males, females []Person) {
	for _, person := range people {
		if person.Male {
			males = append(males, person)
		} else {
			females = append(females, person)
		}
	}
	return
}

func LoadPeopleJSON(filename string) (people []Person) {
	f, err := os.Open(filename)
	mustNotErr(err)
	file := struct {
		People []Person `json:"people"`
	}{}
	err = json.NewDecoder(f).Decode(&file)
	mustNotErr(err)
	return file.People
}

func mustNotErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (person Person) String() string {
	str := strings.Title(person.Name)
	if person.EngagedTo == nil {
		str += " (single)"
		return str
	}
	str += " is engaged to "
	str += strings.Title(person.EngagedTo.Name)
	return str
}
