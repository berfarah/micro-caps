package main

import (
	"net/http"
	"strings"
)

var alphabets = map[string]string{
	"uppercase": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"lowercase": "abcdefghijklmnopqrstuvwxyz",
	"smallcaps": "ᴀʙᴄᴅᴇꜰɢʜɪᴊᴋʟᴍɴᴏᴘǫʀsᴛᴜᴠᴡxʏᴢ",
	"spacedcase": "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ",
}

func getIndex(r rune) int {
	if index := strings.IndexRune(alphabets["lowercase"], r); index > -1 {
		return index;
	}
	return strings.IndexRune(alphabets["uppercase"], r)
}

func convert(reqType, input string) string {
	var output []rune;

	alphabet, ok := alphabets[reqType]
	alph := []rune(alphabet)
	if !ok {
		return input
	}

	for _, char := range input {
		if index := getIndex(char); index > -1 {
			char = alph[index]
		}
		output = append(output, char)
	}

	return string(output)
}

func handler(w http.ResponseWriter, r *http.Request) {
	reqType := r.URL.Query().Get("type")
	input := r.URL.Query().Get("text")
	output := convert(reqType, input)
	w.Write([]byte(output))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
