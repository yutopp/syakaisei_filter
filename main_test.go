package main

import (
	"testing"
)

func TestReplaceAntiSyakaiseiWords(t *testing.T) {
	assertSyakai := func(baseText, expectedText string) {
		text := replaceAntiSyakaiseiWords(baseText)
		if text != expectedText {
			t.Errorf("%s should be %s", text, expectedText)
		}
	}

	assertSyakai("うるせえ黙れ", "承知いたしました")
	assertSyakai("うるせえだまれ", "承知いたしました")
	assertSyakai("モチベーションが完全に無くなった", "テンションMAX！")
	assertSyakai("モチベーション完全に無くなった", "テンションMAX！")
	assertSyakai("モチベーションが無くなった", "テンションMAX！")
	assertSyakai("モチベーションが完全に消えた", "テンションMAX！")
	assertSyakai("モチベーション消えた", "テンションMAX！")		
}
