package koiApi

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	caller "gitea.local/smalloy/caller"
)

// type KoiObject interface {
// 	GetID() string
// 	Create(c *koiClient, o *interface{}) (interface{}, error)
// 	Get(c *koiClient, ID) (interface{}, error)
// 	List(c *koiClient) ([]interface{}, error)
// 	Update(c *koiClient, ID, interface{}) (interface{}, error)
// 	Patch(c *koiClient, ID, interface{}) (interface{}, error)
// 	Delete(c *koiClient, ID) error
// 	ListChildren(c *koiClient, ID) ([]interface{}, error)
// 	UploadImage(c *koiClient, ID ID, file []byte) (interface{}, error)
// 	UploadImageFromFile(c *koiClient, ID ID, filePath string) (interface{}, error)
// 	GetParent(c *koiClient, ID ID) (interface{}, error)
// 	UploadFile(c *koiClient, ID ID, file []byte) (interface{}, error)
// 	UploadFileFromFile(c *koiClient, ID ID, filePath string) (interface{}, error)
// 	Validate(obj *interface{}) error
// }

var basePathForType = map[string]string{
	"Album":      "/api/albums",
	"ChoiceList": "/api/choice_lists",
	"Collection": "/api/collections",
	"Item":       "/api/items",
	"Datum":      "/api/data",
	"Loan":       "/api/loans",
	"Photo":      "/api/photos",
	"Tag":        "/api/tags",
	"Template":   "/api/templates",
	"User":       "/api/users",
	"Wish":       "/api/wishes",
	"Wishlist":   "/api/wishlists",
}

type koiOp struct {
	op   string
	path string
}

type KoiObject interface {
	GetID() string
}

func KoiPathForOp(obj KoiObject) (*koiOp, error) {

	basePath := ""

	if obj != nil {
		// if objname not found, basePath is nil
		typeName := reflect.TypeOf(obj).Elem().Name()
		//fmt.Println("Type name:", typeName)
		parts := strings.Split(typeName, ".")
		typeName = parts[len(parts)-1]
		basePath = basePathForType[typeName]
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
	default:
		return &koiOp{}, fmt.Errorf("unknown operation: %s", op)
	}

	return &retval, nil
}

func GetID(o KoiObject) string {
	return o.GetID()
}

func Create(o KoiObject) (KoiObject, error) {
	result, err := KoiPathForOp(o)
	if err != nil {
		return nil, fmt.Errorf("failed to get operation path: %w", err)
	}
	op := result.op
	path := result.path

	c := GetClient()

	// if op == POST
	if op == http.MethodPost {
		err := c.postResource(path, o, &o)
		return o, err
	}
	fmt.Printf("FAILED: %20s %8s %s\n", caller.ThisFunc(), result.op, result.path)
	return nil, fmt.Errorf("operation %s not implemented for type %T", op, o)
}
