package main

import "testing"

func TestProcess(t *testing.T) {
	cases := []struct{ in, want string }{
		// cas de base
		{"1E (hex) files", "30 files"},
		{"BAMM !!", "BAMM!!"},
		{"a hour", "an hour"},
		// cas de référence de l'audit (analyse § 7.1)
		{"If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?",
			"If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"},
		{"I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure",
			"I have to pack 5 outfits. Packed 26 just to be sure"},
		{"Don not be sad ,because sad backwards is das . And das not good",
			"Don not be sad, because sad backwards is das. And das not good"},
		{"harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '",
			"Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"},
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68."},
		{"Punctuation tests are ... kinda boring ,what do you think ?",
			"Punctuation tests are... kinda boring, what do you think?"},
		// exemples de l'analyse (§ 7.2)
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		// cas limites (analyse § 6)
		{"(up) hello", "hello"},
		{"hello world (up, 5)", "HELLO WORLD"},
		{"hello (low, 0)", "hello"},
		{"go (up) (low)", "go"},
		{"a, b (cap, 2)", "A, B"},
		{"élan (cap)", "Élan"},
		{"zz (hex)", "zz"},
		{"don't stop", "don't stop"},
		{"hello (up)\nworld (cap)", "HELLO\nWorld"},
		{"", ""},
	}
	for _, c := range cases {
		got := Process(c.in)
		if got != c.want {
			t.Errorf("Process(%q) = %q, want %q",
				c.in, got, c.want)
		}
	}
}
