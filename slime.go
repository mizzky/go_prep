// +build ignore

package cli_prep

import (
	"flag"
	"fmt"

	"github.com/mizzky/cli_prep/cmd"
)

// define option
var (
	show_version  = flag.Bool("version", false, "show version")
	show_artClose = flag.Bool("close", false, "show asciiart number")
	show_fav      = flag.Bool("fav", false, "show asciiart number")
)

const (
	version = "v0.10"
)

func main() {
	flag.Parse()

	if *show_version {
		fmt.Printf("version: %s\n", version)
		return
	}

	if *show_artClose {
		cmd.ArtClose()
	}

	if *show_fav {
		cmd.ArtFav()
	}

}
