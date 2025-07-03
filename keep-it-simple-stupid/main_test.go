package main

import "testing"

func Test_getMonthKiss(t *testing.T) {
	type args struct {
		month int
	}
	tests := []struct {
		name string
		want string
		args args
	}{
		{
			name: "Jan",
			args: args{month: 1},
			want: "Jan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMonthKiss(tt.args.month); got != tt.want {
				t.Errorf("getMonthKiss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMonth(t *testing.T) {
	type args struct {
		month int
	}
	tests := []struct {
		name string
		want string
		args args
	}{
		{
			name: "Jan",
			args: args{month: 1},
			want: "Jan",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMonth(tt.args.month); got != tt.want {
				t.Errorf("getMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsKiss(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy case",
			args: args{
				s:   "123456",
				sub: "23",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsKiss(tt.args.s, tt.args.sub); got != tt.want {
				t.Errorf("containsKiss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_contains(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy case",
			args: args{
				s:   "123456",
				sub: "23",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.s, tt.args.sub); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equalKiss(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy case",
			args: args{
				a: 1,
				b: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalKiss(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("equalKiss() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equal(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "happy case",
			args: args{
				a: 1,
				b: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
