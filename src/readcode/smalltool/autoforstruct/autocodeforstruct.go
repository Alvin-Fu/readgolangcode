package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"unicode"
)

type Person struct {
}
type RobotsSignUp struct {
	startTime  int64
	svrType    int64
	svrId      int64
	onceCount  int64
	robotCount int64
	interval   int64
	matchId    string
	gameId     string
}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
func SetReceiverName(name string) string {

	return ""
}

func main() {
	var tmp = new(RobotsSignUp)

	nameType, fieldType, name := GetFieldName(tmp)
	fmt.Println(strings.Title(name))
	var buffer bytes.Buffer
	buffer.WriteString("package main\n")
	for i, s := range nameType {
		buffer.WriteString(fmt.Sprintf("func (t *%s) %s() %s{\n", name, Ucfirst(s), fieldType[i]))
		buffer.WriteString(fmt.Sprintf("return t.%s\n", s))
		buffer.WriteString(fmt.Sprintf("}\n"))
		buffer.WriteString(fmt.Sprintf("func (t *%s) Set%s(%s %s){\n", name, Ucfirst(s), s, fieldType[i]))
		buffer.WriteString(fmt.Sprintf("t.%s = %s\n", s, Lcfirst(s)))
		buffer.WriteString(fmt.Sprintf("}\n"))
	}
	f, _ := os.Create(name + ".go")
	f.Write([]byte(buffer.String()))
	f.Close()
	cmd := exec.Command("goimports", "-w", name)
	cmd.Run()
	fmt.Println("ok！")
}
func GetFieldName(structName interface{}) ([]string, []string, string) {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("Check type error not Struct")
		return nil, nil, ""
	}
	fieldNum := t.NumField()
	fieldName := make([]string, 0, fieldNum)
	fieldType := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		if strings.Contains(t.Field(i).Type.String(), ".") {
			tmp := strings.Split(t.Field(i).Type.String(), ".")
			if strings.Contains(t.Field(i).Type.String(), "[") {
				fieldType = append(fieldType, "[]"+tmp[1])
			} else {
				fieldType = append(fieldType, tmp[1])
			}

		} else {
			fieldType = append(fieldType, t.Field(i).Type.String())
		}
		fieldName = append(fieldName, t.Field(i).Name)
	}
	return fieldName, fieldType, t.Name()
}
