package sources

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Tilemap struct {
	name   string
	width  int
	height int
	data   []int
}

type array struct {
	h string
	c string
}

type Sources struct {
	h_filename string
	c_filename string
	includes   []string
	defines    map[string]string
	arrays     []array
}

func CreateTilemap(name string, width int, height int, data []int) Tilemap {
	return Tilemap{
		name,
		width,
		height,
		data,
	}
}

func (t Tilemap) to_c_src() array {
	data_str_arr := make([]string, 0)
	for _, num := range t.data {
		data_str_arr = append(data_str_arr, fmt.Sprint(num))
	}
	data_str := strings.Join(data_str_arr, ", ")

	h := fmt.Sprintf("extern const uint32_t %v_tilemap[%v];\n", t.name, len(data_str_arr))
	c := fmt.Sprintf("const uint32_t %v_tilemap[] = {%v};\n", t.name, data_str)

	return array{
		h,
		c,
	}
}

func CreateSource(h_filename string, c_filename string) Sources {
	return Sources{
		h_filename,
		c_filename,
		[]string{"#include <stdint.h>\n"},
		make(map[string]string),
		make([]array, 0),
	}
}

func (s *Sources) AddDefine(name string, val string) {
	s.defines[name] = val
}

func (s *Sources) AddArray(a array) {
	fmt.Printf("adding array %v\n", a)
	s.arrays = append(s.arrays, a)
}

func (s *Sources) AddTilemap(name string, width int, height int, data []int) {
	tilemap := CreateTilemap(name, width, height, data)
	a := tilemap.to_c_src()
	s.AddArray(a)
}

func (s Sources) PrintHeader() string {
	o := ""

	for _, v := range s.includes {
		o += v
	}

	for k, v := range s.defines {
		o += fmt.Sprintf("#define %v %v\n", k, v)
	}

	for _, v := range s.arrays {
		fmt.Println(v)
		o += v.h
	}

	return o
}

func (s Sources) PrintSource() string {
	o := ""

	o += fmt.Sprintf("#include \"%v\"\n", path.Base(s.h_filename))

	for _, v := range s.arrays {
		o += v.c
	}

	return o
}

func (s Sources) ToFile() {
	h_file, err := os.Create(s.h_filename)
	if err != nil {
		fmt.Println(err)
	} else {
		h_file.WriteString(s.PrintHeader())
	}
	h_file.Close()

	c_file, err := os.Create(s.c_filename)
	if err != nil {
		fmt.Println(err)
	} else {
		c_file.WriteString(s.PrintSource())
	}
	c_file.Close()
}
