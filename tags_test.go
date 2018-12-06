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
	Name   string `json:"name" select:"persons.[0].name"`
	Age    int    `json:"age" select:"persons.[0].age"`
	NotAge int
}

func TestSelect(t *testing.T) {
	ty := reflect.TypeOf(Person{})
	val, err := Select(ty, personData)
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
	Name  string `json:"name" select:"cars.[0].name"`
	Brand string `json:"brand" select:"cars.[0].brand"`
	Gears int    `json:"gears" select:"cars.[1].gears"`
}

const carData = `
{
	"cars": [
		{ "name": "R8", "brand": "Audi"}
	]
}
`

func TestSelectFail(t *testing.T) {
	ty := reflect.TypeOf(Car{})
	_, err := Select(ty, carData)

	if err == nil {
		t.Fail()
	}
}
