package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{word:"world", want: "dlrow"},
		{word: "siapa", want: "apais"},
	}
	for _, tt := range tests {
		t.Run(tt.word,func(t *testing.T) {
			if got:= ReverseString(tt.word); got!=tt.want{
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}

}