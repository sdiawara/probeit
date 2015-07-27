package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMain(testing *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/", strings.NewReader("z=post&both=y&prio=2&empty="))

	HelloHandler(writer, request)

	assert.Equal(testing, true, strings.Contains(writer.Body.String(), "<img alt=\"logo\" src=\"/images/logo.svg\" id=\"logo\" width=\"150px\" />"))
	assert.Equal(testing, true, strings.Contains(writer.Body.String(), "<h1 class=\"cover-heading\">Nous les sondons pour vous.</h1>"))
}
