package stringnode

import (
	"reflect"
	"testing"
)

func TestJoin(t *testing.T) {
	cases := []struct {
		data []string
		with string
		want []string
	}{
		{
			data: []string{"a", "b", "c"},
			with: "",
			want: []string{"abc"},
		},
		{
			data: []string{"a", "b", "c"},
			with: " and ",
			want: []string{"a and b and c"},
		},
		{
			data: []string{},
			with: "\n",
			want: []string{},
		},
	}

	for _, c := range cases {
		addNode := NewAdd("add1", AddParm{
			adds: c.data,
		})
		joinNode := NewJoin("join1", JoinParm{
			with: c.with,
		})
		joinNode.AddInput(addNode)

		got, err := joinNode.Result()
		if err != nil {
			t.Fatalf("should not have any error, got: %v", err)
		}

		if !reflect.DeepEqual(got, c.want) {
			t.Fatalf("got: %q, want: %q", got, c.want)
		}
	}
}
