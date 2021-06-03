package rst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  []Node
	}{
		{
			name: "sample",
			lines: []string{
				"=====\ntitle\n=====",
				"section\n=======",
				"- item A",
				"- item B",
			},
			want: []Node{
				Title{
					text: "title",
				},
				Section{
					text: "section",
				},
				ListItem{
					text: "item A",
				},
				ListItem{
					text: "item B",
				},
			},
		},
		{
			name: "sample with sub section",
			lines: []string{
				"=====\ntitle\n=====",
				"section\n=======",
				"sub section 1\n-------------",
				"- item A",
				"- item B",
			},
			want: []Node{
				Title{
					text: "title",
				},
				Section{
					text: "section",
				},
				SubSection{
					text: "sub section 1",
				},
				ListItem{
					text: "item A",
				},
				ListItem{
					text: "item B",
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Parse(tc.lines)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestParseTitle(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		wantResult Title
		wantOK     bool
	}{
		{
			name: "sample title",
			line: "=====\ntitle\n=====",
			wantResult: Title{
				text: "title",
			},
			wantOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotOK := parseTitle(tc.line)
			assert.Equal(t, tc.wantOK, gotOK)
			if tc.wantOK {
				assert.Equal(t, tc.wantResult, gotResult)
			}
		})
	}
}

func TestParseSection(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		wantResult Section
		wantOK     bool
	}{
		{
			name: "sample section",
			line: "section\n=======",
			wantResult: Section{
				text: "section",
			},
			wantOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotOK := parseSection(tc.line)
			assert.Equal(t, tc.wantOK, gotOK)
			if tc.wantOK {
				assert.Equal(t, tc.wantResult, gotResult)
			}
		})
	}
}

func TestParseSubSection(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		wantResult SubSection
		wantOK     bool
	}{
		{
			name: "sample sub section",
			line: "sub section\n-----------",
			wantResult: SubSection{
				text: "sub section",
			},
			wantOK: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotResult, gotOK := parseSubSection(tc.line)
			assert.Equal(t, tc.wantOK, gotOK)
			if tc.wantOK {
				assert.Equal(t, tc.wantResult, gotResult)
			}
		})
	}
}

func TestParseListItem(t *testing.T) {
	tests := []struct {
		name string
		line string
		want ListItem
	}{
		{
			name: "sample list item",
			line: "- item",
			want: ListItem{
				text: "item",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := parseListItem(tc.line)
			assert.Equal(t, tc.want, got)
		})
	}
}
