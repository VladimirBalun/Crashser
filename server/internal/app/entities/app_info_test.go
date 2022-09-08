package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppInfo(t *testing.T) {
	appInfo := NewAppInfo()
	require.Equal(t, 0, len(appInfo.Name()))
	require.Equal(t, 0, len(appInfo.Version()))
	require.Equal(t, UnknownProgrammingLanguage, appInfo.ProgrammingLanguage())

	name := "publisher-app"
	version := "v0.0.1"
	language := JavaProgrammingLanguage

	appInfo.SetName(name)
	appInfo.SetVersion(version)
	appInfo.SetProgrammingLanguage(language)

	assert.Equal(t, name, appInfo.Name())
	assert.Equal(t, version, appInfo.Version())
	assert.Equal(t, language, appInfo.ProgrammingLanguage())
}
