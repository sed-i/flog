package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	// App log
	// 2021-06-10 17:25:01 AEST | CORE | TRACE | (pkg/collector/scheduler/job.go:196 in process) | Jobs in bucket: []
	AppLog = "%s | %s | %s | (%s.go:%d) | "
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	ApacheCommonLog = "%s - %s [%s] \"%s %s %s\" %d %d"
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes} "{referrer}" "{agent}"
	ApacheCombinedLog = "%s - %s [%s] \"%s %s %s\" %d %d \"%s\" \"%s\""
	// ApacheErrorLog : [{timestamp}] [{module}:{severity}] [pid {pid}:tid {thread-id}] [client %{client}:{port}] %{message}
	// Message is appened
	ApacheErrorLog = "[%s] [%s:%s] [pid %d:tid %d] [client %s:%d] "
	// RFC3164Log : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	// Message is appened
	RFC3164Log = "<%d>%s %s %s[%d]: "
	// RFC5424Log : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	// Message is appended
	RFC5424Log = "<%d>%d %s %s %s %d ID%d %s "
	// CommonLogFormat : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	CommonLogFormat = "%s - %s [%s] \"%s %s %s\" %d %d"
	// JSONLogFormat : {"host": "{host}", "user-identifier": "{user-identifier}", "datetime": "{datetime}", "method": "{method}", "request": "{request}", "protocol": "{protocol}", "status", {status}, "bytes": {bytes}, "referer": "{referer}", "message": "
	// message is appened
	JSONLogFormat = `{"host":"%s", "user-identifier":"%s", "datetime":"%s", "method": "%s", "request": "%s", "protocol":"%s", "status":%d, "bytes":%d, "referer": "%s", "message": "`
)

var cacheSize = 10000
var cache = []string{}

func buildCache(length int) {
	for i := 0; i < cacheSize; i++ {
		msg := gofakeit.Word()
		for len(msg) <= length {
			msg = msg + " " + gofakeit.Word()
		}
		cache = append(cache, msg[:length-1])
	}
}

func message(length int) string {
	if length < 1 {
		return ""
	}

	msg := gofakeit.Word()
	for len(msg) <= length {
		msg = msg + " " + gofakeit.Word()
	}
	return msg[:length-1]
}

func NewAppLog(t time.Time, length int) string {
	preMsg := fmt.Sprintf(
		AppLog,
		t.Format(RFC3164),
		strings.ToUpper(gofakeit.HackerAbbreviation()),
		strings.ToUpper(gofakeit.LogLevel("general")),
		RandResourceURI(),
		gofakeit.Number(1, 999),
	)
	msg := cache[rand.Intn(cacheSize)]
	return preMsg + msg[:length-len(preMsg)]
}

// NewApacheCommonLog creates a log string with apache common log format
func NewApacheCommonLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCommonLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewApacheCombinedLog creates a log string with apache combined log format
func NewApacheCombinedLog(t time.Time) string {
	return fmt.Sprintf(
		ApacheCombinedLog,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(Apache),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(30, 100000),
		gofakeit.URL(),
		gofakeit.UserAgent(),
	)
}

// NewApacheErrorLog creates a log string with apache error log format
func NewApacheErrorLog(t time.Time, length int) string {
	preMsg := fmt.Sprintf(
		ApacheErrorLog,
		t.Format(ApacheError),
		gofakeit.Word(),
		gofakeit.LogLevel("apache"),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 10000),
		gofakeit.IPv4Address(),
		gofakeit.Number(1, 65535),
	)
	msg := cache[rand.Intn(cacheSize)]
	return preMsg + msg[:length-len(preMsg)]
}

// NewRFC3164Log creates a log string with syslog (RFC3164) format
func NewRFC3164Log(t time.Time, length int) string {
	preMsg := fmt.Sprintf(
		RFC3164Log,
		gofakeit.Number(0, 191),
		t.Format(RFC3164),
		strings.ToLower(gofakeit.Username()),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
	)
	msg := cache[rand.Intn(cacheSize)]
	return preMsg + msg[:length-len(preMsg)]
}

// NewRFC5424Log creates a log string with syslog (RFC5424) format
func NewRFC5424Log(t time.Time, length int) string {
	preMsg := fmt.Sprintf(
		RFC5424Log,
		gofakeit.Number(0, 191),
		gofakeit.Number(1, 3),
		t.Format(RFC5424),
		gofakeit.DomainName(),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 1000),
		"-", // TODO: structured data
	)
	msg := cache[rand.Intn(cacheSize)]
	return preMsg + msg[:length-len(preMsg)]
}

// NewCommonLogFormat creates a log string with common log format
func NewCommonLogFormat(t time.Time) string {
	return fmt.Sprintf(
		CommonLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewJSONLogFormat creates a log string with json log format
func NewJSONLogFormat(t time.Time, length int) string {
	preMsg := fmt.Sprintf(
		JSONLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
		gofakeit.URL(),
	)
	return preMsg + message(length-len(preMsg)-2) + `"}`
}
