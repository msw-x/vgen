package vgen

import (
	"fmt"
	"msw/moon"
)

type Go struct {
	Package  string
	Name     string
	Filename string
}

func (o *Go) WithDefault() *Go {
	o.WithDefaultPackage()
	o.WithDefaultName()
	o.WithDefaultFilename()
	return o
}

func (o *Go) WithDefaultPackage() *Go {
	o.Package = "main"
	return o
}

func (o *Go) WithDefaultName() *Go {
	o.Name = "version"
	return o
}

func (o *Go) WithDefaultFilename() *Go {
	o.Filename = "version.go"
	return o
}

func (o *Go) Gen(version string) string {
	return fmt.Sprintf("package %s\n\nconst %s = \"%s\"\n", o.Package, o.Name, version)
}

func (o *Go) GenFile(version string, directory string) {
	moon.WriteFileStr(moon.PathJoin(directory, o.Filename), o.Gen(version))
}
