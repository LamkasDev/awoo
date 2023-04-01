package instruction

const AwooInstructionFormatR = uint8(0x00)
const AwooInstructionFormatI = uint8(0x01)
const AwooInstructionFormatS = uint8(0x02)
const AwooInstructionFormatU = uint8(0x03)
const AwooInstructionFormatB = uint8(0x04)
const AwooInstructionFormatJ = uint8(0x05)

type AwooInstructionFormat struct {
	Code        AwooInstructionRange
	Destination AwooInstructionRange
	SourceOne   AwooInstructionRange
	SourceTwo   AwooInstructionRange
	Immediate   AwooInstructionRangeExtended
	Argument    AwooInstructionRangeExtended
}

var AwooInstructionFormats = map[uint8]AwooInstructionFormat{
	AwooInstructionFormatR: {
		Code:        AwooInstructionRange{Start: 0, Length: AwooInstructionCodeLength},
		Destination: AwooInstructionRange{Start: 7, Length: 5},
		SourceOne:   AwooInstructionRange{Start: 15, Length: 5},
		SourceTwo:   AwooInstructionRange{Start: 20, Length: 5},
		Argument: AwooInstructionRangeExtended{
			Ranges: []AwooInstructionRange{{Start: 12, Length: 3}, {Start: 25, Length: 7}},
		},
	},
	AwooInstructionFormatI: {
		Code:        AwooInstructionRange{Start: 0, Length: AwooInstructionCodeLength},
		Destination: AwooInstructionRange{Start: 7, Length: 5},
		SourceOne:   AwooInstructionRange{Start: 15, Length: 5},
		Immediate: AwooInstructionRangeExtended{
			Ranges: []AwooInstructionRange{{Start: 20, Length: 12}},
		},
		Argument: AwooInstructionRangeExtended{
			Ranges: []AwooInstructionRange{{Start: 12, Length: 3}},
		},
	},
	AwooInstructionFormatS: {
		Code:      AwooInstructionRange{Start: 0, Length: AwooInstructionCodeLength},
		SourceOne: AwooInstructionRange{Start: 15, Length: 5},
		SourceTwo: AwooInstructionRange{Start: 20, Length: 5},
		Immediate: AwooInstructionRangeExtended{
			Ranges: []AwooInstructionRange{{Start: 7, Length: 5}, {Start: 25, Length: 7}},
		},
		Argument: AwooInstructionRangeExtended{
			Ranges: []AwooInstructionRange{{Start: 12, Length: 3}},
		},
	},
	AwooInstructionFormatU: {
		Code:        AwooInstructionRange{Start: 0, Length: AwooInstructionCodeLength},
		Destination: AwooInstructionRange{Start: 7, Length: 5},
		Immediate: AwooInstructionRangeExtended{
			Offset: 12,
			Ranges: []AwooInstructionRange{{Start: 12, Length: 20}},
		},
	},
	AwooInstructionFormatB: {
		Code:      AwooInstructionRange{Start: 0, Length: AwooInstructionCodeLength},
		SourceOne: AwooInstructionRange{Start: 15, Length: 5},
		SourceTwo: AwooInstructionRange{Start: 20, Length: 5},
		Immediate: AwooInstructionRangeExtended{
			Offset: 1,
			Ranges: []AwooInstructionRange{{Start: 8, Length: 4}, {Start: 25, Length: 6}, {Start: 7, Length: 1}, {Start: 31, Length: 1}},
		},
		Argument: AwooInstructionRangeExtended{
			Ranges: []AwooInstructionRange{{Start: 12, Length: 3}},
		},
	},
	AwooInstructionFormatJ: {
		Code:        AwooInstructionRange{Start: 0, Length: AwooInstructionCodeLength},
		Destination: AwooInstructionRange{Start: 7, Length: 5},
		Immediate: AwooInstructionRangeExtended{
			Offset: 1,
			Ranges: []AwooInstructionRange{{Start: 21, Length: 10}, {Start: 20, Length: 1}, {Start: 12, Length: 8}, {Start: 31, Length: 1}},
		},
	},
}
