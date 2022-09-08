package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCoreDump(t *testing.T) {
	coreDump := NewCoreDump()
	require.Nil(t, coreDump.OSInfo())
	require.Nil(t, coreDump.AppInfo())
	require.Equal(t, ToDoCoreDumpStatus, coreDump.Status())
	require.Equal(t, time.Time{}, coreDump.Timestamp())
	require.Equal(t, 0, len(coreDump.Data()))

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

	assert.Equal(t, osInfo, coreDump.OSInfo())
	assert.Equal(t, appInfo, coreDump.AppInfo())
	assert.Equal(t, timestamp, coreDump.Timestamp())
	assert.Equal(t, status, coreDump.Status())
	assert.Equal(t, data, coreDump.Data())
}
