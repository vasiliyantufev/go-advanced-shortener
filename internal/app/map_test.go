package app

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dataTEST = NewDM()

func TestPutHandler(t *testing.T) {

	tests := []struct {
		name  string
		key   string
		value string
	}{
		{name: "Test1", key: "key1", value: "val1"},
		{name: "Test2", key: "key2", value: "val2"},
		{name: "Test3", key: "key3", value: "val3"},
	}

	for _, tt := range tests {
		dataTEST.Put(tt.key, tt.value)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			val, ok := dataTEST.Get(tt.key)

			assert.Equal(t, ok, true,
				fmt.Sprintf("Element %v not found", val))
		})
	}
}
