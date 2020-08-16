package util

import "os"

// Palindrome ... check if msg is a palindrome
func Palindrome(msg string) (palindrome bool) {
	var revMsg string
	for _, v := range msg {
		revMsg = string(v) + revMsg
	}
	if revMsg == msg {
		palindrome = true
	}
	return
}

// SetDevEnv ... set local development environment
func SetDevEnv() {
	/*
		TODO:
		1) Add these settings to a config file
		2) Read from config file and set env variables
		3) Omit repeatative code
	*/
	os.Setenv("SVC_HOST", "0.0.0.0")
	os.Setenv("SVC_PORT", "8080")
	os.Setenv("SVC_VERSION", "/v1")
	os.Setenv("SVC_PATH_PREFIX", "messages")
	os.Setenv("METRICS_PORT", "6060")
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_HOST", "qlikdb")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "docker")
	os.Setenv("DB_PASSWD", "docker")
	os.Setenv("DB_NAME", "messageDB")
}
