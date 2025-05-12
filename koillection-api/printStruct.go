package koiApi

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// printStruct generically prints the fields of a struct using reflection, aligning values at the same left margin.
func printStruct(v interface{}) (int, error) {
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
	if numFields == 0 {
		fmt.Printf("%s: <empty>\n", typ.Name())
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

	// Print fields with aligned values
	for i := 0; i < numFields; i++ {
		field := val.Field(i)
		name := typ.Field(i).Name
		prefix := typ.Name() + "." + name
		padding := strings.Repeat(" ", maxLen-len(name)+1)

		if field.Kind() == reflect.Ptr {
			if field.IsNil() {
				fmt.Printf("%s:%s<nil>\n", prefix, padding)
			} else {
				// Safely dereference pointer fields
				switch field.Type().Elem().Kind() {
				case reflect.String:
					fmt.Printf("%s:%s%s\n", prefix, padding, field.Elem().String())
				case reflect.Int:
					fmt.Printf("%s:%s%d\n", prefix, padding, field.Elem().Int())
				case reflect.Bool:
					fmt.Printf("%s:%s%t\n", prefix, padding, field.Elem().Bool())
				case reflect.Struct:
					// Handle time.Time and other structs
					if field.Type().Elem() == reflect.TypeOf(time.Time{}) {
						fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().Interface())
					} else {
						fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().String())
					}
				default:
					fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().Interface())
				}
			}
		} else if field.Kind() == reflect.Slice {
			if field.IsNil() {
				fmt.Printf("%s:%s[]\n", prefix, padding)
			} else {
				fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
			}
		} else {
			// Handle non-pointer, non-slice fields
			switch field.Kind() {
			case reflect.String:
				fmt.Printf("%s:%s%s\n", prefix, padding, field.String())
			case reflect.Int:
				fmt.Printf("%s:%s%d\n", prefix, padding, field.Int())
			case reflect.Bool:
				fmt.Printf("%s:%s%t\n", prefix, padding, field.Bool())
			case reflect.Struct:
				// Handle time.Time and other structs
				if field.Type() == reflect.TypeOf(time.Time{}) {
					fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
				} else {
					fmt.Printf("%s:%s%v\n", prefix, padding, field.String())
				}
			default:
				fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
			}
		}
	}

	return numFields, nil
}

// Print methods for each struct type
func (a Album) Print() (int, error) {
	return printStruct(a)
}

func (c ChoiceList) Print() (int, error) {
	return printStruct(c)
}

func (c Collection) Print() (int, error) {
	return printStruct(c)
}

func (d Datum) Print() (int, error) {
	return printStruct(d)
}

func (f Field) Print() (int, error) {
	return printStruct(f)
}

func (i Inventory) Print() (int, error) {
	return printStruct(i)
}

func (i Item) Print() (int, error) {
	return printStruct(i)
}

func (l Loan) Print() (int, error) {
	return printStruct(l)
}

func (l Log) Print() (int, error) {
	return printStruct(l)
}

func (p Photo) Print() (int, error) {
	return printStruct(p)
}

func (t Tag) Print() (int, error) {
	return printStruct(t)
}

func (tc TagCategory) Print() (int, error) {
	return printStruct(tc)
}

func (t Template) Print() (int, error) {
	return printStruct(t)
}

func (u User) Print() (int, error) {
	return printStruct(u)
}

func (w Wish) Print() (int, error) {
	return printStruct(w)
}

func (w Wishlist) Print() (int, error) {
	return printStruct(w)
}
