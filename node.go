// https://docutils.sourceforge.io/docs/user/rst/quickstart.html

package rst

import "strings"

const (
	titleToken      = '='
	sectionToken    = '='
	subSectionToken = '-'
	listToken       = '-'
)

var listTokens = map[rune]struct{}{
	listToken: {},
	'+':       {},
	'*':       {},
}

// Node is single rst syntax representation
// Example: title, list, ...
type Node interface {
	String() string
}

type Title struct {
	text string
}

func NewTitle(text string) Title {
	text = strings.TrimSpace(text)

	return Title{
		text: text,
	}
}

func (t Title) String() string {
	var builder strings.Builder

	for i := 0; i < len(t.text); i++ {
		builder.WriteString(string(titleToken))
	}

	builder.WriteString("\n")
	builder.WriteString(t.text)
	builder.WriteString("\n")

	for i := 0; i < len(t.text); i++ {
		builder.WriteString(string(titleToken))
	}

	return builder.String()
}

type Section struct {
	text string
}

func NewSection(text string) Section {
	text = strings.TrimSpace(text)

	return Section{
		text: text,
	}
}

func (s Section) String() string {
	var builder strings.Builder

	builder.WriteString(s.text)
	builder.WriteString("\n")

	for i := 0; i < len(s.text); i++ {
		builder.WriteString(string(sectionToken))
	}

	return builder.String()
}

type SubSection struct {
	text string
}

func NewSubSection(text string) SubSection {
	text = strings.TrimSpace(text)

	return SubSection{
		text: text,
	}
}

func (s SubSection) String() string {
	var builder strings.Builder

	builder.WriteString(s.text)
	builder.WriteString("\n")

	for i := 0; i < len(s.text); i++ {
		builder.WriteString(string(subSectionToken))
	}

	return builder.String()
}

type ListItem struct {
	text string
}

func NewListItem(text string) ListItem {
	text = strings.TrimSpace(text)

	return ListItem{
		text: text,
	}
}

func (i ListItem) String() string {
	return string(listToken) + " " + i.text
}

func Equal(n1, n2 Node) bool {
	return n1.String() == n2.String()
}
