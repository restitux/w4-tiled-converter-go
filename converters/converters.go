package main

import (
	"fmt"
)

import (
	"encoding/json"
	"fmt"
	"github.com/restitux/w4-tiled-converter/sources"
	"io/ioutil"
	"os"
)

func ConvertTilemap(filename string, h_filename string, c_filename string, name string) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("INFO: Successfully opened %v\n", filename)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tilemap map[string]interface{}
	json.Unmarshal([]byte(byteValue), &tilemap)

	row := tilemap["layers"].([]interface{})[0].(map[string]interface{})
	data_h := int(row["height"].(float64))
	data_w := int(row["width"].(float64))
	data_f := row["data"].([]interface{})
	data := make([]int, 0)
	for _, v := range data_f {
		data = append(data, int(v.(float64)))
	}

	s := sources.CreateSource(h_filename, c_filename)
	s.AddTilemap(name, data_w, data_h, data)
	s.ToFile()
}


import (
	"fmt"
	"image/png"
	"os"
)

type Color struct {
	r int
	g int
	b int
}

func getPixelColorId(c Color) int {
	red := Color{255, 0, 0}
	black := Color{0, 0, 0}
	gray := Color{168, 168, 168}
	white := Color{255, 255, 255}
	if c == red {
		return 0
	} else if c == black {
		return 1
	} else if c == gray {
		return 2
	} else if c == white {
		return 3
	} else {
		panic("ERROR: unknown color")
	}
}

//def convert_region(tile_id, region):
//    result = []
//    for y in range(region.size[1]): #framebuffer coords = y * 160 + x
//        for x in range(region.size[0]):
//            color_id = get_pixel_color_id(region.getpixel((x, y)))
//            result.append(color_id)
//    return result
//
//
//def convert(png_filename : str, h_filename : str, c_filename : str, tilesize : int, name : str):
//    png = Image.open(png_filename)
//    print(f"image is {png.format} of {png.size}")
//
//    tile_id = 0
//    color_ids = []
//    for tile_x in range(0, png.size[0], tilesize):
//        for tile_y in range(0, png.size[1], tilesize):
//            tile_region = (tile_x, tile_y, tile_x + tilesize, tile_y + tilesize)
//            tile_colors = convert_region(tile_id, png.crop(tile_region))
//            color_ids.extend(tile_colors)
//            tile_id += 1
//
//
//    s = sources.Sources(h_filename, c_filename)
//    s.add_tileset(name, tilesize, png.size[0], png.size[1], color_ids)
//    s.to_file()
//}

//func convert_region(tile_id, region):
func ConvertRegion(tile_id string, region string) {

}

//result = []
//for y in range(region.size[1]): #framebuffer coords = y * 160 + x
//    for x in range(region.size[0]):
//        color_id = get_pixel_color_id(region.getpixel((x, y)))
//        result.append(color_id)
//return result

//def convert(png_filename : str, h_filename : str, c_filename : str, tilesize : int, name : str):
func ConvertTileset(png_filename string, h_filename string, c_filename string, tilesize int, name string) {
	pngFile, err := os.Open(png_filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("INFO: Successfully opened %v\n", png_filename)

	defer pngFile.Close()

	pngdata, err := png.Decode(pngFile)
	if err != nil {
		panic(err)
	}

	tile_id := 0
	colors_ids := make([]Color, 0)
	for {

	}

	//png = Image.open(png_filename)
	//print(f"image is {png.format} of {png.size}")

	//tile_id = 0
	//color_ids = []
	//for tile_x in range(0, png.size[0], tilesize):
	//    for tile_y in range(0, png.size[1], tilesize):
	//        tile_region = (tile_x, tile_y, tile_x + tilesize, tile_y + tilesize)
	//        tile_colors = convert_region(tile_id, png.crop(tile_region))
	//        color_ids.extend(tile_colors)
	//        tile_id += 1

	//s = sources.Sources(h_filename, c_filename)
	//s.add_tileset(name, tilesize, png.size[0], png.size[1], color_ids)
	//s.to_file()
}
