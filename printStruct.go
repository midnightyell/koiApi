package koiApi

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
)

// getDefaultCurrency returns the local currency based on the machine's locale or "USD" if undetermined.
func getDefaultCurrency() string {
	// Try LC_ALL or LANG environment variables
	locale := os.Getenv("LC_ALL")
	if locale == "" {
		locale = os.Getenv("LANG")
	}
	if locale != "" {
		// Parse locale (e.g., "en_US.UTF-8" → "en_US")
		parts := strings.Split(locale, ".")
		if len(parts) > 0 {
			tag, err := language.Parse(parts[0])
			if err == nil {
				region, _ := tag.Region()
				if region.IsCountry() {
					unit, ok := currency.FromRegion(region)
					if ok {
						return unit.String()
					}
				}
			}
		}
	}
	return "USD" // Default to USD
}

// validateCurrency checks if a currency code is a valid ISO 4217 code.
func validateCurrency(code string) bool {
	if code == "" {
		return false
	}
	_, err := currency.ParseISO(code)
	return err == nil
}

// validateFloat checks if a string is a valid float.
func validateFloat(value string) bool {
	if value == "" {
		return false
	}
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}

// validateStruct validates currency and price fields in a struct, setting default currency if addressable.
func validateStruct(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			return nil // Skip validation for nil pointers
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("expected a struct or pointer to struct, got %v", val.Kind())
	}

	// Ensure the struct is addressable for modifications
	if !val.CanSet() {
		return fmt.Errorf("struct %v is not addressable; pass a pointer to modify currency fields", val.Type().Name())
	}

	typ := val.Type()
	defaultCurrency := getDefaultCurrency()

	switch typ.Name() {
	case "User":
		currencyField := val.FieldByName("Currency")
		if !currencyField.IsValid() {
			return fmt.Errorf("User.Currency field not found")
		}
		if currencyField.String() == "" && currencyField.CanSet() {
			currencyField.SetString(defaultCurrency)
		}
		if !validateCurrency(currencyField.String()) {
			return fmt.Errorf("invalid ISO 4217 currency code for User.Currency: %s", currencyField.String())
		}
	case "Datum":
		datumTypeField := val.FieldByName("DatumType")
		currencyField := val.FieldByName("Currency")
		valueField := val.FieldByName("Value")
		if !datumTypeField.IsValid() || !currencyField.IsValid() || !valueField.IsValid() {
			return fmt.Errorf("Datum fields (DatumType, Currency, Value) not found")
		}
		if datumTypeField.String() == "price" {
			if currencyField.IsNil() && currencyField.CanSet() {
				currencyField.Set(reflect.ValueOf(&defaultCurrency))
			} else if !currencyField.IsNil() && !validateCurrency(*currencyField.Interface().(*string)) {
				return fmt.Errorf("invalid ISO 4217 currency code for Datum.Currency: %s", *currencyField.Interface().(*string))
			}
			if valueField.IsNil() {
				return fmt.Errorf("Datum.Value must be set for price type")
			}
			if !validateFloat(*valueField.Interface().(*string)) {
				return fmt.Errorf("invalid float for Datum.Value (price): %s", *valueField.Interface().(*string))
			}
		}
	case "Wish":
		currencyField := val.FieldByName("Currency")
		priceField := val.FieldByName("Price")
		if !currencyField.IsValid() || !priceField.IsValid() {
			return fmt.Errorf("Wish fields (Currency, Price) not found")
		}
		if currencyField.IsNil() && priceField.IsValid() && !priceField.IsNil() && currencyField.CanSet() {
			currencyField.Set(reflect.ValueOf(&defaultCurrency))
		} else if !currencyField.IsNil() && !validateCurrency(*currencyField.Interface().(*string)) {
			return fmt.Errorf("invalid ISO 4217 currency code for Wish.Currency: %s", *currencyField.Interface().(*string))
		}
		if priceField.IsValid() && !priceField.IsNil() && !validateFloat(*priceField.Interface().(*string)) {
			return fmt.Errorf("invalid float in Wish.Price: %s", *priceField.Interface().(*string))
		}
	}

	return nil
}

// printStruct generically prints the fields of a struct using reflection, aligning values at the same left margin,
// skipping fields tagged with omitempty if their values would be omitted in JSON marshaling. For Datum in non-verbose
// mode, only DatumType, Label, and Value (if non-nil) are printed. Uses the specified indent level for field lines.
func printStruct(v interface{}, indentLevel int, verbose bool, format string, args ...interface{}) (int, error) {
	if v == nil {
		fmt.Println("<nil>")
		return 0, nil
	}

	// Validate currency and price fields before printing
	if err := validateStruct(v); err != nil {
		return 0, err
	}

	// Get the value and dereference if it's a pointer
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		if val.IsNil() {
			fmt.Println("<nil>")
			return 0, nil
		}
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return 0, fmt.Errorf("expected a struct or pointer to struct, got %v", val.Kind())
	}

	typ := val.Type()
	numFields := val.NumField()

	// Print the formatted header
	if format != "" {
		fmt.Printf(format, args...)
	} else {
		fmt.Printf("%s\n", typ.Name())
	}

	if numFields == 0 {
		fmt.Printf("%s%s: <empty>\n", strings.Repeat(" ", indentLevel), typ.Name())
		return 0, nil
	}

	// Determine fields to print
	var fieldsToPrint []int
	if !verbose && typ.Name() == "Datum" {
		for i := 0; i < numFields; i++ {
			name := typ.Field(i).Name
			if name == "DatumType" || name == "Label" || (name == "Value" && !val.Field(i).IsNil()) {
				fieldsToPrint = append(fieldsToPrint, i)
			}
		}
	} else {
		for i := 0; i < numFields; i++ {
			if jsonTag := typ.Field(i).Tag.Get("json"); strings.Contains(jsonTag, ",omitempty") {
				field := val.Field(i)
				switch field.Kind() {
				case reflect.Ptr:
					if field.IsNil() {
						continue
					}
				case reflect.Slice:
					if field.IsNil() || field.Len() == 0 {
						continue
					}
				case reflect.String:
					if field.String() == "" {
						continue
					}
				case reflect.Int:
					if field.Int() == 0 {
						continue
					}
				case reflect.Bool:
					if !field.Bool() {
						continue
					}
				case reflect.Struct:
					if field.Type() == reflect.TypeOf(time.Time{}) && field.Interface().(time.Time).IsZero() {
						continue
					}
				}
			}
			fieldsToPrint = append(fieldsToPrint, i)
		}
	}

	if len(fieldsToPrint) == 0 {
		fmt.Printf("%s%s: <empty>\n", strings.Repeat(" ", indentLevel), typ.Name())
		return 0, nil
	}

	// Calculate max field name length for alignment
	maxLen := 0
	for _, i := range fieldsToPrint {
		name := typ.Field(i).Name
		if len(name) > maxLen {
			maxLen = len(name)
		}
	}

	// Print fields with aligned values, indented by indentLevel spaces
	printedFields := 0
	for _, i := range fieldsToPrint {
		field := val.Field(i)
		fieldType := typ.Field(i)
		name := fieldType.Name
		//jsonTag := fieldType.Tag.Get("json")

		// Check for non-compliant currency fields from server
		if (typ.Name() == "User" && name == "Currency") ||
			((typ.Name() == "Datum" || typ.Name() == "Wish") && name == "Currency") {
			if field.Kind() == reflect.String && field.String() != "" && !validateCurrency(field.String()) {
				fmt.Printf("%sWARNING: Invalid currency code '%s' for %s.%s\n", strings.Repeat(" ", indentLevel), field.String(), typ.Name(), name)
			} else if field.Kind() == reflect.Ptr && !field.IsNil() && !validateCurrency(field.Elem().String()) {
				fmt.Printf("%sWARNING: Invalid currency code '%s' for %s.%s\n", strings.Repeat(" ", indentLevel), field.Elem().String(), typ.Name(), name)
			}
		}
		if (typ.Name() == "Datum" && name == "Value" && val.FieldByName("DatumType").String() == "price") ||
			(typ.Name() == "Wish" && name == "Price") {
			if field.Kind() == reflect.Ptr && !field.IsNil() && !validateFloat(field.Elem().String()) {
				fmt.Printf("%sWARNING: Invalid float value '%s' for %s.%s\n", strings.Repeat(" ", indentLevel), field.Elem().String(), typ.Name(), name)
			}
		}

		prefix := strings.Repeat(" ", indentLevel) + typ.Name() + "." + name
		padding := strings.Repeat(" ", maxLen-len(name)+1)

		switch field.Kind() {
		case reflect.Ptr:
			if field.IsNil() {
				fmt.Printf("%s:%s<nil>\n", prefix, padding)
			} else {
				switch field.Type().Elem().Kind() {
				case reflect.String:
					fmt.Printf("%s:%s%s\n", prefix, padding, field.Elem().String())
				case reflect.Int:
					fmt.Printf("%s:%s%d\n", prefix, padding, field.Elem().Int())
				case reflect.Bool:
					fmt.Printf("%s:%s%t\n", prefix, padding, field.Elem().Bool())
				case reflect.Struct:
					if field.Type().Elem() == reflect.TypeOf(time.Time{}) {
						fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().Interface())
					} else {
						fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().String())
					}
				default:
					fmt.Printf("%s:%s%v\n", prefix, padding, field.Elem().Interface())
				}
			}
		case reflect.Slice:
			if field.IsNil() {
				fmt.Printf("%s:%s[]\n", prefix, padding)
			} else {
				fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
			}
		default:
			switch field.Kind() {
			case reflect.String:
				fmt.Printf("%s:%s%s\n", prefix, padding, field.String())
			case reflect.Int:
				fmt.Printf("%s:%s%d\n", prefix, padding, field.Int())
			case reflect.Bool:
				fmt.Printf("%s:%s%t\n", prefix, padding, field.Bool())
			case reflect.Struct:
				if field.Type() == reflect.TypeOf(time.Time{}) {
					fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
				} else {
					fmt.Printf("%s:%s%v\n", prefix, padding, field.String())
				}
			default:
				fmt.Printf("%s:%s%v\n", prefix, padding, field.Interface())
			}
		}
		printedFields++
	}

	return printedFields, nil
}

// GetItemAndData retrieves an Item and all associated Datum objects using the Client.
func GetItemAndData(client Client, itemID ID) (*Item, []*Datum, error) {
	// Fetch the Item
	item, err := client.GetItem(itemID)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get item %s: %w", itemID, err)
	}
	// Fetch all Datum objects associated with the item
	data, err := client.ListItemData(itemID)
	return item, data, nil
}

// PrintItemWithData prints the fields of an Item and all associated Datum objects, indenting Datum fields further.
// Datum items are sorted by Position (ascending, with nil/negative at end). In non-verbose mode, only DatumType,
// Label, and Value (if non-nil) are printed for Datum.
func PrintItemWithData(item *Item, data []*Datum, verbose bool, format string, args ...interface{}) (int, error) {
	totalFields := 0

	// Print the formatted header
	if format != "" {
		fmt.Printf(format, args...)
	} else {
		fmt.Println("Item")
	}

	// Print Item fields with 4-space indent
	if item == nil {
		fmt.Println("    Item: <nil>")
	} else {
		itemFields, err := printStruct(item, 4, true, "")
		if err != nil {
			return 0, fmt.Errorf("failed to print item: %w", err)
		}
		totalFields += itemFields
	}

	// Sort Datum by Position (nil/negative to end)
	sortedData := make([]*Datum, len(data))
	copy(sortedData, data)
	sort.SliceStable(sortedData, func(i, j int) bool {
		posI, posJ := sortedData[i].Position, sortedData[j].Position
		if posI == nil || *posI < 0 {
			return false // i is invalid, sort to end
		}
		if posJ == nil || *posJ < 0 {
			return true // j is invalid, sort to end
		}
		return *posI < *posJ
	})

	// Print each Datum with 8-space indent
	for i, datum := range sortedData {
		if datum == nil {
			fmt.Printf("        Datum[%d]: <nil>\n", i)
			continue
		}
		datumFields, err := printStruct(datum, 8, verbose, "        Datum[%d]\n", i)
		if err != nil {
			return totalFields, fmt.Errorf("failed to print datum %d: %w", i, err)
		}
		totalFields += datumFields
	}

	return totalFields, nil
}

// Print methods for each struct type
func (a Album) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&a, 4, true, format, args...)
}

func (c ChoiceList) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&c, 4, true, format, args...)
}

func (c Collection) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&c, 4, true, format, args...)
}

func (d Datum) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&d, 4, true, format, args...)
}

func (f Field) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&f, 4, true, format, args...)
}

func (i Inventory) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&i, 4, true, format, args...)
}

func (i Item) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&i, 4, true, format, args...)
}

func (l Loan) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&l, 4, true, format, args...)
}

func (l Log) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&l, 4, true, format, args...)
}

func (p Photo) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&p, 4, true, format, args...)
}

func (t Tag) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&t, 4, true, format, args...)
}

func (tc TagCategory) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&tc, 4, true, format, args...)
}

func (t Template) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&t, 4, true, format, args...)
}

func (u User) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&u, 4, true, format, args...)
}

func (w Wish) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&w, 4, true, format, args...)
}

func (w Wishlist) Print(format string, args ...interface{}) (int, error) {
	return printStruct(&w, 4, true, format, args...)
}
