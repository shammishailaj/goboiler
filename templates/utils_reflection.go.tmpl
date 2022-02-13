package utils

import (
	"log"
	"reflect"
)

func GetFieldTagMap(d interface{}, tagIdentifier string) map[string]string {
	log.Printf("============================================================GetFieldTagMap()")
	log.Printf("d = %#v", d)
	var fieldTagMap map[string]string

	if d == nil {
		return fieldTagMap
	}

	v := reflect.TypeOf(d)
	reflectValue := reflect.ValueOf(d)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fieldTagMap = make(map[string]string)

	log.Printf("Type of d = %s. Kind: %s", v.Name(), v.Kind())
	if v != nil {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			//fieldTagMap[field.Name] = v.FieldByIndex(i).Tag
			fieldTagMap[field.Name] = field.Tag.Get(tagIdentifier)
			log.Printf("Found field Tag = %s for field named: %s", fieldTagMap[field.Name], field.Name)
		}
	}
	return fieldTagMap
}

func GetTagValueMap(d interface{}, tagIdentifier string) map[string]string {
	log.Printf("============================================================GetFieldTagMap()")
	log.Printf("d = %#v", d)
	var tagValueMap map[string]string

	if d == nil {
		return tagValueMap
	}

	v := reflect.TypeOf(d)
	val := reflect.ValueOf(d)
	reflectValue := reflect.ValueOf(d)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	tagValueMap = make(map[string]string)

	log.Printf("Type of d = %s. Kind: %s", v.Name(), v.Kind())
	if v != nil {
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			//fieldTagMap[field.Name] = v.FieldByIndex(i).Tag
			//tagValueMap[field.Name] = field.Tag.Get(tagIdentifier)
			tagName := field.Tag.Get(tagIdentifier)
			tagValueMap[tagName] = reflect.Indirect(val).FieldByName(field.Name).String()

			log.Printf("Found field Tag = %s, Value: %s", tagName, tagValueMap[tagName])
		}
	}
	return tagValueMap
}

func GetFieldValue(d interface{}, field string) interface{} {
	var fieldValue interface{}

	if d == nil {
		return fieldValue
	}

	v := reflect.TypeOf(d)
	reflectValue := reflect.ValueOf(d)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	log.Printf("Type of d = %s. Kind: %s", v.Name(), v.Kind())

	r := reflect.ValueOf(d)
	fieldValue = reflect.Indirect(r).FieldByName(field)
	return fieldValue
}