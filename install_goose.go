// +build heroku

package main

import (
	// Invoke init() functions within migrations pkg.
	_ "go-vuejs-heroku/cmd/goose"
)
