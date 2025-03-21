package helper

import (
	"reflect"
	"testing"
	"time"
)

var (
	jakartaZone *time.Location = time.FixedZone("Asia/Jakarta", 25200)
	TZZone      *time.Location
)

func init() {
	// loc, err := time.LoadLocation(os.Getenv("TZ"))
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}
	TZZone = loc
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestStartOfDay ./
func TestStartOfDay(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test StartOfDay 1",
			args: InlinePointer(time.Date(2025, 1, 1, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 1, 1, 0, 0, 0, 0, jakartaZone)),
		},
		{
			name: "Test StartOfDay 2",
			args: InlinePointer(time.Date(2025, 12, 31, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 12, 31, 0, 0, 0, 0, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfDay(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfDay() = %v, want %v", got, tt.want)
			} else {
				t.Logf("StartOfDay() = %s", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestEndOfDay ./
func TestEndOfDay(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test EndOfDay 1",
			args: InlinePointer(time.Date(2025, 1, 1, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 1, 1, 23, 59, 59, 999999999, jakartaZone)),
		},
		{
			name: "Test EndOfDay 2",
			args: InlinePointer(time.Date(2025, 12, 31, 7, 8, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 12, 31, 23, 59, 59, 999999999, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfDay(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfDay() = %v, want %v", got, tt.want)
			} else {
				t.Logf("EndOfDay() = %v", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestStartOfWeek ./
func TestStartOfWeek(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test StartOfWeek 01",
			args: InlinePointer(time.Date(2025, 1, 1, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2024, 12, 30, 0, 0, 0, 0, jakartaZone)),
		},
		{
			name: "Test StartOfWeek 02",
			args: InlinePointer(time.Date(2025, 2, 28, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 2, 24, 0, 0, 0, 0, jakartaZone)),
		},
		{
			name: "Test StartOfWeek 03",
			args: InlinePointer(time.Date(2025, 3, 7, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 3, 3, 0, 0, 0, 0, jakartaZone)),
		},
		{
			name: "Test StartOfWeek 04",
			args: InlinePointer(time.Date(2025, 12, 30, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 12, 29, 0, 0, 0, 0, jakartaZone)),
		},
		{
			name: "Test StartOfWeek 05",
			args: InlinePointer(time.Date(2026, 1, 1, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 12, 29, 0, 0, 0, 0, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfWeek(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfWeek() = %v, want %v", got, tt.want)
			} else {
				t.Logf("StartOfWeek() = %s", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestEndOfWeek ./
func TestEndOfWeek(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test EndOfWeek 01",
			args: InlinePointer(time.Date(2025, 1, 1, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 1, 5, 23, 59, 59, 999999999, jakartaZone)),
		},
		{
			name: "Test EndOfWeek 02",
			args: InlinePointer(time.Date(2025, 2, 28, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 3, 2, 23, 59, 59, 999999999, jakartaZone)),
		},
		{
			name: "Test EndOfWeek 03",
			args: InlinePointer(time.Date(2025, 3, 7, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 3, 9, 23, 59, 59, 999999999, jakartaZone)),
		},
		{
			name: "Test EndOfWeek 04",
			args: InlinePointer(time.Date(2025, 12, 30, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2026, 1, 4, 23, 59, 59, 999999999, jakartaZone)),
		},
		{
			name: "Test EndOfWeek 05",
			args: InlinePointer(time.Date(2026, 1, 1, 16, 49, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2026, 1, 4, 23, 59, 59, 999999999, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfWeek(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfWeek() = %v, want %v", got, tt.want)
			} else {
				t.Logf("EndOfWeek() = %s", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestStartOfMonth ./
func TestStartOfMonth(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test StartOfMonth 1",
			args: InlinePointer(time.Date(2025, 1, 25, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 1, 1, 0, 0, 0, 0, jakartaZone)),
		},
		{
			name: "Test StartOfMonth 2",
			args: InlinePointer(time.Date(2025, 12, 31, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 12, 1, 0, 0, 0, 0, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfMonth(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfMonth() = %v, want %v", got, tt.want)
			} else {
				t.Logf("StartOfMonth() = %s", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestEndOfMonth ./
func TestEndOfMonth(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test EndOfMonth 1",
			args: InlinePointer(time.Date(2025, 1, 1, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 1, 31, 23, 59, 59, 999999999, jakartaZone)),
		},
		{
			name: "Test EndOfMonth 2",
			args: InlinePointer(time.Date(2025, 12, 7, 7, 8, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 12, 31, 23, 59, 59, 999999999, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfMonth(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfMonth() = %v, want %v", got, tt.want)
			} else {
				t.Logf("EndOfMonth() = %v", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestStartOfYear ./
func TestStartOfYear(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test StartOfYear 1",
			args: InlinePointer(time.Date(2025, 8, 25, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 1, 1, 0, 0, 0, 0, jakartaZone)),
		},
		{
			name: "Test StartOfYear 2",
			args: InlinePointer(time.Date(2026, 9, 31, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2026, 1, 1, 0, 0, 0, 0, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfYear(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfYear() = %v, want %v", got, tt.want)
			} else {
				t.Logf("StartOfYear() = %s", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestEndOfYear ./
func TestEndOfYear(t *testing.T) {
	tests := []struct {
		name string
		args *time.Time
		want *time.Time
	}{
		{
			name: "Test EndOfYear 1",
			args: InlinePointer(time.Date(2025, 9, 30, 23, 59, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2025, 12, 31, 23, 59, 59, 999999999, jakartaZone)),
		},
		{
			name: "Test EndOfYear 2",
			args: InlinePointer(time.Date(2026, 8, 7, 7, 8, 0, 0, jakartaZone)),
			want: InlinePointer(time.Date(2026, 12, 31, 23, 59, 59, 999999999, jakartaZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfYear(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfYear() = %v, want %v", got, tt.want)
			} else {
				t.Logf("EndOfYear() = %v", got.Format(time.RFC3339))
			}
		})
	}
}

// go test -v -cpu=1 -race -count=1 -failfast -run=TestStrDateToTime ./
func TestStrDateToTime(t *testing.T) {
	tests := []struct {
		name string
		args string
		want *time.Time
	}{
		{
			name: "Test StrDateToTime 1",
			args: "2025-03-06",
			want: InlinePointer(time.Date(2025, 03, 06, 0, 0, 0, 0, TZZone)),
		},
		{
			name: "Test StrDateToTime 2",
			args: "2025-02-28",
			want: InlinePointer(time.Date(2025, 02, 28, 0, 0, 0, 0, TZZone)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StrDateToTime(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrDateToTime() = %v, want %v", got, tt.want)
			} else {
				t.Logf("StrDateToTime() = %v", got.Format(time.RFC3339))
			}
		})
	}
}
