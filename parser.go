package ehsan

import "strings"

func getURI(bs []byte) string {
	// convert by to string
	fullBody := string(bs)
	fullBody = strings.Replace(fullBody, "\r\n", "\n", -1) // windows are \r\n
	fullBody = strings.Replace(fullBody, "\r", "\n", -1)   // mac/osx are \r

	// new band name
	sliceOfBody := strings.Split(fullBody, "\n") // use \n because i use linex

	h1 := sliceOfBody[0]
	h1 = strings.Replace(h1, "GET ", "", 1)      // do not follow
	h1 = strings.Replace(h1, " HTTP/1.1", "", 1) // do not follow

	// return is to return something
	return h1
}
