package utils

import "testing"

func Test_randomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"test", args{10}, "rfBd67ti3S,"},
		{"test2", args{36}, "MtYvSgD6xAV1YU00zampta8Z8S686KLkIZ0P,"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomString(tt.args.n); got != tt.want {
				t.Fatalf("randomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
