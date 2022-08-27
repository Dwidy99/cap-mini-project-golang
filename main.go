package main

import "funding-api/app"

func main() {
	app.Start()
}

// ambil nilai header Authorization: Bearer tokenAsli
// dari header Autorization, kita ambil nilai tokennya saja
// kita validasi token
// kita ambil user_id
// ambil user dari db berdasarkan user_id lewat service
// kita set context isinya user