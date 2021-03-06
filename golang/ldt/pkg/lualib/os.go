package lualib

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
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
		// os.user_exists("myuser")
		Name: "user_exists",
		Function: func(l *lua.State) int {
			username := lua.CheckString(l, 1)
			_, err := user.Lookup(username)
			l.PushBoolean(err == nil)
			return 1
		},
	},
	{
		// os.user_id("myuser")
		Name: "user_id",
		Function: func(l *lua.State) int {
			username := lua.CheckString(l, 1)
			u, err := user.Lookup(username)
			if err != nil {
				lua.Errorf(l, err.Error())
			}
			l.PushString(u.Uid)
			return 1
		},
	},
	{
		// os.group_id("mygroup")
		Name: "group_id",
		Function: func(l *lua.State) int {
			groupname := lua.CheckString(l, 1)
			g, err := user.LookupGroup(groupname)
			if err != nil {
				lua.Errorf(l, err.Error())
			}
			l.PushString(g.Gid)
			return 1
		},
	},
	{
		// os.touch("/tmp/ldt.db")
		Name: "touch",
		Function: func(l *lua.State) int {
			filename := lua.CheckString(l, 1)

			if exist(filename) {
				return 1
			}

			f, err := os.Create(filename)
			if err != nil {
				lua.Errorf(l, err.Error())
			}
			defer f.Close()
			return 1
		},
	},
	{
		// os.exist("~/tmp/binary")
		Name: "exist",
		Function: func(l *lua.State) int {
			path := lua.CheckString(l, 1)
			l.PushBoolean(exist(path))
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
		// os.chown("~/tmp/binary", 1001, 1001, true)  -> recursive
		// os.chown("~/tmp/binary", 1001, 1001, false) -> not recursive
		Name: "chown",
		Function: func(l *lua.State) int {
			root := lua.CheckString(l, 1)
			uid := lua.CheckInteger(l, 2)
			gid := lua.CheckInteger(l, 3)
			recursive := l.ToBoolean(4)

			if !recursive {
				if err := os.Chown(root, uid, gid); err != nil {
					lua.Errorf(l, err.Error())
				}
				return 1
			}

			err := filepath.Walk(root, func(path string, _ os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				return os.Chown(path, uid, gid)
			})

			if err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.symlink("/tmp/ldt.db", "link.db")
		Name: "symlink",
		Function: func(l *lua.State) int {
			src := lua.CheckString(l, 1)
			dst := lua.CheckString(l, 2)
			if err := os.Symlink(src, dst); err != nil {
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
		// os.mv("/src", "/dst")
		Name: "mv",
		Function: func(l *lua.State) int {
			src := lua.CheckString(l, 1)
			dst := lua.CheckString(l, 2)

			if exist(dst) {
				stat, err := os.Stat(dst)
				if err != nil {
					lua.Errorf(l, err.Error())
				}
				if stat.IsDir() {
					dst = filepath.Join(dst, filepath.Base(src))
				}
			}

			if err := os.Rename(src, dst); err != nil {
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
	{
		// os.exec("useradd", "--no-create-home", "--shell", "/sbin/nologin", "myuser")
		Name: "exec",
		Function: func(l *lua.State) int {
			name := lua.CheckString(l, 1)

			var args []string
			for i := 2; i <= l.Top(); i++ {
				s, ok := l.ToString(i)
				if !ok {
					lua.Errorf(l, "arg[%d] = %v is not a string", i, l.ToValue(i))
				}
				args = append(args, s)
			}

			cmd := exec.Command(name, args...)
			std, err := cmd.CombinedOutput()
			if len(std) > 0 {
				fmt.Println(string(std))
			}
			if err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.exec_in("/workdir", "useradd", "--no-create-home", "--shell", "/sbin/nologin", "myuser")
		Name: "exec_in",
		Function: func(l *lua.State) int {
			workdir := lua.CheckString(l, 1)
			name := lua.CheckString(l, 2)

			var args []string
			for i := 3; i <= l.Top(); i++ {
				s, ok := l.ToString(i)
				if !ok {
					lua.Errorf(l, "arg[%d] = %v is not a string", i, l.ToValue(i))
				}
				args = append(args, s)
			}

			cmd := exec.Command(name, args...)
			cmd.Dir = workdir
			std, err := cmd.CombinedOutput()
			if len(std) > 0 {
				fmt.Println(string(std))
			}
			if err != nil {
				lua.Errorf(l, err.Error())
			}
			return 1
		},
	},
	{
		// os.read_file("go.mod")
		Name: "read_file",
		Function: func(l *lua.State) int {
			data, err := os.ReadFile(lua.CheckString(l, 1))
			if err != nil {
				lua.Errorf(l, err.Error())
			}
			l.PushString(string(data))
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

func exist(p string) bool {
	_, err := os.Stat(p)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true // ignoring error
}
