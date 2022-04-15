package helpers

import "strings"

func ArnSplit(arn string) []string {
	return strings.Split(arn, ",")
}
