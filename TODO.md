### valid stix type getter/method?
- map of types to boolean...what if it was map to string and bool is in there?
```
var validStixTypes = map[string]bool{
	attackPatternType:  true,
  ???
```

### should stones return a specific object?
- should a caller have to also call Valid() on a thing?
  - should NewObject only return an object if it's valid?  isn't that the point?
	- if i unmarshal an object from json, unmarshal doesn't call Valid...but it could
