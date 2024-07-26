package main

import (
	"reflect"
	"testing"
)

func TestCreateTopics(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "standard-topics",
			want: []string{"topic-letters-a-to-g", "topic-letters-h-to-n", "topic-letters-o-to-s", "topic-letters-t-to-z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTopics(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createTopics() = %v, want %v", got, tt.want)
			}
		})
	}
}
