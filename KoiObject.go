package koiApi

func Create[T KoiObject](obj T) (T, error) {
	resp, err := doCreate(obj)
	return resp, err
}

func Delete[T KoiObject](obj T) error {
	return doDelete(obj)
}

func Get[T KoiObject](obj T) (T, error) {
	resp, err := doGet(obj)
	return resp.(T), err
}

func GetCollection[T KoiObject](obj T) (*Collection, error) {
	resp, err := doGet(obj)
	return resp.(*Collection), err
}

func GetAlbum[T KoiObject](obj T) (*Album, error) {
	resp, err := doGet(obj)
	return resp.(*Album), err
}

func GetDefaultTemplate[T KoiObject](obj T) (*Template, error) {
	resp, err := doGet(obj)
	return resp.(*Template), err
}

func GetItem[T KoiObject](obj T) (*Item, error) {
	resp, err := doGet(obj)
	return resp.(*Item), err
}

func GetParent[T KoiObject](obj T) (T, error) {
	resp, err := doGet(obj)
	return resp.(T), err
}

func GetTagCategory[T KoiObject](obj T) (*TagCategory, error) {
	resp, err := doGet(obj)
	return resp.(*TagCategory), err
}

func GetTemplate[T KoiObject](obj T) (*Template, error) {
	resp, err := doGet(obj)
	return resp.(*Template), err
}

func List[T KoiObject](obj T) ([]T, error) {
	resp, err := doList(obj)
	return resp.([]T), err
}

func ListChildren[T KoiObject](obj T) ([]T, error) {
	resp, err := doList(obj)
	return resp.([]T), err
}

func ListData[T KoiObject](obj T) ([]*Datum, error) {
	resp, err := doList(obj)
	return resp.([]*Datum), err
}

func ListFields[T KoiObject](obj T) ([]*Field, error) {
	resp, err := doList(obj)
	return resp.([]*Field), err
}

func ListItems[T KoiObject](obj T) ([]*Item, error) {
	resp, err := doList(obj)
	return resp.([]*Item), err
}

func ListLoans[T KoiObject](obj T) ([]*Loan, error) {
	resp, err := doList(obj)
	return resp.([]*Loan), err
}

func ListPhotos[T KoiObject](obj T) ([]*Photo, error) {
	resp, err := doList(obj)
	return resp.([]*Photo), err
}

func ListRelatedItems[T KoiObject](obj T) ([]*Item, error) {
	resp, err := doList(obj)
	return resp.([]*Item), err
}

func ListTags[T KoiObject](obj T) ([]*Tag, error) {
	resp, err := doList(obj)
	return resp.([]*Tag), err
}

func ListWishes[T KoiObject](obj T) ([]*Wish, error) {
	resp, err := doList(obj)
	return resp.([]*Wish), err
}

func Patch[T KoiObject](obj T) (T, error) {
	resp, err := doPatch(obj)
	return resp, err
}

func Update[T KoiObject](obj T) (T, error) {
	resp, err := doUpdate(obj)
	return resp, err
}

func UploadFile[T KoiObject](obj T, file []byte) (any, error) {
	resp, err := doUpload(obj, "fileFile", file)
	return resp, err
}

func UploadFileFromFile[T KoiObject](obj T, filename string) (any, error) {
	resp, err := doUploadFromFile(obj, "fileFile", filename)
	return resp, err
}

func UploadImage[T KoiObject](obj T, file []byte) (any, error) {
	resp, err := doUpload(obj, "fileImage", file)
	return resp, err
}

func UploadImageFromFile[T KoiObject](obj T, filename string) (any, error) {
	resp, err := doUploadFromFile(obj, "fileImage", filename)
	return resp, err
}

func UploadVideo[T KoiObject](obj T, file []byte) (any, error) {
	resp, err := doUpload(obj, "fileVideo", file)
	return resp, err
}

func UploadVideoFromFile[T KoiObject](obj T, filename string) (any, error) {
	resp, err := doUploadFromFile(obj, "fileVideo", filename)
	return resp, err
}
