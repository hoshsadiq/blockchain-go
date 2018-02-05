package server

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

func requestBodyToPayload(writer http.ResponseWriter, request *http.Request, payload interface{}) bool {
    if request.Body == nil {
        http.Error(writer, "Please send a request body", 400)
        return false
    }
    body, _ := ioutil.ReadAll(request.Body)
    if string(body) == "" {
        http.Error(writer, "Please send a request body", 400)
        return false
    }

    err := json.Unmarshal(body, &payload)
    if err != nil {
        http.Error(writer, err.Error(), http.StatusBadRequest)
        return false
    }

    return true
}
