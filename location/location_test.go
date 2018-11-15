package location

import (
	"testing"

	mesg "github.com/mesg-foundation/go-service"
	mesgtest "github.com/mesg-foundation/go-service/mesgtest"
	"github.com/stvp/assert"
)

const (
	maxmindDBPath = "../data/GeoLite2-City/GeoLite2-City.mmdb"

	token    = "token"
	endpoint = "endpoint"
)

func newServiceAndServer(t *testing.T) (*mesg.Service, *mesgtest.Server) {
	testServer := mesgtest.NewServer()
	service, err := mesg.New(
		mesg.DialOption(testServer.Socket()),
		mesg.TokenOption(token),
		mesg.EndpointOption(endpoint),
	)
	assert.Nil(t, err)
	assert.NotNil(t, service)
	return service, testServer
}
