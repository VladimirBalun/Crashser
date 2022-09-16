package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCoreDump(t *testing.T) {
	coreDump := NewCoreDump()
	require.Nil(t, coreDump.GetOSInfo())
	require.Nil(t, coreDump.GetAppInfo())
	require.Equal(t, ToDoCoreDumpStatus, coreDump.GetStatus())
	require.Equal(t, time.Time{}, coreDump.GetTimestamp())
	require.Equal(t, 0, len(coreDump.GetData()))

	osInfo := NewOSInfo()
	appInfo := NewAppInfo()
	timestamp := time.Now()
	status := SolvedCoreDumpStatus
	data := "server/internal/app/entities/core_dump_test.go@TestCoreDump:100"

	coreDump.SetOSInfo(osInfo)
	coreDump.SetAppInfo(appInfo)
	coreDump.SetTimestamp(timestamp)
	coreDump.SetStatus(status)
	coreDump.SetData(data)

	assert.Equal(t, osInfo, coreDump.GetOSInfo())
	assert.Equal(t, appInfo, coreDump.GetAppInfo())
	assert.Equal(t, timestamp, coreDump.GetTimestamp())
	assert.Equal(t, status, coreDump.GetStatus())
	assert.Equal(t, data, coreDump.GetData())
}
