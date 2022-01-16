package tilemap

type Tilemap struct {
	name   string
	width  int
	height int
	data   []int
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
