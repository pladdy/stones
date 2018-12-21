package stones

import (
	"fmt"
	"time"
)

func validTestObject(objectType string) Object {
	now, err := NewTimestamp(time.Now().Format(time.RFC3339Nano))
	if err != nil {
		fmt.Println("Failed to create now as a Timestamp", err)
	}

	theType := objectType
	id, _ := IdentifierFromString(objectType + "--5d0092c5-5f74-4287-9642-33f4c354e56d")

	return Object{ID: id,
		Type:     objectType,
		Created:  now,
		Modified: now,
		Source:   []byte(`{"type": ` + theType + `}`),
	}
}
