package rst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTitleString(t *testing.T) {
	tests := []struct {
		name  string
		title Title
		want  string
	}{
		{
			name: "sample title",
			title: Title{
				text: "title",
			},
			want: "=====\ntitle\n=====",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.title.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSectionString(t *testing.T) {
	tests := []struct {
		name    string
		section Section
		want    string
	}{
		{
			name: "sample section",
			section: Section{
				text: "section A",
			},
			want: "section A\n=========",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.section.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestSubSectionString(t *testing.T) {
	tests := []struct {
		name       string
		subSection SubSection
		want       string
	}{
		{
			name: "sample sub section",
			subSection: SubSection{
				text: "sub section 1",
			},
			want: "sub section 1\n-------------",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.subSection.String()
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestListItemString(t *testing.T) {
	tests := []struct {
		name     string
		listItem ListItem
		want     string
	}{
		{
			name: "sample list item",
			listItem: ListItem{
				text: "item A",
			},
			want: "- item A",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.listItem.String()
			assert.Equal(t, tc.want, got)
		})
	}
}
