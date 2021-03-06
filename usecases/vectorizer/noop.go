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

package vectorizer

import (
	"context"

	"github.com/semi-technologies/weaviate/entities/models"
)

// NoOpVectorizer is a simple stand in that does nothing. Can be used when the
// feature should be turned off overall
type NoOpVectorizer struct{}

// Corpi is not implemented in the NoOpVectorizer
func (n *NoOpVectorizer) Corpi(ctx context.Context, corpi []string) ([]float32, error) {
	return []float32{}, nil
}

// MoveTo is not implemented in the NoOpVectorizer
func (n *NoOpVectorizer) MoveTo(source []float32, target []float32, weight float32) ([]float32, error) {
	return []float32{}, nil
}

// MoveAwayFrom is not implemented in the NoOpVectorizer
func (n *NoOpVectorizer) MoveAwayFrom(source []float32, target []float32, weight float32) ([]float32, error) {
	return []float32{}, nil
}

// NormalizedDistance is not implemented in the NoOpVectorizer
func (n *NoOpVectorizer) NormalizedDistance(a, b []float32) (float32, error) {
	return 0, nil
}

// Thing is not implemented in the NoOpVectorizer
func (n *NoOpVectorizer) Thing(ctx context.Context, concept *models.Thing) ([]float32, error) {
	return []float32{}, nil
}

// Action is not implemented in the NoOpVectorizer
func (n *NoOpVectorizer) Action(ctx context.Context, concept *models.Action) ([]float32, error) {
	return []float32{}, nil
}

// NewNoOp creates a new NoOpVectorizer which can be used when no vectorization
// is desired, i.e. the feature is turned off completely
func NewNoOp() *NoOpVectorizer {
	return &NoOpVectorizer{}
}
