package handler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/spiegel-im-spiegel/text/decode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)

	reader := bytes.NewReader(body)

	res, err := decode.ToUTF8ja(reader)

	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(w, res)
}
