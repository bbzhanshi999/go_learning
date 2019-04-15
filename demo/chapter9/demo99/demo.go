package main

import (
	"fmt"
	"net/http"
	"text/template"

	"google.golang.org/api/urlshortener/v1"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)

	http.ListenAndServe("localhost:8080", nil)
}

func root(rw http.ResponseWriter, r *http.Request) {
	rootHtmlTmpl.Execute(rw, nil)
}

func short(rw http.ResponseWriter, r *http.Request) {
	longUrl := r.FormValue("longUrl")
	urlShorterSvc, _ := urlshortener.New(http.DefaultClient)
	url, _ := urlShorterSvc.Url.Insert(&urlshortener.Url{LongUrl: longUrl}).Do()
	rootHtmlTmpl.Execute(rw, fmt.Sprintf("shortened version of %s is:%s", longUrl, url.Id))
}

func long(rw http.ResponseWriter, r *http.Request) {
	shortUrl := "http://goo.gl/" + r.FormValue("shortUrl")
	urlShortenerSvc, _ := urlshortener.New(http.DefaultClient)
	url, err := urlShortenerSvc.Url.Get(shortUrl).Do()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	rootHtmlTmpl.Execute(rw, fmt.Sprintf("Long version of %s is : %s", shortUrl, url.LongUrl))
}

// the template used to show the forms andx the results web page to the user
var rootHtmlTmpl = template.Must(template.New("rootHtml").Parse(`
	<html><body>
	<h1>URL SHORTENER</h1>
	{{if .}}{{.}}<br /><br />{{end}}
	<form action="/short" type="POST">
	Shorten this: <input type="text" name="longUrl" />
	<input type="submit" value="Give me the short URL" />
	</form>
	<br />
	<form action="/long" type="POST">
	Expand this: http://goo.gl/<input type="text" name="shortUrl" />
	<input type="submit" value="Give me the long URL" />
	</form>
	</body></html>
`))
