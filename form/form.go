package form

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type MyForm struct {
	UserName     string `required:"true" field:"name" name:"Имя пользователя" type:"text"`
	UserPassword string `required:"true" field:"password" name:"Пароль пользователя" type:"password"`
	Resident     bool   `field:"resident" type:"radio" radio:"1;checked" name:"Резидент РФ"`
	NoResident   bool   `field:"resident" type:"radio" radio:"2" name:"Не резидент РФ"`
	Gender       string `field:"gender" name:"Пол" type:"select" select:"Не известный=3;selected,Мужской=1,Женский=2"`
	Age          int64  `field:"age" name:"Возраст" type:"text" default:"true"`
	Token        string `field:"token" type:"hidden" default:"true"`
}

func GenInputWithLabel(tags reflect.StructTag, value interface{}) string {
	var (
		field = tags.Get("field")
		name  = tags.Get("name")
	)
	// add label
	form := fmt.Sprintf(`<label for="%s">%s</label>`, field, name)
	// new line
	form += "\n"
	// add input
	form += GenInput(tags, value)
	return form
}
func GenInput(tags reflect.StructTag, value interface{}) string {
	var field, inputType = tags.Get("field"), tags.Get("type")
	// add input
	form := fmt.Sprintf(`<input type="%s" name="%s"`, inputType, field)
	// get value of default value and add if it is not empty
	val := fmt.Sprintf("%v", value)
	if tags.Get("default") == "true" {
		form += fmt.Sprintf(` value="%v"`, val)
	}
	// if input type is radio, then parse its value
	if inputType == "radio" {
		s := strings.Split(tags.Get("radio"), ";")
		form += fmt.Sprintf(` value="%s" `, s[0])
		if len(s) == 2 {
			form += s[1]
		}
	}
	// at the end add closing '>' and line break symbol '<br>'
	form += "><br>\n"
	return form
}

func GenSelect(tags reflect.StructTag) string {
	var (
		field   = tags.Get("field")
		name    = tags.Get("name")
		options = tags.Get("select")
	)
	// add options
	form := fmt.Sprintf(`<label for="%s">%s</label>`, field, name)
	form += "\n"
	form += fmt.Sprintf(`<select name="%s">`, field)
	form += "\n"
	for _, option := range strings.Split(options, ",") {
		form += "\t"
		selected := ""
		s := strings.Split(option, ";")
		if len(s) > 1 {
			selected = s[1]
		}
		optionValues := strings.Split(s[0], "=")
		form += fmt.Sprintf(`<option value="%s" %s>%s</option>`, optionValues[1], selected, optionValues[0])
		form += "\n"
	}
	form += "</select>"
	form += "\n"
	return form
}
func FormCreate(form *MyForm) (string, error) {
	formType := reflect.TypeOf(*form)
	formValue := reflect.ValueOf(*form)
	var XMLForm string = "<form action='/create' method='post' enctype='multipart/form-data'>\n"
	for i := 0; i < formType.NumField(); i++ {
		field := formType.Field(i)
		value := formValue.Field(i)
		if field.Tag.Get("field") == "" || field.Tag.Get("field") == "-" {
			continue
		}
		switch field.Tag.Get("type") {
		case "":
			fallthrough
		case "text":
			fallthrough
		case "textarea":
			fallthrough
		case "checkbox":
			fallthrough
		case "password":
			fallthrough
		case "button":
			fallthrough
		case "radio":
			XMLForm += GenInputWithLabel(field.Tag, value)
		case "hidden":
			XMLForm += GenInput(field.Tag, value)
		case "select":
			XMLForm += GenSelect(field.Tag)
		default:
			log.Println("nothin")
		}
	}
	XMLForm += "<button type='submit'>send</button>"
	XMLForm += "</form>"
	return XMLForm, nil
}
