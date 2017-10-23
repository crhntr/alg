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

func (person Person) IsSingle() bool {
	return person.EngagedTo == nil
}

func (person *Person) BreakUp() {
	person.EngagedTo.EngagedTo = nil
	person.EngagedTo = nil
}

// Perfers returns
//  n > 1 if person1 a is prefered
//  n = 0 if there is no preference
//  n < 1 if person2 a is prefered
func (person Person) PrefersToExistingEngagement(otherPerson *Person) bool {
	otherPersonIndex := len(person.Preferences)
	engagedToIndex := len(person.Preferences)

	for i, prefName := range person.Preferences {
		if prefName == otherPerson.Name {
			otherPersonIndex = i
		}
		if prefName == person.EngagedTo.Name {
			engagedToIndex = i
		}
	}

	return engagedToIndex > otherPersonIndex
}

func (person *Person) GetPrefered(people []Person) int {
	i := findIndex(people, person.Preferences[0])
	if len(person.Preferences) > 1 {
		person.Preferences = person.Preferences[1:]
	}
	return i
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

func propose(person1, person2 *Person) {
	person1.EngagedTo, person2.EngagedTo = person2, person1
}

func findIndex(people []Person, name string) int {
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
