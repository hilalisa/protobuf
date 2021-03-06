// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: multi/multi3.proto

package multi

import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion2

type Multi3_HatType int

const (
	Multi3_FEDORA Multi3_HatType = 0
	Multi3_FEZ    Multi3_HatType = 1
)

var Multi3_HatType_name = map[int]string{
	0: "FEDORA",
	1: "FEZ",
}
var Multi3_HatType_value = map[string]int{
	"FEDORA": 0,
	"FEZ":    1,
}

func (x Multi3_HatType) String() string {
	return Multi3_HatType_name[int(x)]
}

type Multi3 struct {
	HatType Multi3_HatType
}

// GetHatType gets the HatType of the Multi3.
func (m *Multi3) GetHatType() (x Multi3_HatType) {
	if m == nil {
		return x
	}
	return m.HatType
}

// MarshalToWriter marshals Multi3 to the provided writer.
func (m *Multi3) MarshalToWriter(writer jspb.Writer) {
	if m == nil {
		return
	}

	if int(m.HatType) != 0 {
		writer.WriteEnum(1, int(m.HatType))
	}

	return
}

// Marshal marshals Multi3 to a slice of bytes.
func (m *Multi3) Marshal() []byte {
	writer := jspb.NewWriter()
	m.MarshalToWriter(writer)
	return writer.GetResult()
}

// UnmarshalFromReader unmarshals a Multi3 from the provided reader.
func (m *Multi3) UnmarshalFromReader(reader jspb.Reader) *Multi3 {
	for reader.Next() {
		if m == nil {
			m = &Multi3{}
		}

		switch reader.GetFieldNumber() {
		case 1:
			m.HatType = Multi3_HatType(reader.ReadEnum())
		default:
			reader.SkipField()
		}
	}

	return m
}

// Unmarshal unmarshals a Multi3 from a slice of bytes.
func (m *Multi3) Unmarshal(rawBytes []byte) (*Multi3, error) {
	reader := jspb.NewReader(rawBytes)

	m = m.UnmarshalFromReader(reader)

	if err := reader.Err(); err != nil {
		return nil, err
	}

	return m, nil
}
