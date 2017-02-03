package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		os.Exit(1)
	}

	fmt.Print(replaceAntiSyakaiseiWords(string(body)))
}

var replaceTable = map[*regexp.Regexp]string{
	regexp.MustCompile("うるせえ(黙れ|だまれ)"): "承知いたしました",
}

func replaceAntiSyakaiseiWords(s string) string {
	for k, v := range replaceTable {
		s = k.ReplaceAllString(s, v)
	}
	return s
}
