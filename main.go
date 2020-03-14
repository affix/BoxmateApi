package main

import (
	"os"

	boxmate "github.com/Affix/BoxmateApi/pkg"
)

func main() {
	token, _ := boxmate.BoxmateLogin(os.Getenv("BOXMATE_EMAIL"), os.Getenv("BOXMATE_PASSWORD"))
	println(token.ApiKey)
}
