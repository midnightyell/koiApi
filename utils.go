package koiApi

import (
	"fmt"
	"reflect"
	"strings"

	caller "gitea.local/smalloy/caller-utils"
)

func validationErrors(errs *[]string) error {
	if errs != nil && len(*errs) > 0 {
		return fmt.Errorf("%s failed: %s", caller.ParentFunc(false), strings.Join(*errs, "; "))
	}
	return nil
}

func validateVisibility[T KoiObject](a T, errs *[]string) {
	v := reflect.ValueOf(a)
	if v.Kind() == reflect.Struct {
		if field := v.FieldByName("Visibility"); field.IsValid() {
			switch field.Interface() {
			case VisibilityPublic, VisibilityInternal, VisibilityPrivate:
			default:
				*errs = append(*errs, fmt.Sprintf("invalid visibility: %s; must be public, internal, or private", field.Interface()))
			}
		}
	}
	return
}

// getarg will return s[idx] if s[idx] exists, def otherwise
//
//	e.g. func foo(arg ...int) {var := getarg(100, arg)}
//
// Where var gets a (0th) value passed to foo if present, 100 otherwise
func getArg[T any](def T, s []T, idx ...int) T {
	index := 0
	if len(idx) > 0 {
		index = idx[0]
	}
	if len(s) > index {
		return s[index]
	}
	return def
}

func indentChars(num int, args ...string) string {
	indentStr := getArg(" ", args, 0)
	prefix := getArg("", args, 1)
	return fmt.Sprintf("%s%s", prefix, strings.Repeat(indentStr, num))
}

func lastChars(s string, n int) string {
	return s[len(s)-n:]
}
