package main

import (
	"net/http"
	"os"
)

func listSunDir(dirname string) []string {
	itemDir := make([]string, 0)
	entries, err := os.ReadDir(dirname)
	if err != nil {
		return itemDir
	}
	//infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return itemDir
		}
		if info.IsDir() {
			itemDir = append(itemDir, info.Name())
		}
	}
	return itemDir
}

func main() {
	http.ListenAndServe(":9808", http.FileServer(http.Dir("./app/")))
}
