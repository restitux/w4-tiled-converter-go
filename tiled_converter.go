package main

import (
	"flag"
	"fmt"
	"github.com/restitux/w4-tiled-converter/tilemap"
	"github.com/restitux/w4-tiled-converter/tileset"
	"io/ioutil"
	"encoding/json"
	"os"
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
	ext := path.Ext(filename)
	name := strings.TrimSuffix(filename, ext)
	dir, name := path.Split(name)
	// Create output filenames
	h_filename := path.Join(dir, name+".h")
	c_filename := path.Join(dir, name+".c")

	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tileset map[string]interface{}
	json.Unmarshal([]byte(byteValue), &tileset)

    tile_w = int(tileset_json['tilewidth'].(float64))
    tile_h = int(tileset_json['tileheight'].(float64))
    if tile_w != tile_h:
        panic("ERROR: Tiles of different h / w are not supported")

	png_filename := path.Join(dir, tileset["image"].(string))

	tileset.Convert(png_filename, h_filename, c_filename, tile_w, name)
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
