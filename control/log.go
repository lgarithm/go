package control

import "log"

// LogError is an identity function on error with log side effect
func LogError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
