package koiApi

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// printStruct generically prints the fields of a struct using reflection, aligning values at the same left margin,
// skipping fields tagged with omitempty if their values would be omitted in JSON marshaling. It prints a formatted
// string (using printf-style parameters) at the beginning, and indents all field lines by 4 spaces.
func printStruct(v interface{}, format string, args ...interface{}) (int, error) {
	if v == nil {
		fmt.Println("<nil>")
		return 0, nil
	}

	// Get the value and dereference if it's a pointer
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			fmt.Println("<nil>")
			return 0, nil
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return 0, fmt.Errorf("expected a struct or pointer to struct, got %v", val.Kind())
	}

	typ := val.Type()
	numFields := val.NumField()

	// Print the formatted header
	if format != "" {
		fmt.Printf(format, args...)
	} else {
		fmt.Printf("%s\n", typ.Name())
	}

	if numFields == 0 {
		fmt.Printf("    %s: <empty>\n", typ.Name())
		return 0, nil
	}

	// Calculate max field name length for alignment
	maxLen := 0
	for i := 0; i < numFields; i++ {
		name := typ.Field(i).Name
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	// Print fields with aligned values, indented by 4 spaces
	printedFields := 0
	for i := 0; i < numFields; i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		name := fieldType.Name
		jsonTag := fieldType.Tag.Get("json")
		omitEmpty := strings.Contains(jsonTag, ",omitempty")

		// Check if the field would be omitted in JSON
		shouldSkip := false
		if omitEmpty {
			switch field.Kind() {
			case reflect.Ptr:
				shouldSkip = field.IsNil()
			case reflect.Slice:
				shouldSkip = field.IsNil() || field.Len() == 0
			case reflect.String:
				shouldSkip = field.String() == ""
			case reflect.Int:
				shouldSkip = field.Int() == 0
			case reflect.Bool:
				shouldSkip = !field.Bool()
			case reflect.Struct:
				if field.Type() == reflect.TypeOf(time.Time{}) {
					shouldSkip = field.Interface().(time.Time).IsZero()
				}
			}
		}

		if shouldSkip {
			continue
		}

		prefix := "    " + typ.Name() + "." + name
		padding := strings.Repeat(" ", maxLen-len(name)+1)

		switch field.Kind() {
		case reflect.Ptr:
			if field.IsNil() {
				fmt.Printf("%s:%s<nil>\n", prefix, padding)
			} else {
				switch field.Type().Elem().Kind() {
				case reflect.String:
					fmt.Printf("%s:%s%s\n", prefix, padding, field.Elem().String())
				case reflect.Int:
					fmt.Printf("%s:%s%d\n", prefix, padding, field.Elem().Int())
				case reflect.Bool:
					fmt.Printf("%s:%s%t\n", prefix, padding, field.Elem().Bool())
				case reflect.Struct:
					if field.Type().Elem() == reflect.TypeOf(time.Time{}) {
						fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().Interface())
					} else {
						fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().String())
					}
				default:
					fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().Interface())
				}
			}
		case reflect.Slice:
			if field.IsNil() {
				fmt.Printf("%s:%s[]\n", prefix, padding)
			} else {
				fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
			}
		default:
			switch field.Kind() {
			case reflect.String:
				fmt.Printf("%s:%s%s\n", prefix, padding, field.String())
			case reflect.Int:
				fmt.Printf("%s:%s%d\n", prefix, padding, field.Int())
			case reflect.Bool:
				fmt.Printf("%s:%s%t\n", prefix, padding, field.Bool())
			case reflect.Struct:
				if field.Type() == reflect.TypeOf(time.Time{}) {
					fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
				} else {
					fmt.Printf("%s:%s%v\n", prefix, padding, field.String())
				}
			default:
				fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
			}
		}
		printedFields++
	}

	return printedFields, nil
}

// Print methods for each struct type
func (a Album) Print(format string, args ...interface{}) (int, error) {
	return printStruct(a, format, args...)
}

func (c ChoiceList) Print(format string, args ...interface{}) (int, error) {
	return printStruct(c, format, args...)
}

func (c Collection) Print(format string, args ...interface{}) (int, error) {
	return printStruct(c, format, args...)
}

func (d Datum) Print(format string, args ...interface{}) (int, error) {
	return printStruct(d, format, args...)
}

func (f Field) Print(format string, args ...interface{}) (int, error) {
	return printStruct(f, format, args...)
}

func (i Inventory) Print(format string, args ...interface{}) (int, error) {
	return printStruct(i, format, args...)
}

func (i Item) Print(format string, args ...interface{}) (int, error) {
	return printStruct(i, format, args...)
}

func (l Loan) Print(format string, args ...interface{}) (int, error) {
	return printStruct(l, format, args...)
}

func (l Log) Print(format string, args ...interface{}) (int, error) {
	return printStruct(l, format, args...)
}

func (p Photo) Print(format string, args ...interface{}) (int, error) {
	return printStruct(p, format, args...)
}

func (t Tag) Print(format string, args ...interface{}) (int, error) {
	return printStruct(t, format, args...)
}

func (tc TagCategory) Print(format string, args ...interface{}) (int, error) {
	return printStruct(tc, format, args...)
}

func (t Template) Print(format string, args ...interface{}) (int, error) {
	return printStruct(t, format, args...)
}

func (u User) Print(format string, args ...interface{}) (int, error) {
	return printStruct(u, format, args...)
}

func (w Wish) Print(format string, args ...interface{}) (int, error) {
	return printStruct(w, format, args...)
}

func (w Wishlist) Print(format string, args ...interface{}) (int, error) {
	return printStruct(w, format, args...)
}
