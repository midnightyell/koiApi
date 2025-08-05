package koiApi

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strings"

	caller "gitea.local/smalloy/caller-utils"
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
	op     string
	path   string
	caller string
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
	fn := strings.ToLower(caller.GrandparentFunc())
	retval := koiOp{caller: fn}

	switch fn {
	case "create":
		retval = koiOp{caller: fn, op: http.MethodPost, path: basePath}
	case "delete":
		retval = koiOp{caller: fn, op: http.MethodDelete, path: fmt.Sprintf("%s/%s", basePath, obj.GetID())}
	case "get":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s", basePath, GetID(obj))}
	case "getcollection":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/collection", basePath, obj.GetID())}
	case "getalbum":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/album", basePath, obj.GetID())}
	case "getdefaulttemplate":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/items_default_template", basePath, obj.GetID())}
	case "getitem":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/item", basePath, obj.GetID())}
	case "getparent":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/parent", basePath, obj.GetID())}
	case "gettagcategory":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/category", basePath, obj.GetID())}
	case "gettemplate":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/template", basePath, obj.GetID())}
	case "list":
		retval = koiOp{caller: fn, op: http.MethodGet, path: basePath}
	case "listchildren":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/children", basePath, obj.GetID())}
	case "listdata":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/data", basePath, obj.GetID())}
	case "listfields":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/fields", basePath, obj.GetID())}
	case "listitems":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/items", basePath, obj.GetID())}
	case "listloans":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/loans", basePath, obj.GetID())}
	case "listphotos":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/photos", basePath, obj.GetID())}
	case "listrelateditems":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/related_items", basePath, obj.GetID())}
	case "listtags":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/tags", basePath, obj.GetID())}
	case "listwishes":
		retval = koiOp{caller: fn, op: http.MethodGet, path: fmt.Sprintf("%s/%s/wishes", basePath, obj.GetID())}
	case "patch":
		retval = koiOp{caller: fn, op: http.MethodPatch, path: fmt.Sprintf("%s/%s", basePath, obj.GetID())}
	case "update":
		retval = koiOp{caller: fn, op: http.MethodPut, path: fmt.Sprintf("%s/%s", basePath, obj.GetID())}
	case "uploadfile":
		retval = koiOp{caller: fn, op: http.MethodPost, path: fmt.Sprintf("%s/%s/file", basePath, obj.GetID())}
	case "uploadfilefromfile":
		retval = koiOp{caller: fn, op: http.MethodPost, path: fmt.Sprintf("%s/%s/file", basePath, obj.GetID())}
	case "uploadimage":
		retval = koiOp{caller: fn, op: http.MethodPost, path: fmt.Sprintf("%s/%s/image", basePath, obj.GetID())}
	case "uploadimagefromfile":
		retval = koiOp{caller: fn, op: http.MethodPost, path: fmt.Sprintf("%s/%s/image", basePath, obj.GetID())}
	case "uploadvideo":
		retval = koiOp{caller: fn, op: http.MethodPost, path: fmt.Sprintf("%s/%s/video", basePath, obj.GetID())}
	case "uploadvideofromfile":
		retval = koiOp{caller: fn, op: http.MethodPost, path: fmt.Sprintf("%s/%s/video", basePath, obj.GetID())}
	// Add more cases as needed
	default:
		return &koiOp{caller: fn}, fmt.Errorf("unknown operation: %s for type %T", fn, obj)
	}

	return &retval, nil
}

func GetID(o KoiObject) string {
	return o.GetID()
}

func doCreate[T KoiObject](o T) (T, error) {
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

func doDelete[T KoiObject](o T) error {
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

func doGet[T KoiObject](o T) (any, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return o, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == GET
	if op == http.MethodGet {
		switch result.caller {
		case "getcollection":
			// Datum, Item,
			var resp *Collection
			err := c.getResource(path, &resp)
			return resp, err
		case "getdefaulttemplate":
			// Collection
			var resp *Template
			err := c.getResource(path, &resp)
			return resp, err
		case "gettemplate":
			// Fields
			var resp *Template
			err := c.getResource(path, &resp)
			return resp, err
		case "getalbum":
			// Photo
			var resp *Album
			err := c.getResource(path, &resp)
			return resp, err
		case "getitem":
			// Data, Loan
			var resp *Item
			err := c.getResource(path, &resp)
			return resp, err
		case "gettagcategory":
			// TagCategory
			var resp *TagCategory
			err := c.getResource(path, &resp)
			return resp, err
		default:
			var resp T
			err := c.getResource(path, &resp)
			return resp, err
		}
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func doList[T KoiObject](o T) (any, error) {
	/*
		List can sometimes return a different type than T, such as a list of Photos for an Album,
		or a list of Items for a Collection.
		So we return an interface{} and let the caller handle the type assertion.
	*/
	result, err := KoiPathForOp(o)
	if err != nil {
		return nil, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()
	// if op == GET
	if op == http.MethodGet {

		switch result.caller {
		case "listdata":
			// Collection, Item
			var objs []*Datum
			err := c.listResources(path, &objs)
			return objs, err

		case "listfields":
			// Templates
			var objs []*Field
			err := c.listResources(path, &objs)
			return objs, err

		case "listitems":
			// Collection, Tags
			var objs []*Item
			err := c.listResources(path, &objs)
			return objs, err

		case "listloans":
			// Collection, Tags
			var objs []*Loan
			err := c.listResources(path, &objs)
			return objs, err

		case "listphotos":
			// Album
			var objs []*Photo
			err := c.listResources(path, &objs)
			return objs, err

		case "listtags":
			// Item, TagCategory
			var objs []*Field
			err := c.listResources(path, &objs)
			return objs, err

		case "listwishes":
			// Wishlist
			var objs []*Wish
			err := c.listResources(path, &objs)
			return objs, err

		default:
			var objs []T
			err := c.listResources(path, &objs)
			return objs, err
		}
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return nil, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func doPatch[T KoiObject](o T) (T, error) {
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

func doUpdate[T KoiObject](o T) (T, error) {
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

func doUpload[T KoiObject](o T, magic string, file []byte) (T, error) {
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
		err := c.uploadFile(path, file, magic, &resp)
		return resp, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return o, fmt.Errorf("operation %s not implemented for type %T in %s", op, o, caller.ThisFunc())
}

func doUploadFromFile[T KoiObject](o T, magic string, fname string) (T, error) {
	file, err := os.ReadFile(fname)
	if err != nil {
		return o, fmt.Errorf("failed to read file %s: %w", fname, err)
	}
	return doUpload(o, magic, file)
}
