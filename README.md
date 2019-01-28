[![Go Report Card](https://goreportcard.com/badge/github.com/pladdy/stones)](https://goreportcard.com/report/github.com/pladdy/stones)
[![Code Coverage](https://codecov.io/gh/pladdy/stones/branch/master/graph/badge.svg)](https://codecov.io/gh/pladdy/stones)
[![Build Status](https://travis-ci.org/pladdy/stones.svg?branch=master)](https://travis-ci.org/pladdy/stones)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/pladdy/stones)
[![Go City](https://img.shields.io/badge/go--city-view-blue.svg)](https://go-city.github.io/#/github.com/pladdy/stones)
[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/pladdy/stones/releases/latest)

## stones
A STIX 2.0 Validator and Creator written in Go

This is being written to separate STIX validation/processing from TAXII (github.com/pladdy/cabby).

## Use cases
- [x] Serialize/Deserialize Bundles from JSON
- [x] Serialize/Deserialize Objects from JSON
- [x] Verify Bundles are valid
- [x] Verify Objects are valid
- [ ] Validate objects based on their type
- [ ] Return a generic object that has the attributes of its type?

## Examples
- https://github.com/pladdy/stones/examples/unmarshal_bundle.go

## Resources
- OASIS Doc: https://oasis-open.github.io/cti-documentation/resources
- STIX 2.0 Specs:
  - http://docs.oasis-open.org/cti/stix/v2.0/stix-v2.0-part1-stix-core.html
  - http://docs.oasis-open.org/cti/stix/v2.0/stix-v2.0-part2-stix-objects.html
