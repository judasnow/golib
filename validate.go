package golib

import "regexp"

const EMAIL_REGEXP = `^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`

func IsValidEmail(email string) (bool, error) {
	if email == "" {
		return false, nil
	} else {
		if match, err := regexp.MatchString(EMAIL_REGEXP, email); err != nil {
			return false, err
		} else {
			return match, nil
		}
	}
}
