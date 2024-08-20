package models

//RequestData for request input 
type RequestData struct {
	Numbers []int `json:"numbers"`
	Target int  `json:"target"`
}

//ResponseData for respose output
type ResponseData struct {
	Solutions [][]int `json:"solutions"`
}