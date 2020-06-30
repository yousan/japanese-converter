package handler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spiegel-im-spiegel/text/decode"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// Conversion
func Conversion(inStream io.Reader, outStream io.Writer) error {
	//reader from stream (Shift-JIS to UTF-8)
	reader := transform.NewReader(inStream, japanese.ShiftJIS.NewDecoder())

	_, err := io.Copy(outStream, reader)
	return err
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	//	decoded := json.NewDecoder(r.Body)

	// Do something with the Person struct...
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "Person: %s", body)

	reader := bytes.NewReader(body)

	//fmt.Println("hi")

	res, err := decode.ToUTF8ja(reader)

	if err != nil {
		fmt.Println(err)
		return
	}
	//buf := new(bytes.Buffer)
	io.Copy(w, res)
	// fmt.Println(buf.Bytes())

	//fmt.Fprintf(w, "%s", res)
}
