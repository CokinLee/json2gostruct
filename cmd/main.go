package main

import (
	"fmt"

	"github.com/CokinLee/json2gostruct"
)

// test j
var j = `
{
  "name": "测试",
  "code": "002082",
  "info": {
    "c": "10.56",
    "h": "10.62", 
    "mk": 2,
    "sp": 6.36,
    "isrzrq": false
  },
  "obj": {
	"mapString": "string",
	"mapInt": 1,
	"mapFloat": 1.2,
	"mapBool": true,
	"mapObj": {
		"objStr":"objStr"
	}
  },
  "data": [
    "2006-11-20,3.47,3.58,3.78,3.31,190365,195671708,-",
    "2006-11-21,3.44,3.99,3.99,3.39,109324,120198757,16.7%"
  ],
  "fsli": [
	1.23,
    1,
    3
  ]
}
`

func main() {
	s, _ := json2gostruct.CreateJSONModel(j, "J")
	fmt.Println(s)
}
