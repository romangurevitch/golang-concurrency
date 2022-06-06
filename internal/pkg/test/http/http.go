package http

import (
	"encoding/json"
	jsonpkg "github.com/romangurevitch/golang-concurrency/internal/pkg/json"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CheckAPIResponse(t *testing.T, actualBody io.Reader, actualCode, expectedCode int, expectedObj interface{}) {
	assert.Equal(t, expectedCode, actualCode)

	got, err := ioutil.ReadAll(actualBody)
	assert.NoError(t, err)
	if len(got) > 0 {
		expected, err := json.Marshal(expectedObj)
		assert.NoError(t, err)
		assert.JSONEq(t, string(expected), string(got))
		return
	}

	assert.Equal(t, expectedObj, got)
}

func ToJSONPayloadReader(obj interface{}) io.Reader {
	if obj == nil {
		return nil
	}

	return strings.NewReader(jsonpkg.ToJSONString(obj))
}
