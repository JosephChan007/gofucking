package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"
)

type Student struct {
	Id    int64
	Name  string
	Age   int
	High  float64
	Birth int64
}

func main() {

	stu := Student{
		Id:    1001,
		Name:  "xiaoqiang",
		Age:   34,
		High:  1.75,
		Birth: time.Now().Unix(),
	}

	b := new(bytes.Buffer)
	encoder := gob.NewEncoder(b)
	err := encoder.Encode(stu)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	decoder := gob.NewDecoder(bytes.NewBuffer(b.Bytes()))
	err = decoder.Decode(stu)
	if err != nil {
		fmt.Errorf("%v", err)
	}

	fmt.Printf("%#v", stu)
}
