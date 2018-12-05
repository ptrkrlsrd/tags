# tags
## A Go package for transforming JSON-data on input


### Concept API
``` go
package main

const data = `
  {
    "persons": [
      { "name": "hello" },
    ]
  }
`


type Person struct {
  Name string `json:"name",tags:"persons.0?.name"`
}
```
