package helpers

import "strings"

func GetToolIdFromKey(key string) string {
	splitted := strings.Split(key,"/")
	return splitted[len(splitted) - 1]
}