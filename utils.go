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
