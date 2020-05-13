package list

import "strings"

// First returns the FIRST item in a string-based-list
func First(value string, delimiter string) string {

	index := strings.Index(value, delimiter)

	if index == -1 {
		return value
	}

	return value[:index]
}

// Rest returns any values in the string-based-list AFTER the first item
func Rest(value string, delimiter string) string {
	index := strings.Index(value, delimiter)

	if index == -1 {
		return ""
	}

	return value[index+1:]
}

// Last returns the LAST item in a string-based-list
func Last(value string, delimiter string) string {

	index := strings.LastIndex(value, delimiter)

	if index == -1 {
		return value
	}

	return value[index+1:]
}

// Split returns the FIRST element, and the REST element in one function call
func Split(value string, delimiter string) (string, string) {

	index := strings.Index(value, delimiter)

	if index == -1 {
		return value, ""
	}

	return value[:index], value[index+1:]

}

// At returns the list vaue at a particular index
func At(value string, delimiter string, index int) string {

	if index == 0 {
		return First(value, delimiter)
	}

	rest := Rest(value, delimiter)

	if rest == "" {
		return ""
	}

	return At(rest, delimiter, index-1)
}
