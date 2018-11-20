### valid stix type getter/method?
- map of types to boolean...what if it was map to type string and bool is in there?
```
var validStixTypes = map[string]bool{
	attackPatternType:  true,
  ???
```

### taxii2 server uses stones to validate an object
- should Object have a validate fn?  
- how do i receive an object, validate it based on type and return true/false?
