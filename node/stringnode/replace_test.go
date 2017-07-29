package stringnode

import "testing"

func TestReplaceNode(t *testing.T) {
	cases := []struct {
		input string
		from  string
		to    string
		n     int
		want  string
	}{
		{
			input: "it is mine! not yours.",
			from:  "it",
			to:    "that",
			n:     -1,
			want:  "that is mine! not yours.",
		},
	}

	for _, c := range cases {
		// create mock ReadNode.
		readNode := &ReadNode{
			done:   true,
			result: c.input,
		}

		n := NewReplaceNode(ReplaceNodeParm{
			from: c.from,
			to:   c.to,
			n:    c.n,
		})

		n.AddInput(readNode)

		got, err := n.Result()
		if err != nil {
			t.Fatalf("ReplaceNode.Result(): unexpected error: %v", err)
		}
		if got != c.want {
			t.Fatalf("ReplaceNode.Result(): got \"%v\", want \"%v\"", got, c.want)
		}
	}
}
