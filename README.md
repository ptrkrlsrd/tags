# tags
## A Go package for transforming JSON-data from tags
[![Build Status](https://travis-ci.com/ptrkrlsrd/tags.svg?token=EC6EZTgzr1WN8mybj2yE&branch=master)](https://travis-ci.com/ptrkrlsrd/tags)


## Usage

### Example based on the tests
``` go
package main

type Car struct {
	Name  string `json:"name" select:"cars.[0].name"`
	Brand string `json:"brand" select:"cars.[0].brand"`
	Gears int    `json:"gears" select:"cars.[0].gears"`
}

const carData = `
{
	"cars": [
		{ "name": "R8", "brand": "Audi", "gears": 5}
	]
}
`

func main() {
	ty := reflect.TypeOf(Car{})
	carValue, err := Select(ty, personData)
  car := carValue.Interface().(*Car)
}
```
