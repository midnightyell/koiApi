package koiApi

// Visibility represents the visibility level of a resource.
type Visibility string // Read and write

const (
	VisibilityPublic   Visibility = "public" // Default for most resources
	VisibilityInternal Visibility = "internal"
	VisibilityPrivate  Visibility = "private" // Default for User
)

func (v Visibility) String() string {
	return string(v)
}
