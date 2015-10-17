package form

import (
	"fmt"
	"log"
	"reflect"
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

func GenInputWithLabel(field, name, inputType string) string {
	// add label
	form := fmt.Sprintf(`<label for="%s">%s</label>`, field, name)
	// new line
	form += "\n"
	// add input
	form += fmt.Sprintf(`<input type="%s" name="%s"`, inputType, field)
	// at the end add closing '>' and line break symbol '<br>'
	form += "><br>\n"
	return form
}
func GenInputWithLabel(field, name, inputType string, value interface{}) string {
	// add label
	form := fmt.Sprintf(`<label for="%s">%s</label>`, field, name)
	// new line
	form += "\n"
	// add input
	form += fmt.Sprintf(`<input type="%s" name="%s"`, inputType, field)
	// get value of default value and add if it is not empty
	val := fmt.Sprintf("%v", value)
	if val != "" {
		form += fmt.Sprintf(` value=%v`, val)
	}
	// at the end add closing '>' and line break symbol '<br>'
	form += "><br>\n"
	return form
}
func GenInput(field, name, inputType string, value interface{}) string {
	// add input
	form += fmt.Sprintf(`<input type="%s" name="%s"`, inputType, field)
	// get value of default value and add if it is not empty
	val := fmt.Sprintf("%v", value)
	if val != "" {
		form += fmt.Sprintf(` value=%v`, val)
	}
	// at the end add closing '>' and line break symbol '<br>'
	form += "><br>\n"
	return form
}
func FormCreate(form *MyForm) (string, error) {
	formType := reflect.TypeOf(*form)
	formValue := reflect.ValueOf(*form)
	var XMLForm string
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
			XMLForm += GenInputWithLabel(field.Tag.Get("field"), field.Tag.Get("name"), field.Tag.Get("type"), value)
		case "textarea":
			XMLForm += GenInputWithLabel(field.Tag.Get("field"), field.Tag.Get("name"), field.Tag.Get("type"), value)
		case "radio":
			XMLForm += GenInputWithLabel(field.Tag.Get("field"), field.Tag.Get("name"), field.Tag.Get("type"), value)
		case "password":
			XMLForm += GenInputWithLabel(field.Tag.Get("field"), field.Tag.Get("name"), field.Tag.Get("type"), value)
		case "hidden":
			XMLForm += GenInput(field.Tag.Get("field"), field.Tag.Get("name"), field.Tag.Get("type"), value)
		case "checkbox":
			XMLForm += fmt.Sprintf("<label for=\"%s\">%s</label>\n", field.Tag.Get("field"), field.Tag.Get("name"))
			XMLForm += fmt.Sprintf("<input type=\"%s\" name=\"%s\"><br>\n", field.Tag.Get("type"), field.Tag.Get("name"))
		case "button":
			XMLForm += fmt.Sprintf("<label for=\"%s\">%s</label>\n", field.Tag.Get("field"), field.Tag.Get("name"))
			XMLForm += fmt.Sprintf("<input type=\"%s\" name=\"%s\"><br>\n", field.Tag.Get("type"), field.Tag.Get("name"))
		case "select":
			XMLForm += fmt.Sprintf("<label for=\"%s\">%s</label>\n", field.Tag.Get("field"), field.Tag.Get("name"))
			XMLForm += fmt.Sprintf("<input type=\"%s\" name=\"%s\"><br>\n", field.Tag.Get("type"), field.Tag.Get("name"))
		default:
			log.Println("nothin")
		}
	}
	log.Println(XMLForm)
	return "", nil
}
