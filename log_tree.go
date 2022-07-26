package log

import (
	"fmt"
	"reflect"
	"strconv"
)

func LogTree(name string, x interface{}) {
	var lines []string
	lines = append(lines, fmt.Sprintf("[Display] %s (%T):", name, x))
	display("", reflect.ValueOf(x), &lines, "", true)
	for _, line := range lines{
		fmt.Println(line)
	}
}

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64:
		return strconv.FormatComplex(v.Complex(), 'f', -1, 64)
	case reflect.Complex128:
		return strconv.FormatComplex(v.Complex(), 'f', -1, 128)
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func headIndent(end bool) string {
	if end {
		return "└─ "
	} else {
		return "├─ "
	}
}

func bodyIndent(end bool) string {
	if end {
		return "   "
	} else {
		return "│  "
	}
}

func display(path string, v reflect.Value, lines *[]string, base string, end bool) {
	switch v.Kind() {
	case reflect.Invalid:
		tmpString := fmt.Sprintf("%s = invalid", base+headIndent(end)+path)
		*lines = append(*lines, tmpString)
	case reflect.Slice, reflect.Array:
		if v.Len() == 0 {
			tmpString := fmt.Sprintf("%s = %s [Array: %s]", base+headIndent(end)+path, "[]", v.Type().Elem().Name())
			*lines = append(*lines, tmpString)
		} else {
			tmpString := fmt.Sprintf("%s [Array: %s]", base+headIndent(end)+path, v.Type().Elem().Name())
			*lines = append(*lines, tmpString)
			for i := 0; i < v.Len(); i++ {
				if v.Index(i).Kind() == reflect.Interface {
					if v.Index(i).IsNil() {
						display(fmt.Sprintf("[Interface] "+path+"[%d]", i), v.Index(i), lines, base+bodyIndent(end), i == v.Len()-1)
					} else {
						display(fmt.Sprintf("[Interface] "+path+"[%d]", i), v.Index(i).Elem(), lines, base+bodyIndent(end), i == v.Len()-1)
					}
				} else if v.Kind() == reflect.Ptr {
					if v.Index(i).IsNil() {
						display(fmt.Sprintf("[Ptr] "+path+"[%d]", i), v.Index(i), lines, base+bodyIndent(end), i == v.Len()-1)
					} else {
						display(fmt.Sprintf("[Ptr] "+path+"[%d]", i), v.Index(i).Elem(), lines, base+bodyIndent(end), i == v.Len()-1)
					}
				} else {
					display(fmt.Sprintf("%d", i), v.Index(i), lines, base+bodyIndent(end), i == v.Len()-1)
				}
			}
		}
	case reflect.Struct:
		if v.NumField() == 0 {
			tmpString := fmt.Sprintf("%s = %s [Struct: %s]", base+headIndent(end)+path, "{}", v.Type().Name())
			*lines = append(*lines, tmpString)
		} else {
			tmpString := fmt.Sprintf("%s [Struct: %s]", base+headIndent(end)+path, v.Type().Name())
			*lines = append(*lines, tmpString)
			for i := 0; i < v.NumField(); i++ {
				if v.Field(i).Kind() == reflect.Interface {
					if v.Field(i).IsNil() {
						display("[Interface] "+v.Type().Field(i).Name, v.Field(i), lines, base+bodyIndent(end), i == v.NumField()-1)
					} else {
						display("[Interface] "+v.Type().Field(i).Name, v.Field(i).Elem(), lines, base+bodyIndent(end), i == v.NumField()-1)
					}
				} else if v.Field(i).Kind() == reflect.Ptr {
					if v.Field(i).IsNil() {
						display("[Ptr] "+v.Type().Field(i).Name, v.Field(i), lines, base+bodyIndent(end), i == v.NumField()-1)
					} else {
						display(fmt.Sprintf("[Ptr %s] (*%s)", v.Type().Field(i).Type, v.Type().Field(i).Name), v.Field(i).Elem(), lines, base+bodyIndent(end), i == v.NumField()-1)
					}
				} else {
					display(v.Type().Field(i).Name, v.Field(i), lines, base+bodyIndent(end), i == v.NumField()-1)
				}
			}
		}
	case reflect.Map:
		kType := v.Type().Key().Name()
		if len(kType) == 0 {
			kType = "?" // case: interface{}
		}
		vType := v.Type().Elem().Name()
		if len(vType) == 0 {
			vType = "?" // case: interface{}
		}
		if v.Len() == 0 {
			tmpString := fmt.Sprintf("%s = %s [Map <%s, %s>]", base+headIndent(end)+path, "{}", kType, vType)
			*lines = append(*lines, tmpString)
		} else {
			tmpString := fmt.Sprintf("%s  [Map <%s, %s>]", base+headIndent(end)+path, kType, vType)
			*lines = append(*lines, tmpString)
			res := reflect.MakeMap(v.Type())
			for index, _key := range v.MapKeys() {
				key := _key.Convert(res.Type().Key())
				value := v.MapIndex(_key)
				if key.Kind() != reflect.Struct && key.Kind() != reflect.Slice && key.Kind() != reflect.Array && key.Kind() != reflect.Map &&  key.Kind() != reflect.Interface &&  key.Kind() != reflect.Ptr {
					display(formatAtom(key), value, lines, base+bodyIndent(end), index == v.Len()-1)
				} else {
					if key.IsNil() {
						display("[MapKey]"+formatAtom(key), key, lines, base+bodyIndent(end), false)
					} else {
						display("[MapKey]"+formatAtom(key), key.Elem(), lines, base+bodyIndent(end), false)
					}
					if value.IsNil() {
						display("[MapValue]"+formatAtom(value), value, lines, base+bodyIndent(end), index == v.Len()-1)
					} else {
						display("[MapValue]"+formatAtom(value), value.Elem(), lines, base+bodyIndent(end), index == v.Len()-1)
					}
				}
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			tmpString := fmt.Sprintf("%s = nil [Ptr: %s]", base+headIndent(end)+path, v.Type().Elem().Name())
			*lines = append(*lines, tmpString)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), lines, base+bodyIndent(end), true)
		}
	case reflect.Interface:
		if v.IsNil() {
			tmpString := fmt.Sprintf("%s = nil", base+headIndent(end)+path)
			*lines = append(*lines, tmpString)
		} else {
			tmpString := fmt.Sprintf("%s.type = %s [Interface]", base+headIndent(end)+path, v.Elem().Type())
			*lines = append(*lines, tmpString)
			display(path+".value", v.Elem(), lines, base+bodyIndent(end), true)
		}
	default: // basic types, channels, funcs
		tmpString := fmt.Sprintf("%s = %s (%s)", base+headIndent(end)+path, formatAtom(v), v.Kind())
		*lines = append(*lines, tmpString)
	}
}
