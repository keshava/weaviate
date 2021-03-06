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

package rest

import (
	"context"
	"encoding/json"
	"fmt"

	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/meta"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/p2_p"
	"github.com/semi-technologies/weaviate/adapters/handlers/rest/operations/well_known"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/usecases/config"
	"github.com/semi-technologies/weaviate/usecases/network"
	"github.com/semi-technologies/weaviate/usecases/network/common/peers"
	"github.com/semi-technologies/weaviate/usecases/telemetry"
)

type schemaManager interface {
	GetSchema(principal *models.Principal) (schema.Schema, error)
	GetSchemaSkipAuth() schema.Schema
}

type swaggerJSON struct {
	Info struct {
		Version string `json:"version"`
	} `json:"info"`
}

type c11yMetaProvider interface {
	Version(ctx context.Context) (string, error)
	WordCount(ctx context.Context) (int64, error)
}

func setupMiscHandlers(api *operations.WeaviateAPI, requestsLog *telemetry.RequestsLog,
	serverConfig *config.WeaviateConfig, network network.Network, schemaManager schemaManager, c11y c11yMetaProvider) {

	var swj swaggerJSON
	err := json.Unmarshal(SwaggerJSON, &swj)
	if err != nil {
		panic(err)
	}

	api.MetaMetaGetHandler = meta.MetaGetHandlerFunc(func(params meta.MetaGetParams, principal *models.Principal) middleware.Responder {
		// Create response object
		c11yVersion, err := c11y.Version(context.Background())
		if err != nil {
			return meta.NewMetaGetInternalServerError().WithPayload(errPayloadFromSingleErr(err))
		}

		c11yWordCount, err := c11y.WordCount(context.Background())
		if err != nil {
			return meta.NewMetaGetInternalServerError().WithPayload(errPayloadFromSingleErr(err))
		}

		res := &models.Meta{
			Hostname:               serverConfig.GetHostAddress(),
			Version:                swj.Info.Version,
			ContextionaryVersion:   c11yVersion,
			ContextionaryWordCount: c11yWordCount,
		}

		// Register the request
		go func() {
			requestsLog.Register(telemetry.TypeREST, telemetry.LocalQueryMeta)
		}()
		return meta.NewMetaGetOK().WithPayload(res)
	})

	api.P2pP2pGenesisUpdateHandler = p2_p.P2pGenesisUpdateHandlerFunc(func(params p2_p.P2pGenesisUpdateParams) middleware.Responder {
		newPeers := make([]peers.Peer, 0)

		for _, genesisPeer := range params.Peers {
			peer := peers.Peer{
				ID:         genesisPeer.ID,
				Name:       genesisPeer.Name,
				URI:        genesisPeer.URI,
				SchemaHash: genesisPeer.SchemaHash,
			}

			newPeers = append(newPeers, peer)
		}

		err := network.UpdatePeers(newPeers)

		if err == nil {
			// Register the request
			go func() {
				requestsLog.Register(telemetry.TypeREST, telemetry.NetworkQueryMeta)
			}()
			return p2_p.NewP2pGenesisUpdateOK()
		}
		return p2_p.NewP2pGenesisUpdateInternalServerError()
	})

	api.P2pP2pHealthHandler = p2_p.P2pHealthHandlerFunc(func(params p2_p.P2pHealthParams) middleware.Responder {
		// For now, always just return success.
		return middleware.NotImplemented("operation P2PP2pHealth has not yet been implemented")
	})

	api.WellKnownGetWellKnownOpenidConfigurationHandler = well_known.GetWellKnownOpenidConfigurationHandlerFunc(
		func(params well_known.GetWellKnownOpenidConfigurationParams, principal *models.Principal) middleware.Responder {
			if !serverConfig.Config.Authentication.OIDC.Enabled {
				return well_known.NewGetWellKnownOpenidConfigurationNotFound()
			}

			target := fmt.Sprintf("%s/.well-known/openid-configuration", serverConfig.Config.Authentication.OIDC.Issuer)
			clientID := serverConfig.Config.Authentication.OIDC.ClientID
			body := &well_known.GetWellKnownOpenidConfigurationOKBody{
				Href:     target,
				ClientID: clientID,
			}

			return well_known.NewGetWellKnownOpenidConfigurationOK().WithPayload(body)
		})

	api.WeaviateRootHandler = operations.WeaviateRootHandlerFunc(
		func(params operations.WeaviateRootParams, principal *models.Principal) middleware.Responder {
			origin := serverConfig.Config.Origin
			body := &operations.WeaviateRootOKBody{
				Links: []*models.Link{
					&models.Link{
						Name: "Meta information about this instance/cluster",
						Href: fmt.Sprintf("%s/v1/meta", origin),
					},
					&models.Link{
						Name:              "view complete schema",
						Href:              fmt.Sprintf("%s/v1/schema", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/add-data/define_schema.html",
					},
					&models.Link{
						Name:              "CRUD things schema",
						Href:              fmt.Sprintf("%s/v1/schema/things{/:className}", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/add-data/define_schema.html",
					},
					&models.Link{
						Name:              "CRUD actions schema",
						Href:              fmt.Sprintf("%s/v1/schema/actions{/:className}", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/add-data/define_schema.html",
					},
					&models.Link{
						Name:              "CRUD things",
						Href:              fmt.Sprintf("%s/v1/things{/:id}", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/add-data/add_and_modify.html",
					},
					&models.Link{
						Name:              "CRUD actions",
						Href:              fmt.Sprintf("%s/v1/actions{/:id}", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/add-data/add_and_modify.html",
					},
					&models.Link{
						Name:              "trigger and view status of classifications",
						Href:              fmt.Sprintf("%s/v1/classifications{/:id}", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/features/contextual-classification.html,https://www.semi.technology/documentation/weaviate/current/features/knn-classification.html",
					},
					&models.Link{
						Name:              "search contextionary for concepts",
						Href:              fmt.Sprintf("%s/v1/c11y/concepts/:concept", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/features/adding-synonyms.html",
					},
					&models.Link{
						Name:              "extend contextionary with custom extensions",
						Href:              fmt.Sprintf("%s/v1/c11y/extensions", origin),
						DocumentationHref: "https://www.semi.technology/documentation/weaviate/current/features/adding-synonyms.html",
					},
				},
			}

			return operations.NewWeaviateRootOK().WithPayload(body)
		})
}
