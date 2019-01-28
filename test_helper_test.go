package stones

import (
	"io/ioutil"
	"log"
)

func attackPatternJSON() []byte {
	return readFile("testdata/attack_pattern.json")
}

func readFile(filePath string) []byte {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v, error: %v\n", filePath, err)
	}
	return b
}

func malwareJSON() []byte {
	return readFile("testdata/malware.json")
}

func validTestObject(objectType string) Object {
	now := NewTimestamp()
	theType := objectType
	id, _ := IdentifierFromString(objectType + "--5d0092c5-5f74-4287-9642-33f4c354e56d")

	return Object{ID: id,
		Type:     objectType,
		Created:  now,
		Modified: now,
		Source:   []byte(`{"type": ` + theType + `}`),
	}
}
