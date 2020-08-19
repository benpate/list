package list

import "strings"

// Head returns the FIRST item in a string-based-list
func Head(value string, delimiter string) string {

	index := strings.Index(value, delimiter)

	if index == -1 {
		return value
	}

	return value[:index]
}

// Tail returns any values in the string-based-list AFTER the first item
func Tail(value string, delimiter string) string {
	index := strings.Index(value, delimiter)

	if index == -1 {
		return ""
	}

	return value[index+1:]
}

// RemoveLast returns the full list, with the last element removed.
func RemoveLast(value string, delimiter string) string {

	index := strings.LastIndex(value, delimiter)

	if index == -1 {
		return ""
	}

	return value[:index]
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
		return Head(value, delimiter)
	}

	tail := Tail(value, delimiter)

	if tail == "" {
		return ""
	}

	return At(tail, delimiter, index-1)
}
