package koiApi

// DateFormat represents the date format preference for a user.
type DateFormat string

const (
	DateFormatDMYSlash DateFormat = "d/m/Y"
	DateFormatMDYSlash DateFormat = "m/d/Y"
	DateFormatYMDSlash DateFormat = "Y/m/d"
	DateFormatDMYDash  DateFormat = "d-m-Y"
	DateFormatMDYDash  DateFormat = "m-d-Y"
	DateFormatYMDDash  DateFormat = "Y-m-d" // Default
)

func (df DateFormat) String() string {
	return string(df)
}
