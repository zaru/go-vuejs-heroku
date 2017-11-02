// +build heroku

package goose2

import (
	// Invoke init() functions within migrations pkg.
	_ "github.com/zaru/go-vuejs-heroku/cmd/goose"
)
