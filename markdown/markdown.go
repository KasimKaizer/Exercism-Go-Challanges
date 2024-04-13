package markdown

// implementation to refactor

import (
	"fmt"
	"strings"
)

// Changes to implement
// - split up the code into more manageable function
// - remove instances of string concatenation, with a string builder or Sprintf
// - improve the nesting, the code seems way too nested
// - we work on a single html string
// - a bug, the markdown won't work for text with multiple bold or em
// - remember markdown can consist of utf-8 characters i.e. emojis so develop a solution
// based on that possibility

// Render translates markdown to HTML
func Render(markdown string) string {
	header := 0
	// spin it out into a function, don't use built-in strings.Replace, they don't work well enough
	// my approach:
	// create a string builder and alternate between
	markdown = strings.Replace(markdown, "__", "<strong>", 1)
	markdown = strings.Replace(markdown, "__", "</strong>", 1)
	markdown = strings.Replace(markdown, "_", "<em>", 1)
	markdown = strings.Replace(markdown, "_", "</em>", 1)
	pos := 0
	list := 0
	listOpened := false
	html := ""
	he := false
	for {
		char := markdown[pos]
		if char == '#' {
			for char == '#' {
				header++
				pos++
				char = markdown[pos]
			}
			if header == 7 {
				html += fmt.Sprintf("<p>%s ", strings.Repeat("#", header))
			} else if he {
				html += "# "
				header--
			} else {
				html += fmt.Sprintf("<h%d>", header)
			}
			pos++
			continue
		}
		he = true
		if char == '*' && header == 0 && strings.Contains(markdown, "\n") {
			if list == 0 {
				html += "<ul>"
			}
			list++
			if !listOpened {
				html += "<li>"
				listOpened = true
			} else {
				html += string(char) + " "
			}
			pos += 2
			continue
		}
		if char == '\n' {
			if listOpened && strings.LastIndex(markdown, "\n") == pos && strings.LastIndex(markdown, "\n") > strings.LastIndex(markdown, "*") {
				html += "</li></ul><p>"
				listOpened = false
				list = 0
			}
			if list > 0 && listOpened {
				html += "</li>"
				listOpened = false
			}
			if header > 0 {
				html += fmt.Sprintf("</h%d>", header)
				header = 0
			}
			pos++
			continue
		}
		html += string(char)
		pos++
		if pos >= len(markdown) {
			break
		}
	}
	switch {
	case header == 7:
		return html + "</p>"
	case header > 0:
		return html + fmt.Sprintf("</h%d>", header)
	}
	if list > 0 {
		return html + "</li></ul>"
	}
	if strings.Contains(html, "<p>") {
		return html + "</p>"
	}
	return "<p>" + html + "</p>"

}
