package test

import (
	"encoding/json"
	"fmt"
	"github.com/jaimera/poc-services/server/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

// Test Unmarshal on a real file
func TestParseJsonFile(t *testing.T) {
	bytes, err := os.ReadFile("../ports.json")
	if err != nil {
		fmt.Print(err)
	}
	var jsonMap map[string]map[string]interface{}
	err = json.Unmarshal(bytes, &jsonMap)
	assert.Nil(t, err, "Error on unmarshal")

	ports, objectLength, err := utils.ConvertPortJson(jsonMap)

	assert.Nil(t, err, "Error converting")
	assert.Equal(t, len(ports), objectLength)
}
