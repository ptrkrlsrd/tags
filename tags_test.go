package tags

import (
	"encoding/json"
	"reflect"
	"testing"
)

const personData string = `
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

const carData string = `
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

func BenchmarkSelect100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ty := reflect.TypeOf(Person{})
		val, _ := Select(ty, personData)
		_ = val.Interface().(*Person)
	}
}

func BenchmarkNoSelect100(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var persons struct {
			Persons []Person
		}
		json.Unmarshal([]byte(personData), &persons)
		person := persons.Persons[0]
		_ = person
	}
}
