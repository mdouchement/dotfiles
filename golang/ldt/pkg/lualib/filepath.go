package lualib

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/Shopify/go-lua"
)

var filepathLibrary = []lua.RegistryFunction{
	{
		// filepath.dirname("pkg/go.mod")
		Name: "dirname",
		Function: func(l *lua.State) int {
			path := lua.CheckString(l, 1)
			l.PushString(filepath.Dir(path))
			return 1
		},
	},
	{
		// filepath.basename("pkg/go.mod")
		Name: "basename",
		Function: func(l *lua.State) int {
			path := lua.CheckString(l, 1)
			l.PushString(filepath.Base(path))
			return 1
		},
	},
	{
		// filepath.join("path", "to", "file")
		Name: "join",
		Function: func(l *lua.State) int {
			var vargs []string
			for i := 1; i <= l.Top(); i++ {
				s, ok := l.ToString(i)
				if !ok {
					lua.Errorf(l, "arg[%d] = %v is not a string", i, l.ToValue(i))
				}
				vargs = append(vargs, s)
			}
			l.PushString(filepath.Join(vargs...))
			return 1
		},
	},
	{
		// filepath.expand("~/.go/bin/../")
		Name: "expand",
		Function: func(l *lua.State) int {
			path := lua.CheckString(l, 1)

			// Cleanup separator
			path = filepath.FromSlash(path)

			usr, err := user.Current()
			if err != nil {
				lua.Errorf(l, err.Error())
			}

			// Handle shell hom shortcut.
			if strings.HasPrefix(path, "~") {
				path = strings.Replace(path, "~", usr.HomeDir, 1)
			}

			// Replace environment variables by their values.
			path = os.ExpandEnv(path)

			// Compute absolute path.
			path, err = filepath.Abs(path)
			if err != nil {
				lua.Errorf(l, err.Error())
			}

			l.PushString(path)
			return 1
		},
	},
}

// FilePathOpen opens the filepath library. Usually passed to Require (local filepath = require "lualib/filepath").
func FilePathOpen(l *lua.State) {
	open := func(l *lua.State) int {
		lua.NewLibrary(l, filepathLibrary)
		return 1
	}
	lua.Require(l, "lualib/filepath", open, false)
	l.Pop(1)
}
