### valid stix type getter/method?
- map of types to boolean...what if it was map to string and bool is in there?
```go
var validStixTypes = map[string]bool{
	attackPatternType:  true,
  ???
```

### should stones return a specific object?
- should a caller have to also call Valid() on a thing?
  - should NewObject only return an object if it's valid?  isn't that the point?
	- if i unmarshal an object from json, unmarshal doesn't call Valid...but it could

Options:
- each object has its own type with an embeddded object
  - since this is strongly typed language, you have to know what the object is before hand...this seems gross
- Object is a generic type...with specific types embedded?

```go
type Object {
	# the fields already defined
	AttackPattern
	Malware
	# etc.
	# the specific types will be nil unless the Type field is set to that type...or is that dumb?
}
```
