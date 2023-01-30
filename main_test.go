package main

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {
	// Leer el archivo HTML
	html, err := ioutil.ReadFile("public/index.html")
	if err != nil {
		t.Fatal(err)
	}

	// Utilizar goquery para analizar el HTML
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	// Extraer el título del body
	title := doc.Find("h1").Text()

	// Comprobar que el título del body cumple la expresión regular
	match, _ := regexp.MatchString("(?i)^Hello World from.*", title)
	assert.True(t, match)
}
