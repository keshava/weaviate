//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

package crossrefs

import (
	"fmt"
	"testing"

	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/stretchr/testify/assert"
)

func TestNetworkCrossRefNames(t *testing.T) {
	testData := []struct {
		input          string
		expectedResult bool
	}{
		{"Car", false},
		{"WeviateB/Car", true},
		{"weviateb/Car", true},
		{"weviateb/car", false},
		{"/", false},
		{"/Car", false},
		{"weaviateB/", false},
		{"we4viateB/Car", false},
		{"weaviateB/C4r", false},
	}

	for _, data := range testData {
		var msg string
		if data.expectedResult {
			msg = fmt.Sprintf("'%s' should be a valid network cross-ref", data.input)
		} else {
			msg = fmt.Sprintf("'%s' should NOT be a valid network cross-ref", data.input)
		}
		t.Run(msg, func(t *testing.T) {
			if result := schema.ValidNetworkClassName(data.input); result != data.expectedResult {
				t.Errorf("wanted %v, but got %v", data.expectedResult, result)
			}
		})
	}
}

func TestParseClass(t *testing.T) {

	t.Run("a valid class name", func(t *testing.T) {
		ref, err := ParseClass("WeaviateB/Car")
		assert.Equal(t, nil, err, "should not error")
		assert.Equal(t, NetworkClass{PeerName: "WeaviateB", ClassName: "Car"}, ref,
			"should be parsed correctly into peerName and ClassName")
	})

	t.Run("an invalid class name", func(t *testing.T) {
		_, err := ParseClass("Car")
		assert.NotEqual(t, nil, err, "should error")
	})
}
