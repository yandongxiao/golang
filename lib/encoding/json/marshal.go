// Marshal traverses the value v recursively.
// NOTE: The nil pointer exception is not strictly necessary but mimics a similar, necessary exception in the behavior of UnmarshalJSON.
package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	marshalJSONTest()
	marshalTextTest()
	primaryTypeTest()
	structTypeTest()
}

// If an encountered value implements the Marshaler interface
// and is not a nil pointer Marshal calls its MarshalJSON method to produce JSON
type Float64 float64

func (v Float64) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(v))), nil
}

func (v *Float64) UnmarshalJSON(data []byte) error {
	// 不需要处理data==null的情况，因为Float64的值不可能为null
	i, _ := strconv.Atoi(string(data))
	*v = Float64(i)
	return nil
}

func marshalJSONTest() {
	f := Float64(6.4)
	checkMarshal(json.Marshal(f))

	data, _ := json.Marshal(f)
	json.Unmarshal(data, &f)
	fmt.Println(f)

	data, _ = json.Marshal(nil)
	json.Unmarshal(data, &f)
	fmt.Println(f)
}

// If no MarshalJSON method is present but the value implements encoding.TextMarshaler instead,
// Marshal calls its MarshalText method and encodes the result as a JSON **string**.
type Float32 float32

func (v Float32) MarshalText() ([]byte, error) {
	//return []byte(strconv.Itoa(int(v))), nil
	return nil, nil
}

func marshalTextTest() {
	f := Float32(6.4)
	checkMarshal(json.Marshal(f))
}

func primaryTypeTest() {
	// nil
	//var data []byte = nil
	checkMarshal(json.Marshal(nil))
	checkMarshal(json.Marshal(3))
	checkMarshal(json.Marshal(3.5))
	checkMarshal(json.Marshal("name"))
	// NOTE: 注意与string的不同. []byte encodes as a base64-encoded string
	checkMarshal(json.Marshal([]byte("name")))
}

// Struct values encode as JSON objects. 即struct被JSON化以后，本身就是一个完整的json对象
// Each exported struct field becomes a member of the object, using the field name as the object key
func structTypeTest() {
	v := struct {
		// the format string: stored under the "json" key in the struct field's tag
		//	  The format string gives the name of the field, possibly followed by a comma-separated list of options.
		//     The name may be empty in order to specify options without overriding the default field name.
		Name   string      `json:"name"`
		Age    int         `json:",omitempty"` // omitempty: the field should be omitted from the encoding if the field has an empty value
		Desc   string      `json:"-"`          // As a special case, if the field tag is "-", the field is always omitted. 注意与json:"-,"的区别（按照普通语义理解）
		Sex    bool        `json:",string"`    // value以字符串形式输出. only to fields of string, floating point, integer, or boolean types.
		Child              // Anonymous struct fields are usually marshaled as if their inner exported fields were fields in the outer struct
		Child2 `json:"c2"` // An anonymous struct field with a name given in its JSON tag is treated as having that name, rather than being anonymous.
		Child3             // n anonymous struct field of interface type is treated the same as having that type as its name, rather than being anonymous.
		Child4 `json:"c4"` // 和Child都拥有Num，按照golang的语义v.Num是不能被访问的。在json中，因为Child4存在tag，所以它会被编码. 否则，Num不会被编码
	}{
		Name: "jack",
		Age:  10,
		Desc: "a ha",
	}
	//fmt.Println(v.Num) // ambiguous selector v.Num
	checkMarshal(json.Marshal(v))
}

type Child struct {
	Num int
}

type Child4 struct {
	Num int
}

type Child2 struct {
	Num2 int
}

type Child3 interface {
	String() string
}

func checkMarshal(v []byte, err error) {
	if err != nil {
		panic(err)
	}

	trace()
	fmt.Println(string(v))
}

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(3, pc)
	f := runtime.FuncForPC(pc[0])
	//file, line := f.FileLine(pc[0])
	fmt.Printf("%s: ", f.Name())
}
