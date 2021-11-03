package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// IfThenElse evaluates a condition, if true returns the first parameter otherwise the second
func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// UpdateIfNew returns the updated value or the same value if not changed
func UpdateIfNew(new interface{}, old interface{}) interface{} {
	return IfThenElse(new==old, old, new)
}

// MakeSyncPatchCall makes a patch call to the specified url to keep gorm models in sync
func MakeSyncPatchCall(url string, body interface{}, method string) ([]byte, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewBuffer(b)

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	res, _ := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}