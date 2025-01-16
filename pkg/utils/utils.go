package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil { //reading the body of the request
		if err := json.Unmarshal([]byte(body), x); err != nil { //unmarshalling the body of the request
			return

		}
	}
}
