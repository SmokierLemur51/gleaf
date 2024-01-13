/*
File: errors.go

Trying to create a map of common errors to to be used in the application,
*/
package data

var ERROR_MAP map[int]string = map[int]string{
	1: "Doesn't exist",
	2: "Already exists",
	3: "Missing Requirements",
}

func CheckErr(e error) string {
	if e != nil {
		panic(e)
	}
	return ""
}
