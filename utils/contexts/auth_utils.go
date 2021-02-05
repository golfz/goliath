package contexts

import "strings"

func isBearerAuth(s string) bool {
	bearerIndex := strings.Index(s, bearerStartPattern)
	return bearerIndex == 0
}
