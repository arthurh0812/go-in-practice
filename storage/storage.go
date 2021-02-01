package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]uint64)

// returns the bytes that are used by the user with the given username
func bytesInUsage(username string) uint64 {
	return usage[username]
}

// Email sender configuration
const sender = "notifications@example.com"
const password = "securepassword1234"
const hostname = "smtp.example.com"

var template = `Warning: you are using %d bytes of storage, %d%% of your quota`

// CheckQuota uses the given username to check if he has overstepped the storage quota
func CheckQuota(username string) {
	used := usage[username]
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":smtp", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Fatalf("smtp.SendMail(%q) failed: %s", username, err)
	}
}
