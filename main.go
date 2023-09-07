package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		html := `
    <!DOCTYPE html>
    <html>
    <head>
      <title>HTMX Go Example</title>
      <script src="https://unpkg.com/htmx.org@1.9.5/dist/htmx.min.js"></script>
    </head>
    <body>
      <h1>Go HTMX QR Code generator</h1>
      <textarea
        rows=5
        cols=45
        hx-post="/generate"
        hx-trigger="keyup changed delay:500ms"
        hx-target="#output" name="qrcontent" id="qrcontent"></textarea>
      <div id="output" />
    </body>
    </html>
    `

		fmt.Fprint(w, html)
	})

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		content := r.FormValue("qrcontent")

		log.Printf("Generating code for %s", content)

		png, err := qrcode.Encode(content, qrcode.Medium, 512)
		if err != nil {
			log.Fatal(err)
		}

		base64img := base64.RawStdEncoding.EncodeToString(png)
		log.Printf("QR code base64: %s", base64img)

		image := fmt.Sprintf("<img src=\"data:image/png;base64,%s\">", base64img)

		fmt.Fprint(w, image)
	})

	http.ListenAndServe(":8080", nil)
}
