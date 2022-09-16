package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAppInfo(t *testing.T) {
	appInfo := NewAppInfo()
	require.Equal(t, 0, len(appInfo.GetName()))
	require.Equal(t, 0, len(appInfo.GetVersion()))
	require.Equal(t, UnknownProgrammingLanguage, appInfo.GetProgrammingLanguage())

	name := "publisher-app"
	version := "v0.0.1"
	language := JavaProgrammingLanguage

	appInfo.SetName(name)
	appInfo.SetVersion(version)
	appInfo.SetProgrammingLanguage(language)

	assert.Equal(t, name, appInfo.GetName())
	assert.Equal(t, version, appInfo.GetVersion())
	assert.Equal(t, language, appInfo.GetProgrammingLanguage())
}
