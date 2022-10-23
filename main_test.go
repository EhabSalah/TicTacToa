package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestName(t *testing.T) {

	type args struct {
		input  [3][3]string
		player string
	}

	tests := map[string]struct {
		args args
		want bool
	}{
		"Horizontal win": {
			args: args{
				input: [3][3]string{
					{x, x, o},
					{x, o, x},
					{o, o, o},
				},
				player: o,
			},
			want: true,
		},
		"Vertical win": {
			args: args{
				input: [3][3]string{
					{x, o, o},
					{x, x, o},
					{x, o, x},
				},
				player: x,
			},
			want: true,
		},
		"Diagonal win (from top start to bottom end)": {
			args: args{
				input: [3][3]string{
					{x, o, o},
					{o, x, o},
					{o, o, x},
				},
				player: x,
			},
			want: true,
		},
		"Diagonal win (from top end to bottom start)": {
			args: args{
				input: [3][3]string{
					{x, o, x},
					{o, x, o},
					{x, o, x},
				},
				player: x,
			},
			want: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := isThePlayerWon(tt.args.input, tt.args.player)

			require.True(t, result == tt.want)
		})
	}
}

//	type WinResult int64
//
//	const (
//		Win WinResult = iota
//		Autumn
//		Winter
//		Spring
//)
