package main

import (
	"testing"
)

func TestreplaceAntiSyakaiseiWords(t *testing.T) {
	assertSyakai := func(baseText, expectedText string) {
		text := replaceAntiSyakaiseiWords(baseText)
		if text != expectedText {
			t.Errorf("%s should be %s", text, expectedText)
		}
	}

	assertSyakai("うるせえ黙れ", "承知いたしました")
	assertSyakai("うるせえだまれ", "承知いたしました")
}
