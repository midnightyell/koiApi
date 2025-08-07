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

func (i *Inventory) Summary() string {
	return fmt.Sprintf("%-40s %s", i.Name, i.ID)
}

func (i *Item) Summary() string {
	return fmt.Sprintf("%8.8s   %s", i.ID[len(i.ID)-8:], i.Name)
}

func (l *Loan) Summary() string {
	return fmt.Sprintf("%-40s %s", l.LentTo, l.ID)
}

func (l *Log) Summary() string {
	return fmt.Sprintf("%-40s %s", l.ObjectLabel, l.ID)
}

func (m *Metrics) Summary() string {
	return fmt.Sprintf("%-40s %s", string(*m), "")
}

func (p *Photo) Summary() string {
	return fmt.Sprintf("%-40s %s", p.Title, p.ID)
}

func (t *Tag) Summary() string {
	return fmt.Sprintf("%-40s %s", t.Label, t.ID)
}

func (tc *TagCategory) Summary() string {
	return fmt.Sprintf("%-40s %s", tc.Label, tc.ID)
}

func (t *Template) Summary() string {
	return fmt.Sprintf("%-40s %s", t.Name, t.ID)
}

func (u *User) Summary() string {
	return fmt.Sprintf("%-40s %s", u.Username, u.ID)
}

func (w *Wish) Summary() string {
	return fmt.Sprintf("%-40s %s", w.Name, w.ID)
}

func (w *Wishlist) Summary() string {
	return fmt.Sprintf("%-40s %s", w.Name, w.ID)
}
