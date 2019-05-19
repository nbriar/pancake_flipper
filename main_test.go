package main

import (
	"testing"
)

func TestPancakes_SortStack(t *testing.T) {
	type fields struct {
		Stack string
		Count int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Pancakes
	}{
		{name: "flips a single pancake", fields: fields{Stack: "-"}, want: &Pancakes{Stack: "+", Count: 1}},
		{name: "flips the first pancake", fields: fields{Stack: "-+"}, want: &Pancakes{Stack: "++", Count: 1}},
		{name: "flips the last pancake with 2 steps", fields: fields{Stack: "+-"}, want: &Pancakes{Stack: "++", Count: 2}},
		{name: "doesn't flip any panckages", fields: fields{Stack: "+++"}, want: &Pancakes{Stack: "+++", Count: 0}},
		{name: "flips a stack with one face up pancake", fields: fields{Stack: "--+-"}, want: &Pancakes{Stack: "++++", Count: 3}},
		{name: "flips a stack with one face down pancake", fields: fields{Stack: "++-+"}, want: &Pancakes{Stack: "++++", Count: 2}},
		{name: "flips a mixed stack", fields: fields{Stack: "++-+++----+--+-+++----+++-+-+++-"}, want: &Pancakes{Stack: "++++++++++++++++++++++++++++++++", Count: 16}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pancakes{
				Stack: tt.fields.Stack,
			}

			if got := p.SortStack(); got.Stack != tt.want.Stack {
				t.Errorf("Pancakes.SortStack() = %v, want %v", got, tt.want)
				return
			}
			if got := p.SortStack(); got.Count != tt.want.Count {
				t.Errorf("Pancakes.SortStack() = %v, want %v: not optimized", got, tt.want)
			}
		})
	}
}
