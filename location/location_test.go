package location

import (
	"testing"

	"github.com/mesg-foundation/core/client/service"
	"github.com/mesg-foundation/core/client/service/servicetest"
	"github.com/stretchr/testify/require"
)

const (
	maxmindDBPath = "../data/GeoLite2-City/GeoLite2-City.mmdb"

	token    = "token"
	endpoint = "endpoint"
)

func newServiceAndServer(t *testing.T) (*service.Service, *servicetest.Server) {
	testServer := servicetest.NewServer()
	s, err := service.New(
		service.DialOption(testServer.Socket()),
		service.TokenOption(token),
		service.EndpointOption(endpoint),
	)
	require.NoError(t, err)
	require.NotNil(t, s)
	return s, testServer
}
