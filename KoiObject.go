package koiApi

/*

func (a *Object)
() (*Object, error) {

{
	resp, err := Get(obj)
	return resp.(T), err
}


	func (a *Object) Create
	func (a *Object) Delete
	func (a *Object) Get
	func (a *Object) GetCollection
	func (a *Object) GetAlbum
	func (a *Object) GetDefaultTemplate
	func (a *Object) GetItem
	func (a *Object) GetParent
	func (a *Object) GetTagCategory
	func (a *Object) GetTemplate
	func (a *Object) List
	func (a *Object) ListChildren
	func (a *Object) ListData
	func (a *Object) ListFields
	func (a *Object) ListItems
	func (a *Object) ListLoans
	func (a *Object) ListPhotos
	func (a *Object) ListRelatedItems
	func (a *Object) ListTags
	func (a *Object) ListWishes
	func (a *Object) Patch
	func (a *Object) Update
	func (a *Object) UploadFile
	func (a *Object) UploadFileFromFile
	func (a *Object) UploadImage
	func (a *Object) UploadImageFromFile
	func (a *Object) UploadVideo
	func (a *Object) UploadVideoFromFile
*/

func Create[T KoiObject](obj T) (T, error) {
	resp, err := doCreate(obj)
	return resp.(T), err
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

func List[T KoiObject](obj T) ([]*T, error) {
	resp, err := doList(obj)
	return resp.([]*T), err
}

func ListChildren[T KoiObject](obj T) ([]*T, error) {
	resp, err := doList(obj)
	return resp.([]*T), err
}

func ListData[T KoiObject](obj T) ([]*Datum, error) {
	resp, err := doList(obj)
	return resp.(T), err
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
	return resp.(*[]Photo), err
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
	return resp.(T), err
}

func Update[T KoiObject](obj T) (T, error) {
	resp, err := doUpdate(obj)
	return resp.(T), err
}

func UploadFile[T KoiObject](obj T, file []byte) (*Datum, error) {
	resp, err := doUpload(obj, file)
	return resp.(*Datum), err
}

func UploadFileFromFile[T KoiObject](obj T, filename string) (*Datum, error) {
	resp, err := doUploadFromFile(obj, filename)
	return resp.(*Datum), err
}

func UploadImage[T KoiObject](obj T, file []byte) (T, error) {
	resp, err := doUpload(obj, file)
	return resp.(T), err
}

func UploadImageFromFile[T KoiObject](obj T, filename string) (T, error) {
	resp, err := doUploadFromFile(obj, filename)
	return resp.(T), err
}

func UploadVideo[T KoiObject](obj T, file []byte) (T, error) {
	resp, err := doUpload(obj, "fileVideo", file)
	return resp.(T), err
}

func UploadVideoFromFile[T KoiObject](obj T, filename string) (T, error) {
	resp, err := doUploadFromFile(obj, "fileVideo", filename)
	return resp.(T), err
}
