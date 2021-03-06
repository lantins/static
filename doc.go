/*
Package static is a handler for the Core (https://github.com/volatile/core).
It serves files from the "static" directory on the "/static" URL path.

Installation

In the terminal:

	$ go get github.com/volatile/static

Usage

Example:

	package main

	import (
		"github.com/volatile/core"
		"github.com/volatile/static"
	)

	func main() {
		static.Use(static.DefaultMaxAge)

		core.Run()
	}

Cache max-age

The Use method receives a single parameter: the "max-age" (in seconds) for all resources.
This setting is only used in a production environment.
*/
package static
