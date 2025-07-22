package utils

import "log"

func SendEmail(to, subject, body string) error {
	log.Printf("[EMAIL] To: %s | Subject: %s | Body: %s", to, subject, body)
	return nil
}
