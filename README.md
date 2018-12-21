[![Go Report Card](https://goreportcard.com/badge/github.com/pladdy/stones)](https://goreportcard.com/report/github.com/pladdy/stones)
[![Code Coverage](https://codecov.io/gh/pladdy/stones/branch/master/graph/badge.svg)](https://codecov.io/gh/pladdy/stones)
[![Build Status](https://travis-ci.org/pladdy/stones.svg?branch=master)](https://travis-ci.org/pladdy/stones)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/pladdy/stones)
[![Go City](https://img.shields.io/badge/go--city-view-blue.svg)](https://go-city.github.io/#/github.com/pladdy/stones)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/pladdy/stones/releases/latest)

## stones
A STIX 2.0 Validator written in Go

This is being written to separate STIX validation/processing from TAXII (github.com/pladdy/cabby).

## Use cases
- [x] Serialize/Deserialize Bundles from JSON
- [x] Serialize/Deserialize Objects from JSON
- [x] Verify Bundles are valid
- [x] Verify Objects are valid
- [ ] Validate objects based on their type
- [ ] Return a generic object that has the attributes of its type?

## How to Use
```go
import (
  "fmt"

  "github.com/pladdy/stones"
)

rawJson := []byte(`{
              "type": "malware",
              "id": "malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b",
              "created": "2016-04-06T20:07:09.000Z",
              "modified": "2016-04-06T20:07:09.000Z",
              "created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
              "name": "Poison Ivy"
            }`)

var o stones.Object
err := stones.UnmarshalJSON(rawJson, o)
if err != nil {
  fmt.Println("something went wrong, is it a valid object")
}

fmt.Println(o)
```

## Resources
- OASIS Doc: https://oasis-open.github.io/cti-documentation/resources
- STIX 2.0 Specs:
  - http://docs.oasis-open.org/cti/stix/v2.0/stix-v2.0-part1-stix-core.html
  - http://docs.oasis-open.org/cti/stix/v2.0/stix-v2.0-part2-stix-objects.html
