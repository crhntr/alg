package stablemarriage

import (
	"bytes"
	_ "embed"
	"testing"
)

var (
	//go:embed testdata/people00.json
	people00JSON []byte

	//go:embed testdata/people01.json
	people01JSON []byte

	//go:embed testdata/people02.json
	people02JSON []byte

	//go:embed testdata/badData00.json
	badData00JSON []byte

	//go:embed testdata/badData01.json
	badData01JSON []byte

	//go:embed testdata/badData02.json
	badData02JSON []byte
)

func TestAlg00(t *testing.T) {
	people, err := LoadPeople(bytes.NewReader(people00JSON))
	if err != nil {
		t.Fatal(err)
	}
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
	people, err := LoadPeople(bytes.NewReader(people01JSON))
	if err != nil {
		t.Fatal(err)
	}
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
	people, err := LoadPeople(bytes.NewBuffer(people02JSON))
	if err != nil {
		t.Fatal(err)
	}
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

func TestLoadPeopleFromFile(t *testing.T) {
	if _, err := LoadPeople(bytes.NewBuffer(badData00JSON)); err == nil {
		t.Error("testdata/badData00.json should err")
	}
	if _, err := LoadPeople(bytes.NewBuffer(badData01JSON)); err == nil {
		t.Error("testdata/badData01.json should err")
	}
	if _, err := LoadPeople(bytes.NewBuffer(badData02JSON)); err == nil {
		t.Error("testdata/badData02.json should err")
	}
}
