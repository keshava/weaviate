/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package adapters




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

/*WeaveAdaptersActivateOK Successful response

swagger:response weaveAdaptersActivateOK
*/
type WeaveAdaptersActivateOK struct {

	// In: body
	Payload *models.AdaptersActivateResponse `json:"body,omitempty"`
}

// NewWeaveAdaptersActivateOK creates WeaveAdaptersActivateOK with default headers values
func NewWeaveAdaptersActivateOK() *WeaveAdaptersActivateOK {
	return &WeaveAdaptersActivateOK{}
}

// WithPayload adds the payload to the weave adapters activate o k response
func (o *WeaveAdaptersActivateOK) WithPayload(payload *models.AdaptersActivateResponse) *WeaveAdaptersActivateOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weave adapters activate o k response
func (o *WeaveAdaptersActivateOK) SetPayload(payload *models.AdaptersActivateResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaveAdaptersActivateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}