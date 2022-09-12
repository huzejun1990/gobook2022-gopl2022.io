// Package storage is part of a hypothetical cloud storage server.
// 包存储是假设的云存储服务器的一部分。
//!+main
package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 {
	return usage[username]
}

// Email sender configuration.
// 电子邮件发件人配置。
// NOTE: never put passwords in source code!
// 注意：永远不要在源代码中输入密码！
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"

const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // ok
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}

//!-main
