package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOSInfo(t *testing.T) {
	osInfo := NewOSInfo()
	require.Equal(t, 0, len(osInfo.Name()))
	require.Equal(t, 0, len(osInfo.Version()))
	require.Equal(t, 0, len(osInfo.Architecture()))

	name := "windows"
	version := "10"
	architecture := "x64"

	osInfo.SetName(name)
	osInfo.SetVersion(version)
	osInfo.SetArchitecture(architecture)

	assert.Equal(t, name, osInfo.Name())
	assert.Equal(t, version, osInfo.Version())
	assert.Equal(t, architecture, osInfo.Architecture())
}
