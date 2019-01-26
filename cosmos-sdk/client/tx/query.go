package tx

import (
	"io/ioutil"
	"net/http"
)

func TxQuery(hash string) string {

	res, err := http.Get("http://localhost:26657/tx?hash=0x" + hash)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return string(data)
}
