What if I use a tool for validation? https://github.com/thedevsaddam/govalidator
- validates UUIDs
- JSON
- etc.

Need to add ability to add custom properties to objects
- eg:
```js
{
	// snip
	"x_pladdy_custom1": "pants"
}
```

### valid stix type getter/method?
- map of types to boolean...what if it was map to string and bool is in there?
Ex:
```
var validStixTypes = map[string]bool{
	attackPatternType:  true,
  // ???
```

### Validation Goals
take a JSON object, pass it to stones and see if it's valid
- stones.Validate(rawJSON)
- regardless of the object, run validation for generic objects and the specific type

### Generation Goals
- can create a new generic object given a type: `stones.NewObject("malware")`
- can create a specific object: `stones.NewAttackPattern(// signature`
- can unmarshal a specific object:
```go
var ap stones.AttackPattern
err := json.Unmarshal(someRawJSON, &ap)
```
- each object can be validated
  - given 'ap' is an attack pattern object: `ap.Valid()`
- can add custom properties?
  - how though?  go is strict typing
	- maybe a doc on wrapping and implementing your own for custom?
		- example:
			```go
			package main

		  import stones

			type AttackPattern {
				stones.AttackPattern
				FooField string
				//... etc.
			}
			```

### JSON libraries and ideas for marshalling/unmarshalling
https://github.com/json-iterator/go
- https://godoc.org/github.com/json-iterator/go

https://medium.com/@irshadhasmat/golang-simple-json-parsing-using-empty-interface-and-without-struct-in-go-language-e56d0e69968
https://github.com/valyala/fastjson
- i don't think this is a good fit
https://github.com/tidwall/gjson
- flexible, might be good for handling custom json fields
