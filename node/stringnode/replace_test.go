package stringnode

import "testing"

func TestReplace(t *testing.T) {
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
		addNode := NewAdd("add", AddParm{
			adds: []string{c.input},
		})

		n := NewReplace("replace", ReplaceParm{
			from: c.from,
			to:   c.to,
			n:    c.n,
		})

		n.AddInput(addNode)

		got, err := n.Result()
		if err != nil {
			t.Fatalf("Replace.Result(): unexpected error: %v", err)
		}
		if got[0] != c.want {
			t.Fatalf("Replace.Result(): got \"%v\", want \"%v\"", got, c.want)
		}
	}
}
