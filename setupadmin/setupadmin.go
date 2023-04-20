package main

import (
	f "forum"
	"os"
)

func main() {
	list_username := os.Args[1:]
	f.Init()
	for _, username := range list_username {
		f.RankupAdmin(username)
	}
}
