package SystemNX

func getVariableType(variable interface{}) string {
	var ret string
	switch variable.(type) {
	case int64:
		ret = "int64"
	case int32:
		ret = "int32"
	case int16:
		ret = "int16"
	case int8:
		ret = "int8"
	case uint64:
		ret = "uint64"
	case uint32:
		ret = "uint32"
	case uint16:
		ret = "uint16"
	}
	return ret
}
