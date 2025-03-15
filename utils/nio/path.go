package nio

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root             = filepath.Join(filepath.Dir(b), "../..")
	configFolderName = "config"
	resFolderName    = "resources"
	buildFolderName  = "build"
)

type Path struct {
	uri string
}

func (p Path) String() string {
	return p.uri
}

func RootPath() Path {
	return Path{Root}
}

func ResPath() Path {
	return RootPath().Resolve(resFolderName)
}

func ConfigPath() Path {
	return RootPath().Resolve(configFolderName)
}

func BuildPath() Path {
	return RootPath().Resolve(buildFolderName)
}

func (p Path) Resolve(elem ...string) Path {
	return Path{filepath.Join(p.uri, filepath.Join(elem...))}
}

func (p Path) Parent() Path {
	return Path{filepath.Dir(p.String())}
}

func (p Path) Name() string {
	return filepath.Base(p.String())
}
