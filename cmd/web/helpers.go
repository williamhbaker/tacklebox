package main

import "encoding/json"

func validJSONBytes(b []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(b, &js) == nil
}
