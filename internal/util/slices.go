package util

import "slices"

func AppendUniqueString(slice []string, value string) []string {
	if slices.Contains[[]string, string](slice, value) {
		return slice
	}

	return append(slice, value)
}

func DeleteString(slice []string, value string) []string {
	return slices.DeleteFunc(slice, func(s string) bool {
		return s == value
	})
}
