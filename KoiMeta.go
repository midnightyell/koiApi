package koiApi

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	caller "gitea.local/smalloy/caller"
)

var basePathForType = map[string]string{
	"album":      "/api/albums",
	"choiceList": "/api/choice_lists",
	"collection": "/api/collections",
	"item":       "/api/items",
	"datum":      "/api/data",
	"loan":       "/api/loans",
	"photo":      "/api/photos",
	"tag":        "/api/tags",
	"template":   "/api/templates",
	"user":       "/api/users",
	"wish":       "/api/wishes",
	"wishlist":   "/api/wishlists",
}

type koiOp struct {
	op   string
	path string
}

type KoiObject interface {
	GetID() string
	IRI() string
	Validate() error
}

func IRI[T KoiObject](o T) string {
	return fmt.Sprintf("%s/%s", BaseObjPath(o), o.GetID())
}

func BaseObjPath[T KoiObject](o T) string {
	typeName := strings.ToLower(reflect.TypeOf(o).Elem().Name())
	parts := strings.Split(typeName, ".")
	typeName = parts[len(parts)-1]
	return basePathForType[typeName]
}

func KoiPathForOp(obj KoiObject) (*koiOp, error) {
	basePath := ""

	if obj != nil {
		basePath = BaseObjPath(obj) // Get the base path for the type
	}

	// Get caller name
	op := strings.ToLower(caller.ParentFunc())

	retval := koiOp{}

	switch op {
	case "create":
		retval = koiOp{op: http.MethodPost, path: basePath}
	case "get":
		retval = koiOp{op: http.MethodGet, path: fmt.Sprintf("%s/%s", basePath, GetID(obj))}
	case "list":
		retval = koiOp{op: http.MethodGet, path: basePath}
	case "update":
		retval = koiOp{op: http.MethodPut, path: fmt.Sprintf("%s/%s", basePath, obj.GetID())}
	case "delete":
		retval = koiOp{op: http.MethodDelete, path: fmt.Sprintf("%s/%s", basePath, obj.GetID())}
	case "patch":
		retval = koiOp{op: http.MethodPatch, path: fmt.Sprintf("%s/%s", basePath, obj.GetID())}
	case "listphotos":
		retval = koiOp{op: http.MethodGet, path: fmt.Sprintf("%s/%s/photos", basePath, obj.GetID())}
	case "listchildren":
		retval = koiOp{op: http.MethodGet, path: fmt.Sprintf("%s/%s/children", basePath, obj.GetID())}
	case "uploadimage":
		retval = koiOp{op: http.MethodPost, path: fmt.Sprintf("%s/%s/image", basePath, obj.GetID())}
	case "uploadimagefromfile":
		retval = koiOp{op: http.MethodPost, path: fmt.Sprintf("%s/%s/image", basePath, obj.GetID())}
	case "getparent":
		retval = koiOp{op: http.MethodGet, path: fmt.Sprintf("%s/%s/parent", basePath, obj.GetID())}
	case "uploadfile":
		retval = koiOp{op: http.MethodPost, path: fmt.Sprintf("%s/%s/file", basePath, obj.GetID())}
	case "uploadfilefromfile":
		retval = koiOp{op: http.MethodPost, path: fmt.Sprintf("%s/%s/file", basePath, obj.GetID())}
	case "getdefaulttemplate":
		retval = koiOp{op: http.MethodGet, path: fmt.Sprintf("%s/%s/items_default_template", basePath, obj.GetID())}
	// Add more cases as needed
	default:
		return &koiOp{}, fmt.Errorf("unknown operation: %s", op)
	}

	return &retval, nil
}

func GetID(o KoiObject) string {
	return o.GetID()
}

func Create[T KoiObject](o T) (T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == POST
	if op == http.MethodPost {
		var resp T
		err := c.postResource(path, o, &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func Delete[T KoiObject](o T) error {
	result, err := KoiPathForOp(o)
	if err != nil {
		return fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == DELETE
	if op == http.MethodDelete {
		return c.deleteResource(path)
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func Get[T KoiObject](o T) (T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == GET
	if op == http.MethodGet {
		var resp T
		err := c.getResource(path, &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func GetParent[T KoiObject](o T) (T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == GET
	if op == http.MethodGet {
		var resp T
		err := c.getResource(path, &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func List[T KoiObject](o T) ([]T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return nil, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == GET
	if op == http.MethodGet {
		var items []T
		err := c.listResources(path, &items)
		return items, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return nil, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func ListChildren[T KoiObject](o T) ([]T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return nil, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == GET
	if op == http.MethodGet {
		var items []T
		err := c.listResources(path, &items)
		return items, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return nil, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func ListPhotos[T KoiObject](o T) ([]*Photo, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return nil, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == GET
	if op == http.MethodGet {
		var items []*Photo
		err := c.listResources(path, &items)
		return items, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return nil, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func Patch[T KoiObject](o T) (T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == PATCH
	if op == http.MethodPatch {
		var resp T
		err := c.patchResource(path, o, &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}
func Update[T KoiObject](o T) (T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == PUT
	if op == http.MethodPut {
		var resp T
		err := c.putResource(path, o, &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func UploadImage[T KoiObject](o T, file []byte) (T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == POST
	if op == http.MethodPost {
		var resp T
		err := c.uploadFile(path, file, "file", &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func UploadImageFromFile[T KoiObject](o T, filename string) (T, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == POST
	if op == http.MethodPost {
		file, err := os.ReadFile(filename)
		if err != nil {
			return o, fmt.Errorf("failed to read file %s: %w", filename, err)
		}
		var resp T
		err = c.uploadFile(path, file, "file", &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func GetDefaultTemplate[T KoiObject](o T) (*Template, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return nil, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == GET
	if op == http.MethodGet {
		var resp *Template
		err := c.getResource(path, &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return nil, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}
