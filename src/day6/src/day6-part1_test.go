package src

import (
	"reflect"
	"testing"
)

func Test_findNearest(t *testing.T) {
	type args struct {
		x           int
		y           int
		coordinates []coordinate
	}
	tests := []struct {
		name          string
		args          args
		wantNearestId int
	}{
		{
			name:          "shouldFindNearestWhenSameCoordinates",
			args:          args{x: 5, y: 5, coordinates: []coordinate{{5, 5, 5}}},
			wantNearestId: 5,
		},
		{
			name:          "shouldFindNearestWhenDifferentCoordinates",
			args:          args{x: 0, y: 0, coordinates: []coordinate{{5, 5, 5}}},
			wantNearestId: 5,
		},
		{
			name:          "shouldFindNearestTwoCoordinates",
			args:          args{x: 0, y: 0, coordinates: []coordinate{{5, 5, 5}, {1, 1, 1}}},
			wantNearestId: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNearestId := findNearest(tt.args.x, tt.args.y, tt.args.coordinates, 1, 1); gotNearestId != tt.wantNearestId {
				t.Errorf("findNearest() = %v, want %v", gotNearestId, tt.wantNearestId)
			}
		})
	}
}

func Test_getTaxicabDistance(t *testing.T) {
	type args struct {
		firstX  int
		firstY  int
		secondX int
		secondY int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero distance", args: args{firstX: 0, firstY: 0, secondX: 0, secondY: 0}, want: 0},
		{name: "firstX positive", args: args{firstX: 1, firstY: 0, secondX: 0, secondY: 0}, want: 1},
		{name: "firstY positive", args: args{firstX: 0, firstY: 1, secondX: 0, secondY: 0}, want: 1},
		{name: "secondX positive", args: args{firstX: 0, firstY: 0, secondX: 1, secondY: 0}, want: 1},
		{name: "secondX positive", args: args{firstX: 0, firstY: 0, secondX: 0, secondY: 1}, want: 1},
		{name: "firstX negative", args: args{firstX: -1, firstY: 0, secondX: 0, secondY: 0}, want: 1},
		{name: "firstY negative", args: args{firstX: 0, firstY: -1, secondX: 0, secondY: 0}, want: 1},
		{name: "secondX negative", args: args{firstX: 0, firstY: 0, secondX: -1, secondY: 0}, want: 1},
		{name: "secondX negative", args: args{firstX: 0, firstY: 0, secondX: 0, secondY: -1}, want: 1},
		{name: "combination", args: args{firstX: 1, firstY: -1, secondX: 1, secondY: -1}, want: 0},
		{name: "combination", args: args{firstX: 1, firstY: 1, secondX: 5, secondY: 5}, want: 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTaxicabDistance(tt.args.firstX, tt.args.firstY, tt.args.secondX, tt.args.secondY); got != tt.want {
				t.Errorf("getTaxicabDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGrid(t *testing.T) {
	type args struct {
		coordinates []coordinate
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"should create grid for two coordinates",
			args{coordinates: []coordinate{{1, 0, 0}, {2, 1, 1}}},
			[][]int{{0, 0}, {0, 0}},
		},
		{"should create grid for two coordinates not starting at 0",
			args{coordinates: []coordinate{{1, 2, 2}, {2, 4, 4}}},
			[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
		},
		{"should create grid for three coordinates not starting at 0",
			args{coordinates: []coordinate{{1, 2, 2}, {2, 4, 4}, {3, 5, 5}}},
			[][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
		},
		{"should create grid for example",
			args{coordinates: []coordinate{
				{0, 1, 1},
				{1, 1, 6},
				{2, 8, 3},
				{3, 3, 4},
				{4, 5, 5},
				{5, 8, 9},
			}},
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _, _ := getGrid(tt.args.coordinates); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero no should stay same", args: args{x: 0}, want: 0},
		{name: "positive no should stay same", args: args{x: 5}, want: 5},
		{name: "negative no should be positive", args: args{x: -5}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}
