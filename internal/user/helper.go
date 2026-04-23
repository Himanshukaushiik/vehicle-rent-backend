package user

import "strconv"

func parseID(id string) uint {
	parsed, _ := strconv.Atoi(id)
	return uint(parsed)
}
