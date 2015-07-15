package main

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestMain(testing *testing.T) {
	writer := httptest.NewRecorder()

	HelloHandler(writer, nil)

	assert.Equal(testing, "Hello World!", writer.Body.String())
}
