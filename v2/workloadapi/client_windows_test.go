//go:build windows
// +build windows

package workloadapi

import (
	"context"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vishnusomank/go-spiffe/v2/internal/test"
	"github.com/vishnusomank/go-spiffe/v2/internal/test/fakeworkloadapi"
)

func TestWithNamedPipeName(t *testing.T) {
	ca := test.NewCA(t, td)
	wl := fakeworkloadapi.NewWithNamedPipeListener(t)
	defer wl.Stop()

	pipeName := strings.TrimPrefix(wl.Addr(), "npipe:")
	c, err := New(context.Background(), WithNamedPipeName(pipeName))
	require.NoError(t, err)
	defer c.Close()
	require.Equal(t, pipeName, c.config.namedPipeName)

	resp := &fakeworkloadapi.X509SVIDResponse{
		Bundle: ca.X509Bundle(),
		SVIDs:  makeX509SVIDs(ca, fooID, barID),
	}
	wl.SetX509SVIDResponse(resp)
	svid, err := c.FetchX509SVID(context.Background())
	require.NoError(t, err)
	assertX509SVID(t, svid, fooID, resp.SVIDs[0].Certificates)
}

func TestWithNamedPipeNameError(t *testing.T) {
	wl := fakeworkloadapi.NewWithNamedPipeListener(t)
	defer wl.Stop()

	c, err := New(context.Background(), WithNamedPipeName("ohno"))
	require.NoError(t, err)
	defer c.Close()

	wl.SetX509SVIDResponse(&fakeworkloadapi.X509SVIDResponse{})
	_, err = c.FetchX509SVID(context.Background())
	require.Error(t, err)
	require.Contains(t, err.Error(), `ohno: The system cannot find the file specified`)
}
