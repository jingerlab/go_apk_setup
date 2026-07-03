package slidevapp

import (
	"fmt"
	"net/http"
	"time"
)

const (
	slidevURL = "https://jingerlab.github.io/2026_daup_lecture_2/"
	localPort = "127.0.0.1:9999"
)

const errorHTML = `
<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style>
        body { font-family: system-ui, sans-serif; text-align: center; padding: 40px; background-color: #0f172a; color: #f8fafc; }
        h1 { color: #f43f5e; font-size: 24px; }
        p { color: #94a3b8; font-size: 16px; }
    </style>
</head>
<body>
    <h1>Presentation Offline</h1>
    <p>Unable to reach the lecture slides. Please check your internet connection.</p>
</body>
</html>`

// InitializeAndroidApp runs the network check and local server loop, returning the proper view URL
func InitializeAndroidApp() string {
	// 1. Boot up the local offline loopback server fallback execution path
	go func() {
		http.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			fmt.Fprint(w, errorHTML)
		})
		_ = http.ListenAndServe(localPort, nil)
	}()

	// 2. Health check the target Slidev presentation deployment endpoint
	target := slidevURL
	client := http.Client{Timeout: 3 * time.Second}
	if _, err := client.Get(slidevURL); err != nil {
		target = "http://" + localPort + "/404"
	}

	return target
}
