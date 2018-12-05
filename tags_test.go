package tags

import (
	"reflect"
	"testing"
)

const personData = `
{
	"persons": [
		{ "name": "Person", "age": 10}
	]
}
`

type Person struct {
	Name   string `json:"name" tags:"persons.[0].name"`
	Age    int    `json:"age" tags:"persons.[0].age"`
	NotAge int
}

func TestParse(t *testing.T) {
	ty := reflect.TypeOf(Person{})
	val, err := Parse(ty, personData)
	person := val.Interface().(*Person)

	if err != nil {
		t.Fail()
	}

	if person.Age != 10 {
		t.Fail()
	}

	if person.Name != "Person" {
		t.Fail()
	}
}

type Car struct {
	Name  string `json:"name" tags:"cars.[0].name"`
	Brand string `json:"brand" tags:"cars.[0].brand"`
	Gears int    `json:"gears" tags:"cars.[1].gears"`
}

const carData = `
{
	"cars": [
		{ "name": "R8", "brand": "Audi"}
	]
}
`

func TestParseFail(t *testing.T) {
	ty := reflect.TypeOf(Car{})
	_, err := Parse(ty, personData)

	if err == nil {
		t.Fail()
	}
}
