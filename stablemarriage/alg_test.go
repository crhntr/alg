package stablemarriage

import (
	"encoding/json"
	"os"
	"testing"
)

func TestAlg00(t *testing.T) {
	people := loadJSONPeople(t, "testdata/set00.json")
	males, females := SortPeopleByGender(people)
	Alg(males, females)

	if SinglesExist(males) || SinglesExist(females) {
		t.Fail()
	}
	// for _, person := range males {
	// 	t.Log(person)
	// }
}

func TestAlg01(t *testing.T) {
	people := loadJSONPeople(t, "testdata/set01.json")
	males, females := SortPeopleByGender(people)
	Alg(males, females)

	if SinglesExist(males) || SinglesExist(females) {
		t.Fail()
	}
	// for _, person := range males {
	// 	t.Log(person)
	// }
}

func TestAlg02(t *testing.T) {
	people := loadJSONPeople(t, "testdata/set02.json")
	males, females := SortPeopleByGender(people)

	for _, person := range append(males, females...) {
		t.Log(person)
	}

	Alg(males, females)

	if SinglesExist(males) || SinglesExist(females) {
		t.Fail()
	}
	for _, person := range males {
		t.Log(person)
	}
}

func loadJSONPeople(t *testing.T, filename string) (people []Person) {
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal(err)
	}
	file := struct {
		People []Person `json:"people"`
	}{}
	err = json.NewDecoder(f).Decode(&file)
	if err != nil {
		t.Fatal(err)
	}
	return file.People
}
