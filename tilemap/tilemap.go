package tilemap

import (
	"encoding/json"
	"fmt"
	"github.com/restitux/w4-tiled-converter/sources"
	"io/ioutil"
	"os"
)

func Convert(filename string, h_filename string, c_filename string, name string) {
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
