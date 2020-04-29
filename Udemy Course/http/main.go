package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// in order to see the body of the response we use an io reader , in this case a reader interface
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// fmt.Println(resp)

	//creates a slice of type byte with 99999 available spaces
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// this line replaces the above 3 lines condensing the code and simplifying it
	lw := logWriter{}

	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Bytes Written: ", len(bs))
	return len(bs), nil
}
