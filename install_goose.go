// +build heroku

package main

import (
	// Invoke init() functions within migrations pkg.
	_ "github.com/zaru/go-vuejs-heroku/cmd"
)
