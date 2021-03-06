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

package testhelper

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fakeT struct {
	lastError error
}

func (f *fakeT) Reset() {
	f.lastError = nil
}

func (f *fakeT) Errorf(msg string, args ...interface{}) {
	f.lastError = fmt.Errorf(msg, args...)
}

// AssertEventuallyEqual retries the 'actual' thunk every 10ms for a total of
// 300ms. If a single one succeeds, it returns, if all fails it eventually
// fails
func AssertEventuallyEqual(t *testing.T, expected interface{}, actualThunk func() interface{}, msg ...interface{}) {
	interval := 10 * time.Millisecond
	timeout := 4000 * time.Millisecond
	elapsed := 0 * time.Millisecond
	fakeT := &fakeT{}

	for elapsed < timeout {
		fakeT.Reset()
		actual := actualThunk()
		assert.Equal(fakeT, expected, actual, msg...)

		if fakeT.lastError == nil {
			return
		}

		time.Sleep(interval)
		elapsed += interval
	}

	t.Errorf("waiting for %s, but never succeeded:\n\n%s", elapsed, fakeT.lastError)
}

func AssertEventuallyEqualWithFrequencyAndTimeout(t *testing.T, expected interface{}, actualThunk func() interface{},
	interval time.Duration, timeout time.Duration, msg ...interface{}) {
	elapsed := 0 * time.Millisecond
	fakeT := &fakeT{}

	for elapsed < timeout {
		fakeT.Reset()
		actual := actualThunk()
		assert.Equal(fakeT, expected, actual, msg...)

		if fakeT.lastError == nil {
			return
		}

		time.Sleep(interval)
		elapsed += interval
	}

	t.Errorf("waiting for %s, but never succeeded:\n\n%s", elapsed, fakeT.lastError)
}
