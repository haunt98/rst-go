package rst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name  string
		nodes []Node
		want  string
	}{
		{
			name: "sample",
			nodes: []Node{
				NewTitle("title"),
				NewSection("section"),
				NewListItem("item A"),
				NewListItem("item B"),
			},
			want: "=====\ntitle\n=====\n\nsection\n=======\n\n- item A\n\n- item B\n",
		},
		{
			name: "sample with sub section",
			nodes: []Node{
				NewTitle("title"),
				NewSection("section"),
				NewSubSection("sub section 1"),
				NewListItem("item A"),
				NewListItem("item B"),
			},
			want: "=====\ntitle\n=====\n\nsection\n=======\n\nsub section 1\n-------------\n\n- item A\n\n- item B\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GenerateText(tc.nodes)
			assert.Equal(t, tc.want, got)
		})
	}
}
