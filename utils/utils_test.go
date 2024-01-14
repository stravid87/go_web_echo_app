package utils

import "testing"

// func BuildResponse(status bool, message string, data interface{}) Response {
// 	res := Response{
// 		Status:  status,
// 		Message: message,
// 		Error:   nil,
// 		Data:    data,
// 	}
// 	return res
// }

func TestBuildResponse(t *testing.T) {
	// t.Log()
	// t.Fail()
	// t.FailNow()
	// t.Error()
	// t.Fatal()
	res := BuildResponse(true, "404", map[string]string{"name": "bro", "data": "you failed?"})
	if res.Status != true {
		t.Fatal("Status False and should be True")
	}
	if res.Message != "404" {
		t.Error("Message was not '404'. Received this instead: ", res.Message)
	}
	t.Log("res.Data", res.Data)

}
