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
