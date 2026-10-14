package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var markerRe = regexp.MustCompile(`\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)`)
var punctBeforeRe = regexp.MustCompile(`\s+([.,!?:;]+)`)
var punctAfterRe = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)

func Process(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		line = markerRe.ReplaceAllStringFunc(line, normalizeMarker)
		words := strings.Fields(line)
		var out []string
		for _, word := range words {
			m := markerRe.FindStringSubmatch(word)
			if m == nil {
				out = append(out, word)
				continue
			}
			count := 1
			if m[2] != "" {
				count, _ = strconv.Atoi(m[2])
			}
			applyMarker(out, m[1], count)
		}
		lines[i] = fixArticles(fixQuotes(fixPunctuation(strings.Join(out, " "))))
	}
	return strings.Join(lines, "\n")
}

func capitalize(word string) string {
	r := []rune(strings.ToLower(word))
	if len(r) == 0 {
		return word
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func convertBase(word string, base int) string {
	n, err := strconv.ParseInt(word, base, 64)
	if err != nil {
		return word
	}
	return strconv.FormatInt(n, 10)
}

func normalizeMarker(marker string) string {
	marker = strings.ReplaceAll(marker, " ", "")
	return " " + marker + " "
}

func applyMarker(out []string, cmd string, count int) {
	start := len(out) - count
	if start < 0 {
		start = 0
	}
	for i := start; i < len(out); i++ {
		switch cmd {
		case "up":
			out[i] = strings.ToUpper(out[i])
		case "low":
			out[i] = strings.ToLower(out[i])
		case "cap":
			out[i] = capitalize(out[i])
		case "hex":
			out[i] = convertBase(out[i], 16)
		case "bin":
			out[i] = convertBase(out[i], 2)
		}
	}
}

func fixPunctuation(line string) string {
	line = punctBeforeRe.ReplaceAllString(line, "$1")
	line = punctAfterRe.ReplaceAllString(line, "$1 $2")
	return line
}

func fixQuotes(line string) string {
	words := strings.Fields(line)
	var out []string
	open := false
	glueNext := false
	for _, w := range words {
		if w == "'" {
			if !open {
				open = true
				glueNext = true
			} else {
				open = false
				if len(out) > 0 {
					out[len(out)-1] += "'"
				} else {
					out = append(out, "'")
				}
			}
			continue
		}
		if glueNext {
			w = "'" + w
			glueNext = false
		}
		out = append(out, w)
	}
	if glueNext {
		out = append(out, "'")
	}
	return strings.Join(out, " ")
}

func fixArticles(line string) string {
	words := strings.Fields(line)
	for i := 0; i < len(words)-1; i++ {
		if words[i] != "a" && words[i] != "A" {
			continue
		}
		next := strings.TrimLeft(words[i+1], "'")
		if next != "" && strings.ContainsRune("aeiouhAEIOUH", rune(next[0])) {
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}
