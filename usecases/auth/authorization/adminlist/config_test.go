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

package adminlist

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Validation(t *testing.T) {
	t.Run("with only an admin list set", func(t *testing.T) {
		cfg := Config{
			Enabled: true,
			Users: []string{
				"alice",
				"johndoe",
			},
		}

		err := cfg.Validate()
		assert.Nil(t, err)
	})

	t.Run("with only an read only list set", func(t *testing.T) {
		cfg := Config{
			Enabled: true,
			ReadOnlyUsers: []string{
				"alice",
				"johndoe",
			},
		}

		err := cfg.Validate()
		assert.Nil(t, err)
	})

	t.Run("with both lists present, but no overlap", func(t *testing.T) {
		cfg := Config{
			Enabled: true,
			Users: []string{
				"alice",
			},
			ReadOnlyUsers: []string{
				"johndoe",
			},
		}

		err := cfg.Validate()
		assert.Nil(t, err)
	})

	t.Run("with one subject part of both lists", func(t *testing.T) {
		cfg := Config{
			Enabled: true,
			Users: []string{
				"alice",
				"johndoe",
			},
			ReadOnlyUsers: []string{
				"johndoe",
			},
		}

		err := cfg.Validate()
		assert.Equal(t, err, fmt.Errorf("admin list: subject 'johndoe' is present on both admin and read-only list"))
	})
}
