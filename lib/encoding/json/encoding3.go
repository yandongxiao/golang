package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Server struct {
	Name string /* Why the field Name must be accessed */
	Ip   string
	Age  int
	D1   *int /* nil */
	D2   *int
}

type ServerSlice struct {
	Servers []Server
}

func (slice *ServerSlice) String() string {
	var output string
	for i := range slice.Servers {
		output += "D2 : " + strconv.Itoa(*slice.Servers[i].D2) + "\n"
	}
	return output
}

func main() {
	val := 2
	/* D2 point to the same value */
	slice := ServerSlice{
		[]Server{{"shanghai", "127.0.0.1", 10, nil, &val},
			{"beijing", "127.0.0.2", 20, nil, &val}}}

	data, _ := json.Marshal(slice)
	fmt.Println(string(data))

	/* D2 point to the different value */
	/* this is OK, because go do not allowed to compare pointers */
	inst := new(ServerSlice)
	json.Unmarshal(data, &inst)
	*inst.Servers[0].D2 = 100
	fmt.Println(inst)
}
