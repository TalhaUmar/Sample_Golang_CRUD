package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	// fmt.Println(printEmoji([]int{1, 3, 4}, 0, ""))

	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func printEmoji(nums []int, currIndex int, emoji string) string {
	emap := map[int]string{
		0: ":zero:",
		1: ":one:",
		2: ":two:",
		3: ":three:",
		4: ":four:",
		5: ":five:",
		6: ":six:",
		7: ":seven:",
		8: ":eight:",
		9: ":nine:",
	}

	if len(nums) == currIndex {
		return emoji
	}

	return printEmoji(nums, currIndex+1, emoji+emap[nums[currIndex]])
}

func foo(w http.ResponseWriter, req *http.Request) {

	var s string

	if req.Method == http.MethodPost {
		f, _, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer f.Close()

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)

	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q" />
	<input type="submit" />
	</form>
	<br>`+s)
}
