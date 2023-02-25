package validators

import "regexp"

func IsEmailValid(email string) bool {
	var reEmail = regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

	if len(email) < 3 || len(email) > 254 || !reEmail.MatchString(email) {
		return false
	}
	return true
}
