package stringnode

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		want string
	}{
		{want: ""},
		{want: "hello"},
	}

	for _, c := range cases {
		addNode := NewAdd(AddParm{
			adds: []string{c.want},
		})

		data, err := addNode.Result()
		if err != nil {
			t.Fatal("should not have any error")
		}
		if data == nil {
			t.Fatal("data should not nil")
		}
		if len(data) != 1 {
			t.Fatalf("length of data should 1 but it's %v", len(data))
		}

		got := data[0]
		if got != c.want {
			t.Fatalf("got: %v, want: %v", got, c.want)
		}
	}
}

func TestTwoAdd(t *testing.T) {
	cases := []struct {
		a    []string
		b    []string
		want []string
	}{
		{
			a:    []string{"a"},
			b:    []string{"b"},
			want: []string{"a", "b"},
		},
		{
			a:    []string{"a", "b"},
			b:    []string{"c"},
			want: []string{"a", "b", "c"},
		},
	}

	for _, c := range cases {
		firstNode := NewAdd(AddParm{
			adds: c.a,
		})
		secondNode := NewAdd(AddParm{
			adds: c.b,
		})
		secondNode.AddInput(firstNode)

		data, err := secondNode.Result()
		if err != nil {
			t.Fatal("should not have any error")
		}
		if data == nil {
			t.Fatal("data should not nil")
		}
		for i, got := range data {
			if got != c.want[i] {
				t.Fatalf("got: %v, want: %v", got, c.want[i])
			}
		}
	}
}
