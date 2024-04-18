package main

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_incrementGoroutine(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{count: 100},
			want: 100,
		},
		{
			name: "100",
			args: args{count: 1000},
			want: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := incrementGoroutine(tt.args.count); got != tt.want {
				t.Errorf("incrementGoroutine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_incrementWithoutDefer(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{count: 100},
			want: 100,
		},
		{
			name: "1000",
			args: args{count: 1000},
			want: 1000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var wg sync.WaitGroup
			var num int64
			for range test.args.count {
				wg.Add(1)
				go increment(&num, 1, &wg)
			}

			wg.Wait()

			if got := int(num); got != test.want {
				t.Errorf("incrementGoroutine() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_incrementWithEventually(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{count: 100},
			want: 100,
		},
		{
			name: "1000",
			args: args{count: 1000},
			want: 1000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var num int64
			for range test.args.count {
				go incrementWithoutDefer(&num, 1)
			}

			require.Eventually(t, func() bool {
				return int(num) == test.want
			}, time.Second, time.Millisecond)

			if got := int(num); got != test.want {
				t.Errorf("incrementGoroutine() = %v, want %v", got, test.want)
			}
		})
	}
}

func Test_incrementWithEventuallyWithDefer(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "100",
			args: args{count: 100},
			want: 100,
		},
		{
			name: "1000",
			args: args{count: 1000},
			want: 1000,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var wg sync.WaitGroup
			var num int64
			for range test.args.count {
				wg.Add(1)
				go increment(&num, 1, &wg)
			}

			require.Eventually(t, func() bool {
				return int(num) == test.want
			}, time.Second, time.Millisecond)

			if got := int(num); got != test.want {
				t.Errorf("incrementGoroutine() = %v, want %v", got, test.want)
			}
		})
	}
}
