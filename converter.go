package main

import (
	"flag"
	"fmt"
	"github.com/restitux/w4-tiled-converter/tilemap"
	"path"
	"strings"
)

func tilemap_subcommand(filename string) {
	fmt.Printf("INFO: Generating tilename from file %v\n", filename)

	ext := path.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	dir, name := path.Split(name)
	// Create output filenames
	h_filename := path.Join(dir, name+".h")
	c_filename := path.Join(dir, name+".c")

	tilemap.Convert(filename, h_filename, c_filename, name)
}

func tileset_subcommand(filename string) {
}

func main() {

	flag.Parse()
	mode := flag.Args()[0]
	filename := flag.Args()[1]

	if flag.NArg() < 2 {
		panic("Error: Incorrect number of arguments")
	}

	if mode == "tilemap" {
		tilemap_subcommand(filename)
	} else if mode == "tileset" {
		tileset_subcommand(filename)
	} else {
		panic("Unkown command")
	}

}
