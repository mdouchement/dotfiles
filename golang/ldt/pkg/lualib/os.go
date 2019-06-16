package lualib

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Shopify/go-lua"
)

var osLibrary = []lua.RegistryFunction{
	{
		// os.osname()
		Name: "osname",
		Function: func(l *lua.State) int {
			l.PushString(runtime.GOOS)
			return 1
		},
	},
	{
		// os.chmod("~/tmp/binary", 0755)
		Name: "chmod",
		Function: func(l *lua.State) int {
			name := lua.CheckString(l, 1)
			mode, err := strconv.ParseUint(strconv.Itoa(lua.CheckInteger(l, 2)), 8, 32) // convert base10 to base8
			if err != nil {
				lua.Errorf(l, err.Error())
			}

			if err = os.Chmod(name, os.FileMode(mode)); err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.mkdir("~/tmp/something")
		Name: "mkdir",
		Function: func(l *lua.State) int {
			folder := lua.CheckString(l, 1)
			if err := os.Mkdir(folder, 0755); err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.mkdir_p("~/tmp/something")
		Name: "mkdir_p",
		Function: func(l *lua.State) int {
			folder := lua.CheckString(l, 1)
			if err := os.MkdirAll(folder, 0755); err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.cp("go.mod", "/tmp")
		Name: "cp",
		Function: func(l *lua.State) int {
			src := lua.CheckString(l, 1)
			dst := lua.CheckString(l, 2)
			if err := copy(src, dst); err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.cp_rf(".", "/tmp/project")
		Name: "cp_rf",
		Function: func(l *lua.State) int {
			src := lua.CheckString(l, 1)
			dst := lua.CheckString(l, 2)
			if err := copyRF(src, dst); err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.cp("go.mod")
		Name: "rm",
		Function: func(l *lua.State) int {
			filename := lua.CheckString(l, 1)
			if err := os.Remove(filename); err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.cp("/tmp/project")
		Name: "rm_rf",
		Function: func(l *lua.State) int {
			dst := lua.CheckString(l, 1)
			if err := os.RemoveAll(dst); err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
}

// OSOpen opens the os library. Usually passed to Require (local os = require "lualib/os").
func OSOpen(l *lua.State) {
	open := func(l *lua.State) int {
		lua.NewLibrary(l, osLibrary)
		return 1
	}
	lua.Require(l, "lualib/os", open, false)
	l.Pop(1)
}

// -------------
// --------------------------- //
// Utils                       //
// --------------------------- //
// --------

func copyRF(src, dst string) (err error) {
	if !isDir(src) {
		return copy(src, dst)
	}
	srcs, err := filepath.Glob(filepath.Join(src, "*"))
	if err != nil {
		return err
	}

	for _, s := range srcs {
		// Get destination path and create directory
		d := strings.Replace(s, src, "", 1)
		if src == "." {
			d = s // Rollback the replacement
		}
		d = filepath.Join(dst, d)
		mkdirAllWithFilename(d)
		//
		err = copyRF(s, d)
		if err != nil {
			return err
		}
	}
	return
}

func copy(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	return out.Sync()
}

func isDir(p string) bool {
	s, err := os.Stat(p)
	if err != nil {
		panic(err)
	}
	return s.IsDir()
}

func mkdirAllWithFilename(p string) {
	mkdirAll(filepath.Dir(p))
}

func mkdirAll(p string) {
	_ = os.MkdirAll(p, 0755)
}
