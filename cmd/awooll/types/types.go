package types

const AwooTypeFlagsSign = 1

type AwooType struct {
	Key    string
	Type   uint16
	Length uint8
	Size   uint16
	Flags  uint64
}

type AwooTypeMap struct {
	All      map[uint16]AwooType
	Lookup   map[string]*AwooType
	Position uint16
}

func AddType(m *AwooTypeMap, key string, size uint16, flags uint64) {
	awooType := AwooType{
		Key:    key,
		Type:   m.Position,
		Length: uint8(len(key)),
		Size:   size,
		Flags:  flags,
	}
	m.All[m.Position] = awooType
	m.Lookup[key] = &awooType
	m.Position++
}

func SetupTypeMap() AwooTypeMap {
	m := AwooTypeMap{
		All:    make(map[uint16]AwooType),
		Lookup: make(map[string]*AwooType),
	}

	AddType(&m, "bool", 1, 0)
	AddType(&m, "byte", 1, 0)
	AddType(&m, "char", 4, 0)

	AddType(&m, "int", 4, AwooTypeFlagsSign)
	AddType(&m, "int8", 1, AwooTypeFlagsSign)
	AddType(&m, "int16", 2, AwooTypeFlagsSign)
	AddType(&m, "int32", 4, AwooTypeFlagsSign)
	AddType(&m, "int64", 8, AwooTypeFlagsSign)

	AddType(&m, "uint", 4, 0)
	AddType(&m, "uint8", 1, 0)
	AddType(&m, "uint16", 2, 0)
	AddType(&m, "uint32", 4, 0)
	AddType(&m, "uint64", 8, 0)

	/* AddType(&m, "float", 4)
	AddType(&m, "float32", 4)
	AddType(&m, "float64", 8) */

	return m
}
