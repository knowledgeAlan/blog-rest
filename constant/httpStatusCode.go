package constant


const(
	BlogSuccess = 200
	BlogFailure = 400
)


var statusText = map[int]string{
	BlogSuccess: "success" ,
	BlogFailure : "blog failure",
	
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	return statusText[code]
}