package json2gostruct

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// CreateJSONModel 根据json生成go struct
// jsonStr:需要生成的标准json格式字符串
// structName:struct名，需要导出则首字符大写
func CreateJSONModel(jsonStr string, structName string) (string, error) {
	m, err := json2map(jsonStr)
	if err != nil {
		return "", fmt.Errorf("json转换成map对象失败，%s", err.Error())
	}
	return createStruct(m, structName), nil
}

// json2map json转换成map对象
func json2map(jsonStr string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

//根据map对象字符串生成结构体
func createStruct(m map[string]interface{}, structName string) string {
	var buffer bytes.Buffer
	var innerM []string
	buffer.WriteString("type ")
	buffer.WriteString(structName)
	buffer.WriteString(" struct {\n")
	for k, v := range m {
		buffer.WriteString("\t")
		runes := []rune(k)
		fieldName := strings.ToUpper(string(runes[0])) + string(runes[1:])
		buffer.WriteString(fieldName)
		buffer.WriteString("   ")
		vts := reflect.TypeOf(v).String()
		if vts == "map[string]interface {}" {
			innerStructName := structName + fieldName
			buffer.WriteString(innerStructName)
			innerM = append(innerM, createStruct(v.(map[string]interface{}), innerStructName))
		} else if vts == "[]interface {}" {
			var vType string
			var same = true
			for i, sliV := range v.([]interface{}) {
				sliVT := reflect.TypeOf(sliV).String()
				if sliVT == "float64" && floatIsInt(sliV.(float64)) {
					sliVT = "int64"
				}
				if i == 0 {
					vType = sliVT
				} else if sliVT != vType {
					same = false
				}
			}
			if same {
				buffer.WriteString("[]")
				buffer.WriteString(vType)
			} else {
				buffer.WriteString(strings.TrimSpace(vts))
			}
		} else if vts == "float64" && floatIsInt(v.(float64)) {
			buffer.WriteString("int64")
		} else {
			buffer.WriteString(strings.TrimSpace(vts))
		}
		// 写tag
		buffer.WriteString("     `")
		buffer.WriteString("json:\""+k+"\" ")
		buffer.WriteString("bson:\""+k+"\" ")
		buffer.WriteString("`")
		buffer.WriteString("\n")
	}
	buffer.WriteString("}\n\n")
	for _, in := range innerM {
		buffer.WriteString(in)
	}

	return buffer.String()
}

// floatIsInt 判断float小数部是否为0
func floatIsInt(f float64) bool {
	s := fmt.Sprintf("%.10f", f)
	fs := strings.Split(s, ".")
	i, err := strconv.Atoi(fs[1])
	if err == nil && i == 0 {
		return true
	}
	return false
}
