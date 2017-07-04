// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: my_test/test.proto

/*
	Package my_test is a generated protocol buffer package.

	This package holds interesting messages.

	It is generated from these files:
		my_test/test.proto

	It has these top-level messages:
		Request
		Reply
		OtherBase
		OtherReplyExtensions
		Communique
*/
package my_test

import js "github.com/gopherjs/gopherjs/js"
import jspb "github.com/johanbrandhorst/protobuf/jspb"
import multitest2 "github.com/johanbrandhorst/protobuf/protoc-gen-gopherjs/testdata/multi"

type HatType int

const (
	HatType_FEDORA HatType = 0
	HatType_FEZ    HatType = 1
)

var HatType_name = map[int]string{
	0: "FEDORA",
	1: "FEZ",
}
var HatType_value = map[string]int{
	"FEDORA": 0,
	"FEZ":    1,
}

func (x HatType) String() string {
	return HatType_name[int(x)]
}

// This enum represents days of the week.
type Days int

const (
	Days_MONDAY  Days = 0
	Days_TUESDAY Days = 2
	Days_LUNDI   Days = 0
)

var Days_name = map[int]string{
	0: "MONDAY",
	2: "TUESDAY",
	// Duplicate value: 0: "LUNDI",
}
var Days_value = map[string]int{
	"MONDAY":  0,
	"TUESDAY": 2,
	"LUNDI":   0,
}

func (x Days) String() string {
	return Days_name[int(x)]
}

type Request_Color int

const (
	Request_RED   Request_Color = 0
	Request_GREEN Request_Color = 1
	Request_BLUE  Request_Color = 2
)

var Request_Color_name = map[int]string{
	0: "RED",
	1: "GREEN",
	2: "BLUE",
}
var Request_Color_value = map[string]int{
	"RED":   0,
	"GREEN": 1,
	"BLUE":  2,
}

func (x Request_Color) String() string {
	return Request_Color_name[int(x)]
}

type Reply_Entry_Game int

const (
	Reply_Entry_FOOTBALL Reply_Entry_Game = 0
	Reply_Entry_TENNIS   Reply_Entry_Game = 2
)

var Reply_Entry_Game_name = map[int]string{
	0: "FOOTBALL",
	2: "TENNIS",
}
var Reply_Entry_Game_value = map[string]int{
	"FOOTBALL": 0,
	"TENNIS":   2,
}

func (x Reply_Entry_Game) String() string {
	return Reply_Entry_Game_name[int(x)]
}

// This is a message that might be sent somewhere.
type Request struct {
	*js.Object
}

// New creates a new Request.
// This is a map field. It will generate map[int32]string.
// This is a map field whose value type is a message.
// This field should not conflict with any getters.
// Imported type
// Property with name of Go keyword
func (m *Request) New(key []int64, hue Request_Color, hat HatType, deadline float32, nameMapping map[int32]string, msgMapping map[int64]*Reply, reset int32, getKey string, multi *multitest2.Multi1, _fallthrough string) *Request {
	m = &Request{
		Object: js.Global.Get("proto").Get("my").Get("test").Get("Request").New([]interface{}{
			js.Undefined,
			hue,
			hat,
			deadline,
			js.Undefined,
			js.Undefined,
			reset,
			getKey,
			multi.Call("toArray"),
			_fallthrough,
		}),
	}

	arr := js.Global.Get("Array").New(len(key))
	for i, value := range key {
		arr.SetIndex(i, value)
		m.Call("setKeyList", arr)
	}
	for key, value := range nameMapping {
		m.Call("getNameMappingMap").Call("set", key, value)
	}
	for key, value := range msgMapping {
		m.Call("getMsgMappingMap").Call("set", key, value)
	}

	return m
}

// GetKey gets the Key of the Request.
// Warning: mutating the returned slice will not be reflected in the message.
// Use the setter to make changes to the slice in the message.
func (m *Request) GetKey() []int64 {
	x := []int64{}
	arrFunc := func(value *js.Object) {
		x = append(x, value.Int64())
	}
	m.Call("getKeyList").Call("forEach", arrFunc)
	return x
}

// SetKey sets the Key of the Request.
func (m *Request) SetKey(v []int64) {
	arr := js.Global.Get("Array").New(len(v))
	for i, value := range v {
		arr.SetIndex(i, value)
	}
	m.Call("setKeyList", arr)
}

// GetHue gets the Hue of the Request.
func (m *Request) GetHue() Request_Color {
	return Request_Color(m.Call("getHue").Int())
}

// SetHue sets the Hue of the Request.
func (m *Request) SetHue(v Request_Color) {
	m.Call("setHue", v)
}

// GetHat gets the Hat of the Request.
func (m *Request) GetHat() HatType {
	return HatType(m.Call("getHat").Int())
}

// SetHat sets the Hat of the Request.
func (m *Request) SetHat(v HatType) {
	m.Call("setHat", v)
}

// GetDeadline gets the Deadline of the Request.
func (m *Request) GetDeadline() float32 {
	return float32(m.Call("getDeadline").Float())
}

// SetDeadline sets the Deadline of the Request.
func (m *Request) SetDeadline(v float32) {
	m.Call("setDeadline", v)
}

// GetNameMapping gets the NameMapping of the Request.
// This is a map field. It will generate map[int32]string.
func (m *Request) GetNameMapping() map[int32]string {
	x := map[int32]string{}
	mapFunc := func(value *js.Object, key *js.Object) {
		x[int32(key.Int())] = value.String()
	}
	m.Call("getNameMappingMap").Call("forEach", mapFunc)
	return x
}

// SetNameMapping sets the NameMapping of the Request.
// This is a map field. It will generate map[int32]string.
func (m *Request) SetNameMapping(v map[int32]string) {
	m.Call("clearNameMappingMap")
	for key, value := range v {
		m.Call("getNameMappingMap").Call("set", key, value)
	}
}

// GetMsgMapping gets the MsgMapping of the Request.
// This is a map field whose value type is a message.
func (m *Request) GetMsgMapping() map[int64]*Reply {
	x := map[int64]*Reply{}
	mapFunc := func(value *js.Object, key *js.Object) {
		x[key.Int64()] = &Reply{Object: value}
	}
	m.Call("getMsgMappingMap").Call("forEach", mapFunc)
	return x
}

// SetMsgMapping sets the MsgMapping of the Request.
// This is a map field whose value type is a message.
func (m *Request) SetMsgMapping(v map[int64]*Reply) {
	m.Call("clearMsgMappingMap")
	for key, value := range v {
		m.Call("getMsgMappingMap").Call("set", key, value)
	}
}

// GetReset gets the Reset of the Request.
func (m *Request) GetReset() int32 {
	return int32(m.Call("getReset").Int())
}

// SetReset sets the Reset of the Request.
func (m *Request) SetReset(v int32) {
	m.Call("setReset", v)
}

// GetGetKey gets the GetKey of the Request.
// This field should not conflict with any getters.
func (m *Request) GetGetKey() string {
	return m.Call("getGetKey").String()
}

// SetGetKey sets the GetKey of the Request.
// This field should not conflict with any getters.
func (m *Request) SetGetKey(v string) {
	m.Call("setGetKey", v)
}

// GetMulti gets the Multi of the Request.
// Imported type
func (m *Request) GetMulti() *multitest2.Multi1 {
	return &multitest2.Multi1{Object: m.Call("getMulti")}
}

// SetMulti sets the Multi of the Request.
// Imported type
func (m *Request) SetMulti(v *multitest2.Multi1) {
	m.Call("setMulti", v.Call("toArray"))
}

// GetFallthrough gets the Fallthrough of the Request.
// Property with name of Go keyword
func (m *Request) GetFallthrough() string {
	return m.Call("getFallthrough").String()
}

// SetFallthrough sets the Fallthrough of the Request.
// Property with name of Go keyword
func (m *Request) SetFallthrough(v string) {
	m.Call("setFallthrough", v)
}

// Serialize marshals Request to a slice of bytes.
func (m *Request) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Request from a slice of bytes.
func (m *Request) Deserialize(rawBytes []byte) (*Request, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("my").Get("test").Get("Request"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Request{
		Object: obj,
	}, nil
}

type Reply struct {
	*js.Object
}

// New creates a new Reply.
func (m *Reply) New(found []*Reply_Entry, compactKeys []int32) *Reply {
	m = &Reply{
		Object: js.Global.Get("proto").Get("my").Get("test").Get("Reply").New([]interface{}{
			js.Undefined,
			js.Undefined,
		}),
	}

	arr := js.Global.Get("Array").New(len(found))
	for i, value := range found {
		arr.SetIndex(i, value)
		m.Call("setFoundList", arr)
	}
	arr_ := js.Global.Get("Array").New(len(compactKeys))
	for i, value := range compactKeys {
		arr_.SetIndex(i, value)
		m.Call("setCompactKeysList", arr_)
	}

	return m
}

// GetFound gets the Found of the Reply.
// Warning: mutating the returned slice will not be reflected in the message.
// Use the setter to make changes to the slice in the message.
func (m *Reply) GetFound() []*Reply_Entry {
	x := []*Reply_Entry{}
	arrFunc := func(value *js.Object) {
		x = append(x, &Reply_Entry{Object: value})
	}
	m.Call("getFoundList").Call("forEach", arrFunc)
	return x
}

// SetFound sets the Found of the Reply.
func (m *Reply) SetFound(v []*Reply_Entry) {
	arr := js.Global.Get("Array").New(len(v))
	for i, value := range v {
		arr.SetIndex(i, value)
	}
	m.Call("setFoundList", arr)
}

// GetCompactKeys gets the CompactKeys of the Reply.
// Warning: mutating the returned slice will not be reflected in the message.
// Use the setter to make changes to the slice in the message.
func (m *Reply) GetCompactKeys() []int32 {
	x := []int32{}
	arrFunc := func(value *js.Object) {
		x = append(x, int32(value.Int()))
	}
	m.Call("getCompactKeysList").Call("forEach", arrFunc)
	return x
}

// SetCompactKeys sets the CompactKeys of the Reply.
func (m *Reply) SetCompactKeys(v []int32) {
	arr := js.Global.Get("Array").New(len(v))
	for i, value := range v {
		arr.SetIndex(i, value)
	}
	m.Call("setCompactKeysList", arr)
}

// Serialize marshals Reply to a slice of bytes.
func (m *Reply) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Reply from a slice of bytes.
func (m *Reply) Deserialize(rawBytes []byte) (*Reply, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("my").Get("test").Get("Reply"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Reply{
		Object: obj,
	}, nil
}

type Reply_Entry struct {
	*js.Object
}

// New creates a new Reply_Entry.
func (m *Reply_Entry) New(keyThatNeeds1234camelCasIng int64, value int64, MyFieldName2 int64) *Reply_Entry {
	m = &Reply_Entry{
		Object: js.Global.Get("proto").Get("my").Get("test").Get("Reply").Get("Entry").New([]interface{}{
			keyThatNeeds1234camelCasIng,
			value,
			MyFieldName2,
		}),
	}

	return m
}

// GetKeyThatNeeds_1234Camel_CasIng gets the KeyThatNeeds_1234Camel_CasIng of the Reply_Entry.
func (m *Reply_Entry) GetKeyThatNeeds_1234Camel_CasIng() int64 {
	return m.Call("getKeyThatNeeds_1234Camel_CasIng").Int64()
}

// SetKeyThatNeeds_1234Camel_CasIng sets the KeyThatNeeds_1234Camel_CasIng of the Reply_Entry.
func (m *Reply_Entry) SetKeyThatNeeds_1234Camel_CasIng(v int64) {
	m.Call("setKeyThatNeeds_1234Camel_CasIng", v)
}

// GetValue gets the Value of the Reply_Entry.
func (m *Reply_Entry) GetValue() int64 {
	return m.Call("getValue").Int64()
}

// SetValue sets the Value of the Reply_Entry.
func (m *Reply_Entry) SetValue(v int64) {
	m.Call("setValue", v)
}

// GetXMyFieldName_2 gets the XMyFieldName_2 of the Reply_Entry.
func (m *Reply_Entry) GetXMyFieldName_2() int64 {
	return m.Call("getXMyFieldName_2").Int64()
}

// SetXMyFieldName_2 sets the XMyFieldName_2 of the Reply_Entry.
func (m *Reply_Entry) SetXMyFieldName_2(v int64) {
	m.Call("setXMyFieldName_2", v)
}

// Serialize marshals Reply_Entry to a slice of bytes.
func (m *Reply_Entry) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Reply_Entry from a slice of bytes.
func (m *Reply_Entry) Deserialize(rawBytes []byte) (*Reply_Entry, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("my").Get("test").Get("Reply").Get("Entry"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Reply_Entry{
		Object: obj,
	}, nil
}

type OtherBase struct {
	*js.Object
}

// New creates a new OtherBase.
func (m *OtherBase) New(name string) *OtherBase {
	m = &OtherBase{
		Object: js.Global.Get("proto").Get("my").Get("test").Get("OtherBase").New([]interface{}{
			name,
		}),
	}

	return m
}

// GetName gets the Name of the OtherBase.
func (m *OtherBase) GetName() string {
	return m.Call("getName").String()
}

// SetName sets the Name of the OtherBase.
func (m *OtherBase) SetName(v string) {
	m.Call("setName", v)
}

// Serialize marshals OtherBase to a slice of bytes.
func (m *OtherBase) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a OtherBase from a slice of bytes.
func (m *OtherBase) Deserialize(rawBytes []byte) (*OtherBase, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("my").Get("test").Get("OtherBase"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &OtherBase{
		Object: obj,
	}, nil
}

type OtherReplyExtensions struct {
	*js.Object
}

// New creates a new OtherReplyExtensions.
func (m *OtherReplyExtensions) New(key int32) *OtherReplyExtensions {
	m = &OtherReplyExtensions{
		Object: js.Global.Get("proto").Get("my").Get("test").Get("OtherReplyExtensions").New([]interface{}{
			key,
		}),
	}

	return m
}

// GetKey gets the Key of the OtherReplyExtensions.
func (m *OtherReplyExtensions) GetKey() int32 {
	return int32(m.Call("getKey").Int())
}

// SetKey sets the Key of the OtherReplyExtensions.
func (m *OtherReplyExtensions) SetKey(v int32) {
	m.Call("setKey", v)
}

// Serialize marshals OtherReplyExtensions to a slice of bytes.
func (m *OtherReplyExtensions) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a OtherReplyExtensions from a slice of bytes.
func (m *OtherReplyExtensions) Deserialize(rawBytes []byte) (*OtherReplyExtensions, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("my").Get("test").Get("OtherReplyExtensions"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &OtherReplyExtensions{
		Object: obj,
	}, nil
}

type Communique struct {
	*js.Object
}

// New creates a new Communique.
func (m *Communique) New(makeMeCry bool) *Communique {
	m = &Communique{
		Object: js.Global.Get("proto").Get("my").Get("test").Get("Communique").New([]interface{}{
			makeMeCry,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
			js.Undefined,
		}),
	}

	return m
}

// This is a oneof, called "union".
//
// Types that are valid to be assigned to Union:
//	*Communique_Number
//	*Communique_Name
//	*Communique_Data
//	*Communique_TempC
//	*Communique_Height
//	*Communique_Today
//	*Communique_Maybe
//	*Communique_Delta_
//	*Communique_Msg
type isCommunique_Union interface {
	isCommunique_Union()
}

type Communique_Number struct {
	Number int32
}
type Communique_Name struct {
	Name string
}
type Communique_Data struct {
	Data []byte
}
type Communique_TempC struct {
	TempC float64
}
type Communique_Height struct {
	Height float32
}
type Communique_Today struct {
	Today Days
}
type Communique_Maybe struct {
	Maybe bool
}
type Communique_Delta_ struct {
	Delta int32
}
type Communique_Msg struct {
	Msg *Reply
}

func (*Communique_Number) isCommunique_Union() {}
func (*Communique_Name) isCommunique_Union()   {}
func (*Communique_Data) isCommunique_Union()   {}
func (*Communique_TempC) isCommunique_Union()  {}
func (*Communique_Height) isCommunique_Union() {}
func (*Communique_Today) isCommunique_Union()  {}
func (*Communique_Maybe) isCommunique_Union()  {}
func (*Communique_Delta_) isCommunique_Union() {}
func (*Communique_Msg) isCommunique_Union()    {}

func (m *Communique) GetUnion() (x isCommunique_Union) {
	switch m.Call("getUnionCase").Int() {
	case 5:
		x = &Communique_Number{
			Number: m.GetNumber(),
		}
	case 6:
		x = &Communique_Name{
			Name: m.GetName(),
		}
	case 7:
		x = &Communique_Data{
			Data: m.GetData(),
		}
	case 8:
		x = &Communique_TempC{
			TempC: m.GetTempC(),
		}
	case 9:
		x = &Communique_Height{
			Height: m.GetHeight(),
		}
	case 10:
		x = &Communique_Today{
			Today: m.GetToday(),
		}
	case 11:
		x = &Communique_Maybe{
			Maybe: m.GetMaybe(),
		}
	case 12:
		x = &Communique_Delta_{
			Delta: m.GetDelta(),
		}
	case 13:
		x = &Communique_Msg{
			Msg: m.GetMsg(),
		}
	}
	return x
}

// GetMakeMeCry gets the MakeMeCry of the Communique.
func (m *Communique) GetMakeMeCry() bool {
	return m.Call("getMakeMeCry").Bool()
}

// SetMakeMeCry sets the MakeMeCry of the Communique.
func (m *Communique) SetMakeMeCry(v bool) {
	m.Call("setMakeMeCry", v)
}

// GetNumber gets the Number of the Communique.
func (m *Communique) GetNumber() int32 {
	return int32(m.Call("getNumber").Int())
}

// SetNumber sets the Number of the Communique.
func (m *Communique) SetNumber(v int32) {
	m.Call("setNumber", v)
}

// GetName gets the Name of the Communique.
func (m *Communique) GetName() string {
	return m.Call("getName").String()
}

// SetName sets the Name of the Communique.
func (m *Communique) SetName(v string) {
	m.Call("setName", v)
}

// GetData gets the Data of the Communique.
func (m *Communique) GetData() []byte {
	return m.Call("getData").Interface().([]byte)
}

// SetData sets the Data of the Communique.
func (m *Communique) SetData(v []byte) {
	m.Call("setData", v)
}

// GetTempC gets the TempC of the Communique.
func (m *Communique) GetTempC() float64 {
	return m.Call("getTempC").Float()
}

// SetTempC sets the TempC of the Communique.
func (m *Communique) SetTempC(v float64) {
	m.Call("setTempC", v)
}

// GetHeight gets the Height of the Communique.
func (m *Communique) GetHeight() float32 {
	return float32(m.Call("getHeight").Float())
}

// SetHeight sets the Height of the Communique.
func (m *Communique) SetHeight(v float32) {
	m.Call("setHeight", v)
}

// GetToday gets the Today of the Communique.
func (m *Communique) GetToday() Days {
	return Days(m.Call("getToday").Int())
}

// SetToday sets the Today of the Communique.
func (m *Communique) SetToday(v Days) {
	m.Call("setToday", v)
}

// GetMaybe gets the Maybe of the Communique.
func (m *Communique) GetMaybe() bool {
	return m.Call("getMaybe").Bool()
}

// SetMaybe sets the Maybe of the Communique.
func (m *Communique) SetMaybe(v bool) {
	m.Call("setMaybe", v)
}

// GetDelta gets the Delta of the Communique.
func (m *Communique) GetDelta() int32 {
	return int32(m.Call("getDelta").Int())
}

// SetDelta sets the Delta of the Communique.
func (m *Communique) SetDelta(v int32) {
	m.Call("setDelta", v)
}

// GetMsg gets the Msg of the Communique.
func (m *Communique) GetMsg() *Reply {
	return &Reply{Object: m.Call("getMsg")}
}

// SetMsg sets the Msg of the Communique.
func (m *Communique) SetMsg(v *Reply) {
	m.Call("setMsg", v.Call("toArray"))
}

// Serialize marshals Communique to a slice of bytes.
func (m *Communique) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Communique from a slice of bytes.
func (m *Communique) Deserialize(rawBytes []byte) (*Communique, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("my").Get("test").Get("Communique"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Communique{
		Object: obj,
	}, nil
}

type Communique_Delta struct {
	*js.Object
}

// New creates a new Communique_Delta.
func (m *Communique_Delta) New() *Communique_Delta {
	m = &Communique_Delta{
		Object: js.Global.Get("proto").Get("my").Get("test").Get("Communique").Get("Delta").New([]interface{}{}),
	}

	return m
}

// Serialize marshals Communique_Delta to a slice of bytes.
func (m *Communique_Delta) Serialize() ([]byte, error) {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Communique_Delta from a slice of bytes.
func (m *Communique_Delta) Deserialize(rawBytes []byte) (*Communique_Delta, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("my").Get("test").Get("Communique").Get("Delta"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Communique_Delta{
		Object: obj,
	}, nil
}
