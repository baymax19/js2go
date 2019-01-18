package main

import (
	"github.com/baymax19/js2go/cosmos-sdk/client/keys"
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Module.Get("exports").Set("createKey", keys.CreateKey)
}
