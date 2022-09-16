// Package params provides a reflection-based parser for URL parameters.
// Package params 为 URL 参数提供了一个基于反射的解析器。
package params

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//!+unpack

// Unpack populates the fields of the struct pointed to by ptr
// 解包填充 ptr 指向的结构的字段
// from the HTTP request parameters in req.
// 来自 req 中的 HTTP 请求参数。
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// Build map of fields keyed by effective name.
	// 构建以有效名称为键的字段映射。
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem() // the struct variable // 结构变量
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) //a reflect.StructField // 一个反射 StructField
		tag := fieldInfo.Tag           //  // a reflect.StructTag // 一个反射 StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// Update struct field for each parameter in the request.
	// 更新请求中每个参数的结构字段。
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // ignore unrecognized HTTP parameters // 忽略无法识别的 HTTP 参数
		}
		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem().Elem())
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

//!-Unpack

//!+populate
func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

//!-populate
