package stablemarriage

import (
	"testing"
)

func TestStableMatching00(t *testing.T) {
	people := LoadPeopleJSON("testdata/set00.json")
	males, females := SortPeopleByGender(people)
	StableMatching(males, females)

	for _, person := range males {
		t.Log(person)
	}
}

func TestStableMatching01(t *testing.T) {
	people := LoadPeopleJSON("testdata/set01.json")
	males, females := SortPeopleByGender(people)
	StableMatching(males, females)

	for _, person := range males {
		t.Log(person)
	}
}

func TestStableMatching02(t *testing.T) {
	people := LoadPeopleJSON("testdata/set02.json")
	males, females := SortPeopleByGender(people)

	for _, person := range append(males, females...) {
		t.Log(person)
	}

	StableMatching(males, females)

	for _, person := range males {
		t.Log(person)
	}
}
