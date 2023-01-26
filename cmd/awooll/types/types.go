package types

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

func AddTypeAt(m *AwooTypeMap, t uint16, key string, size uint16, flags uint64) uint16 {
	awooType := AwooType{
		Key:    key,
		Type:   t,
		Length: uint8(len(key)),
		Size:   size,
		Flags:  flags,
	}
	m.All[t] = awooType
	m.Lookup[key] = &awooType

	return t
}

func AddType(m *AwooTypeMap, key string, size uint16, flags uint64) uint16 {
	t := AddTypeAt(m, m.Position, key, size, flags)
	m.Position++

	return t
}

func SetupTypeMap() AwooTypeMap {
	m := AwooTypeMap{
		All:    make(map[uint16]AwooType),
		Lookup: make(map[string]*AwooType),
	}

	AddTypeAt(&m, AwooTypeBoolean, "bool", 1, 0)
	AddTypeAt(&m, AwooTypeByte, "byte", 1, 0)
	AddTypeAt(&m, AwooTypeChar, "char", 4, 0)

	AddTypeAt(&m, AwooTypeInt8, "int8", 1, AwooTypeFlagsSign)
	AddTypeAt(&m, AwooTypeInt16, "int16", 2, AwooTypeFlagsSign)
	AddTypeAt(&m, AwooTypeInt32, "int32", 4, AwooTypeFlagsSign)
	AddTypeAt(&m, AwooTypeInt64, "int64", 8, AwooTypeFlagsSign)

	AddTypeAt(&m, AwooTypeUInt8, "uint8", 1, 0)
	AddTypeAt(&m, AwooTypeUIn16, "uint16", 2, 0)
	AddTypeAt(&m, AwooTypeUInt32, "uint32", 4, 0)
	AddTypeAt(&m, AwooTypeUInt64, "uint64", 8, 0)

	/* AddType(&m, "float", 4)
	AddType(&m, "float32", 4)
	AddType(&m, "float64", 8) */

	return m
}
