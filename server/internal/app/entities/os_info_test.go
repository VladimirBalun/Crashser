package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOSInfo(t *testing.T) {
	osInfo := NewOSInfo()
	require.Equal(t, 0, len(osInfo.GetName()))
	require.Equal(t, 0, len(osInfo.GetVersion()))
	require.Equal(t, 0, len(osInfo.GetArchitecture()))

	name := "windows"
	version := "10"
	architecture := "x64"

	osInfo.SetName(name)
	osInfo.SetVersion(version)
	osInfo.SetArchitecture(architecture)

	assert.Equal(t, name, osInfo.GetName())
	assert.Equal(t, version, osInfo.GetVersion())
	assert.Equal(t, architecture, osInfo.GetArchitecture())
}
