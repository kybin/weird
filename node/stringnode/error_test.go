package stringnode

import "testing"

func TestErrorPassing(t *testing.T) {
	cases := []struct {
		a    Node
		b    Node
		want string
	}{
		{
			a: NewAdd("a", AddParm{
				adds: nil, // Error!
			}),
			b: NewReplace("b", ReplaceParm{
				from: "x",
				to:   "y",
				n:    -1,
			}),
			want: "a (Add): parm.adds should not nil",
		},
	}

	for _, c := range cases {
		c.b.AddInput(c.a)

		_, err := c.b.Result()
		if err == nil {
			t.Fatal("should have error, got no error")
		}
		got := err.Error()

		if got != c.want {
			t.Fatalf("got: \"%v\", want: \"%v\"", got, c.want)
		}
	}
}
