package location

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocateSuccess(t *testing.T) {
	var (
		data = locateInputs{
			IP: "212.2.212.126",
		}
		service, server = newServiceAndServer(t)
		ls              = New(service, maxmindDBPath)
	)

	go server.Start()
	go ls.Start()

	_, execution, err := server.Execute("locate", data)
	require.NoError(t, err)
	require.Equal(t, "location", execution.Key())

	var outputs locateLocationOutputs
	require.NoError(t, execution.Data(&outputs))
	require.Equal(t, "Istanbul", outputs.City)
	require.Equal(t, "Turkey", outputs.Country)
	require.Equal(t, "TR", outputs.CountryISOCode)
	require.Equal(t, "Asia", outputs.Continent)
	require.Equal(t, "AS", outputs.ContinentCode)
	require.Equal(t, "Europe/Istanbul", outputs.TimeZone)
	require.Equal(t, 41.0483, outputs.Latitude)
	require.Equal(t, 28.9521, outputs.Longitude)
	require.Equal(t, "34445", outputs.PostalCode)
}

func TestLocateError(t *testing.T) {
	var (
		data = locateInputs{
			IP: "invalid",
		}
		service, server = newServiceAndServer(t)
		ls              = New(service, maxmindDBPath)
	)
	go server.Start()
	go ls.Start()

	_, execution, err := server.Execute("locate", data)
	require.NoError(t, err)
	require.Equal(t, "error", execution.Key())

	var outputs errorOutput
	require.NoError(t, execution.Data(&outputs))
	require.Contains(t, "IP address is not valid", outputs.Message)
}
