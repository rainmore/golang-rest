package nio

import (
	"github.com/rainmore/rest-api/tests"
	"github.com/rainmore/rest-api/utils/log"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	tests.SetupUnitTests()
	code := m.Run()
	tests.TearDownUnitTests()
	os.Exit(code)
}

func TestRootPath(t *testing.T) {
	rootPath := RootPath()
	log.Logger().Infof("rootPath: %q", rootPath)
	assert.True(t, strings.HasSuffix(rootPath.String(), "golang-rest"))
}

func TestResPath(t *testing.T) {
	resPath := ResPath()
	log.Logger().Infof("resPath: %q", resPath)
	assert.True(t, strings.HasSuffix(resPath.String(), "resources"))
}

func TestConfigPath(t *testing.T) {
	configPath := ConfigPath()
	log.Logger().Infof("configPath: %q", configPath)
	assert.True(t, strings.HasSuffix(configPath.String(), "config"))
}

func TestBuildPath(t *testing.T) {
	buildPath := BuildPath()
	log.Logger().Infof("buildPath: %q", buildPath)
	assert.True(t, strings.HasSuffix(buildPath.String(), "build"))
}

func TestPath_Parent(t *testing.T) {
	resPath := ResPath()
	parentPath := resPath.Parent()
	log.Logger().Infof("parentPath: %q", parentPath)
	assert.Equal(t, RootPath(), parentPath)
}

func TestPath_Name(t *testing.T) {
	resPath := ResPath()
	log.Logger().Infof("resPath: %q", resPath)
	assert.Equal(t, "resources", resPath.Name())
}
