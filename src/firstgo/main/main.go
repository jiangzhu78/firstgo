
// [_Command-line flags_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// are a common way to specify options for command-line
// programs. For example, in `wc -l` the `-l` is a
// command-line flag.

package main

// Go provides a `flag` package supporting basic
// command-line flag parsing. We'll use this package to
// implement our example command-line program.
import (
	"flag"
	"firstgo/httpserver"
	)

func main() {

	var port int

    flag.IntVar(&port, "port", 8080, "default port is 8080")
    flag.Parse()

    httpserver.Start(port)

}