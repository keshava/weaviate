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

package schema

import (
	"testing"
)

func TestValidateOKClassName(t *testing.T) {
	_, err := ValidateClassName("FooBar")
	if err != nil {
		t.Fail()
	}
}

func TestFailValidateBadClassName(t *testing.T) {
	_, err := ValidateClassName("Foo Bar")
	if err == nil {
		t.Fail()
	}

	_, err = ValidateClassName("foo")
	if err == nil {
		t.Fail()
	}

	_, err = ValidateClassName("fooBar")
	if err == nil {
		t.Fail()
	}
}

func TestValidateOKPropertyName(t *testing.T) {
	_, err := ValidatePropertyName("fooBar")
	if err != nil {
		t.Fail()
	}
}

func TestFailValidateBadPropertyName(t *testing.T) {
	_, err := ValidatePropertyName("foo Bar")
	if err == nil {
		t.Fail()
	}

	_, err = ValidatePropertyName("Foo")
	if err == nil {
		t.Fail()
	}

	_, err = ValidatePropertyName("FooBar")
	if err == nil {
		t.Fail()
	}
}
