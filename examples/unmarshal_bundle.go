package main

import (
	"encoding/json"
	"fmt"

	"github.com/pladdy/stones"
)

func main() {
	rawJSON := []byte(`{
                      "type": "bundle",
                      "id": "bundle--5d0092c5-5f74-4287-9642-33f4c354e56d",
                      "spec_version": "2.0",
                      "objects": [
                        {
                          "type": "indicator",
                          "id": "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
                          "created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
                          "created": "2016-04-06T20:03:48.000Z",
                          "modified": "2016-04-06T20:03:48.000Z",
                          "labels": ["malicious-activity"],
                          "name": "Poison Ivy Malware",
                          "description": "This file is part of Poison Ivy",
                          "pattern": "[ file:hashes.'SHA-256' = '4bac27393bdd9777ce02453256c5577cd02275510b2227f473d03f533924f877' ]",
                          "valid_from": "2016-01-01T00:00:00Z"
                        },
                        {
                          "type": "relationship",
                          "id": "relationship--44298a74-ba52-4f0c-87a3-1824e67d7fad",
                          "created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
                          "created": "2016-04-06T20:06:37.000Z",
                          "modified": "2016-04-06T20:06:37.000Z",
                          "relationship_type": "indicates",
                          "source_ref": "indicator--8e2e2d2b-17d4-4cbf-938f-98ee46b3cd3f",
                          "target_ref": "malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b"
                        },
                        {
                          "type": "malware",
                          "id": "malware--31b940d4-6f7f-459a-80ea-9c1f17b5891b",
                          "created": "2016-04-06T20:07:09.000Z",
                          "modified": "2016-04-06T20:07:09.000Z",
                          "created_by_ref": "identity--f431f809-377b-45e0-aa1c-6a4751cae5ff",
                          "name": "Poison Ivy"
                        }
                      ]
                    }`)

	var b stones.Bundle
	err := json.Unmarshal(rawJSON, &b)
	if err != nil {
		fmt.Println("something went wrong, is it a valid bundle?")
	}

	// print the object
	fmt.Println("type:", b.Type)
	fmt.Println("id:", b.ID.String())
	fmt.Println("spec_version:", b.SpecVersion)
	fmt.Println("first object:", string(b.Objects[0]))

	// print results of validation
	fmt.Print("It should be valid ('true []'): ")
	fmt.Println(b.Valid())
}
