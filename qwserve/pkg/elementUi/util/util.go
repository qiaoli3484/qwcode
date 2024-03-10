package util

import (
	"reflect"
	"strconv"
	"strings"
)

func AttrRef(el interface{}) string {

	valueof := reflect.Indirect(reflect.ValueOf(el))
	typeof := valueof.Type()
	num := typeof.NumField()
	arr := make([]string, 0)
	sty := make([]string, 0)

	for i := 0; i < num; i++ {
		sf := typeof.Field(i)
		sv := valueof.Field(i)
		def := sf.Tag.Get("default")
		if key := sf.Tag.Get("att"); key != "" {
			arr = attr(sf, sv, key, def, arr)
			continue
		} else if key := sf.Tag.Get("style"); key != "" {
			sty = style(sf, sv, key, def, sty)
			continue
		}

	}

	return strings.Join(arr, " ") + "  style=\"" + strings.Join(sty, " ") + "\" "
}

func attr(sf reflect.StructField, sv reflect.Value, key, def string, arr []string) []string {
	if sf.Type.Kind() == reflect.String {
		if sv.String() == "" {
			arr = append(arr, key+"=\""+def+"\"")
			return arr
		}
		arr = append(arr, key+"=\""+sv.String()+"\"")

	} else if sf.Type.Kind() == reflect.Int {
		if sv.Int() == 0 {
			arr = append(arr, ":"+key+"=\""+def+"\"")
			return arr
		}
		arr = append(arr, ":"+key+"=\""+strconv.Itoa(int(sv.Int()))+"\"")
	} else if sf.Type.Kind() == reflect.Bool {
		bb, _ := strconv.ParseBool(def)

		if sv.Bool() || bb {
			arr = append(arr, key)
			return arr
		}
		arr = append(arr, ":"+key+"='false'")
	} else if sf.Type.Kind() == reflect.Slice {

		if sv.Len() == 0 {
			arr = append(arr, ":"+key+"=\"[]\"")
			return arr
		}
		ar1 := make([]string, sv.Len())
		for i := 0; i < sv.Len(); i++ {
			v := sv.Index(i)
			ar1[i] = v.String()
		}
		arr = append(arr, ":"+key+"=\"['"+strings.Join(ar1, " ','")+"']\"")
		//sv.Elem().Slice()
	}
	return arr
}

func style(sf reflect.StructField, sv reflect.Value, key, def string, arr []string) []string {
	if sf.Type.Kind() == reflect.String {
		if sv.String() == "" {
			arr = append(arr, key+":"+def+";")
			return arr
		}
		arr = append(arr, key+"=\""+sv.String()+"\"")

	} else if sf.Type.Kind() == reflect.Int {
		if sv.Int() == 0 {
			arr = append(arr, ":"+key+"=\""+def+"\"")
			return arr
		}
		arr = append(arr, ":"+key+"=\""+strconv.Itoa(int(sv.Int()))+"\"")
	} else if sf.Type.Kind() == reflect.Bool {
		bb, _ := strconv.ParseBool(def)

		if sv.Bool() || bb {
			arr = append(arr, key)
			return arr
		}
		arr = append(arr, ":"+key+"='false'")
	} else if sf.Type.Kind() == reflect.Slice {

		if sv.Len() == 0 {
			arr = append(arr, ":"+key+"=\"[]\"")
			return arr
		}
		ar1 := make([]string, sv.Len())
		for i := 0; i < sv.Len(); i++ {
			v := sv.Index(i)
			ar1[i] = v.String()
		}
		arr = append(arr, ":"+key+"=\"['"+strings.Join(ar1, " ','")+"']\"")
		//sv.Elem().Slice()
	}
	return arr
}

func Iif[T int | string](b bool, s1, s2 T) T {
	if b {
		return s1
	}
	return s2
}
