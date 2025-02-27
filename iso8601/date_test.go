package iso8601

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_ParseDate(t *testing.T) {
	tests := []struct {
		name    string
		want    DateLike
		wantErr error
	}{
		{
			name: "20121224",
			want: Date{
				Year:  2012,
				Month: time.December,
				Day:   24,
			},
		},
		{
			name: "+20121224",
			want: Date{
				Year:  2012,
				Month: time.December,
				Day:   24,
			},
		},
		{
			name: "00001224",
			want: Date{
				Year:  0,
				Month: time.December,
				Day:   24,
			},
		},
		{
			name: "2012359",
			want: OrdinalDate{
				Year: 2012,
				Day:  359,
			},
		},
		{
			name: "+2012359",
			want: OrdinalDate{
				Year: 2012,
				Day:  359,
			},
		},
		{
			name: "2012W521",
			want: WeekDate{
				Year: 2012,
				Week: 52,
				Day:  1,
			},
		},
		{
			name: "+2012W521",
			want: WeekDate{
				Year: 2012,
				Week: 52,
				Day:  1,
			},
		},
		{
			name: "2012Q485",
			want: QuarterDate{
				Year:    2012,
				Quarter: 4,
				Day:     85,
			},
		},
		{
			name: "+2012Q485",
			want: QuarterDate{
				Year:    2012,
				Quarter: 4,
				Day:     85,
			},
		},
		{
			name: "2012-12-24",
			want: Date{
				Year:  2012,
				Month: time.December,
				Day:   24,
			},
		},
		{
			name: "+2012-12-24",
			want: Date{
				Year:  2012,
				Month: time.December,
				Day:   24,
			},
		},
		{
			name: "0000-12-24",
			want: Date{
				Year:  0,
				Month: time.December,
				Day:   24,
			},
		},
		{
			name: "+0000-12-24",
			want: Date{
				Year:  0,
				Month: time.December,
				Day:   24,
			},
		},
		{
			name: "2012-359",
			want: OrdinalDate{
				Year: 2012,
				Day:  359,
			},
		},
		{
			name: "+2012-359",
			want: OrdinalDate{
				Year: 2012,
				Day:  359,
			},
		},
		{
			name: "0000-366",
			want: OrdinalDate{
				Year: 0,
				Day:  366,
			},
		},
		{
			name: "2012-W52-1",
			want: WeekDate{
				Year: 2012,
				Week: 52,
				Day:  1,
			},
		},
		{
			name: "0000-W52-1",
			want: WeekDate{
				Year: 0,
				Week: 52,
				Day:  1,
			},
		},
		{
			name: "2012-Q4-85",
			want: QuarterDate{
				Year:    2012,
				Quarter: 4,
				Day:     85,
			},
		},
		{
			name: "0000-Q4-85",
			want: QuarterDate{
				Year:    0,
				Quarter: 4,
				Day:     85,
			},
		},
		{
			name: "20",
			wantErr: &UnexpectedTokenError{
				Value:    "20",
				Token:    humanizeDigits(2),
				Expected: "date format",
			},
		},
		{
			name: "2000/",
			wantErr: &UnexpectedTokenError{
				Value:      "2000/",
				Token:      "/",
				AfterToken: "2000",
				Expected:   "8 or more characters",
			},
		},
		{
			name: "2000Q1",
			wantErr: &UnexpectedTokenError{
				Value:      "2000Q1",
				Token:      "Q1",
				AfterToken: "2000",
				Expected:   "8 or more characters",
			},
		},
		{
			name: "2000Q12",
			wantErr: &UnexpectedTokenError{
				Value:      "2000Q12",
				Token:      "Q12",
				AfterToken: "2000",
				Expected:   "8 or more characters",
			},
		},
		{
			name: "2000W1",
			wantErr: &UnexpectedTokenError{
				Value:      "2000W1",
				Token:      "W1",
				AfterToken: "2000",
				Expected:   "8 or more characters",
			},
		},
		{
			name: "2000W12",
			wantErr: &UnexpectedTokenError{
				Value:      "2000W12",
				Token:      "W12",
				AfterToken: "2000",
				Expected:   "8 or more characters",
			},
		},
		{
			name: "2000X1234",
			wantErr: &UnexpectedTokenError{
				Value:      "2000X1234",
				Token:      "X1234",
				AfterToken: "2000",
				Expected:   "- or Q or W",
			},
		},
		{
			name: "2000Q1234",
			wantErr: &UnexpectedTokenError{
				Value:      "2000Q1234",
				Token:      humanizeDigits(4),
				AfterToken: "Q",
				Expected:   humanizeDigits(3),
			},
		},
		{
			name: "2000W1234",
			wantErr: &UnexpectedTokenError{
				Value:      "2000W1234",
				Token:      humanizeDigits(4),
				AfterToken: "W",
				Expected:   humanizeDigits(3),
			},
		},
		{
			name: "2000-Q12-34",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-Q12-34",
				Token:      humanizeDigits(2),
				AfterToken: "Q",
				Expected:   humanizeDigits(1),
			},
		},
		{
			name: "2000-Q1=34",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-Q1=34",
				Token:      "=",
				AfterToken: "Q1",
				Expected:   "-",
			},
		},
		{
			name: "2000-Q1-123",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-Q1-123",
				Token:      humanizeDigits(3),
				AfterToken: "Q1-",
				Expected:   humanizeDigits(2),
			},
		},
		{
			name: "2000-W123-34",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-W123-34",
				Token:      humanizeDigits(3),
				AfterToken: "W",
				Expected:   humanizeDigits(2),
			},
		},
		{
			name: "2000-W1-234",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-W1-234",
				Token:      humanizeDigits(1),
				AfterToken: "W",
				Expected:   humanizeDigits(2),
			},
		},
		{
			name: "2000-W12=34",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-W12=34",
				Token:      "=",
				AfterToken: "W12",
				Expected:   "-",
			},
		},
		{
			name: "2000-W12-34",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-W12-34",
				Token:      humanizeDigits(2),
				AfterToken: "W12-",
				Expected:   humanizeDigits(1),
			},
		},
		{
			name: "2000-12~34",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-12~34",
				Token:      "~",
				AfterToken: "-12",
				Expected:   "-",
			},
		},
		{
			name: "2000-12-345",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-12-345",
				Token:      humanizeDigits(3),
				AfterToken: "-12-",
				Expected:   humanizeDigits(2),
			},
		},
		{
			name: "2000-12345",
			wantErr: &UnexpectedTokenError{
				Value:      "2000-12345",
				Token:      humanizeDigits(5),
				AfterToken: "2000-",
				Expected:   "like -Q4-85 or -W52-1 or -359",
			},
		},
		// valid format but range is invalid
		{
			name: "20121324",
			wantErr: &DateLikeRangeError{
				Element: "month",
				Value:   13,
				Year:    2012,
				Min:     1,
				Max:     12,
			},
		},
		{
			name: "20110229",
			wantErr: &DateLikeRangeError{
				Element: "day of month",
				Value:   29,
				Year:    2011,
				Min:     1,
				Max:     28,
			},
		},
		{
			name: "20120230",
			wantErr: &DateLikeRangeError{
				Element: "day of month",
				Value:   30,
				Year:    2012,
				Min:     1,
				Max:     29,
			},
		},
		{
			name: "2012367",
			wantErr: &DateLikeRangeError{
				Element: "day of year",
				Value:   367,
				Year:    2012,
				Min:     1,
				Max:     366,
			},
		},
		{
			name: "2013366",
			wantErr: &DateLikeRangeError{
				Element: "day of year",
				Value:   366,
				Year:    2013,
				Min:     1,
				Max:     365,
			},
		},
		{
			name: "2012W018",
			wantErr: &DateLikeRangeError{
				Element: "day of week",
				Value:   8,
				Year:    2012,
				Min:     1,
				Max:     7,
			},
		},
		{
			name: "2012W010",
			wantErr: &DateLikeRangeError{
				Element: "day of week",
				Value:   0,
				Year:    2012,
				Min:     1,
				Max:     7,
			},
		},
		{
			name: "2012W532",
			wantErr: &DateLikeRangeError{
				Element: "week",
				Value:   53,
				Year:    2012,
				Min:     1,
				Max:     52,
			},
		},
		{
			name: "2012Q585",
			wantErr: &DateLikeRangeError{
				Element: "quarter",
				Value:   5,
				Year:    2012,
				Min:     1,
				Max:     4,
			},
		},
		{
			name: "2012Q192",
			wantErr: &DateLikeRangeError{
				Element: "day of quarter",
				Value:   92,
				Year:    2012,
				Min:     1,
				Max:     91,
			},
		},
		{
			name: "2013Q191",
			wantErr: &DateLikeRangeError{
				Element: "day of quarter",
				Value:   91,
				Year:    2013,
				Min:     1,
				Max:     90,
			},
		},
		{
			name: "20121224Hello",
			wantErr: &UnexpectedTokenError{
				Value:      "20121224Hello",
				Token:      "Hello",
				AfterToken: "20121224",
				Expected:   "20121224",
			},
		},
		{
			name: "+0000-366Hello",
			wantErr: &UnexpectedTokenError{
				Value:      "+0000-366Hello",
				Token:      "Hello",
				AfterToken: "+0000-366",
				Expected:   "+0000-366",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDate([]byte(tt.name))
			if tt.wantErr != nil {
				if diff := cmp.Diff(tt.wantErr, err); diff != "" {
					t.Errorf("error: (-want, +got)\n%s", diff)
				}
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
			if err != nil {
				t.Error(err)
			}
		})
	}
}

func Test_countDigits(t *testing.T) {
	type args struct {
		b []byte
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "start from position 0",
			args: args{
				b: []byte("20121224"),
				i: 0,
			},
			want: 8,
		},
		{
			name: "start from position 4",
			args: args{
				b: []byte("20121224"),
				i: 4,
			},
			want: 4,
		},
		{
			name: "stop at 4",
			args: args{
				b: []byte("2012T1224"),
				i: 0,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countDigits(tt.args.b, tt.args.i); got != tt.want {
				t.Errorf("countDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseNumber(t *testing.T) {
	type args struct {
		b     []byte
		start int
		width int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				b:     []byte("987654321"),
				start: 0,
				width: 9,
			},
			want: 987654321,
		},
		{
			args: args{
				b:     []byte("4321"),
				start: 0,
				width: 4,
			},
			want: 4321,
		},
		{
			args: args{
				b:     []byte("4321"),
				start: 1,
				width: 3,
			},
			want: 321,
		},
		{
			args: args{
				b:     []byte("4321"),
				start: 2,
				width: 2,
			},
			want: 21,
		},
		{
			args: args{
				b:     []byte("4321"),
				start: 3,
				width: 1,
			},
			want: 1,
		},
		{
			args: args{
				b:     []byte("4321"),
				start: 4,
				width: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s start: %d, width: %d", tt.args.b, tt.args.start, tt.args.width)
		t.Run(name, func(t *testing.T) {
			if got := parseNumber(tt.args.b, tt.args.start, tt.args.width); got != tt.want {
				t.Errorf("parseNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_IsValid(t *testing.T) {
	tests := []struct {
		name string
		date Date
		want bool
	}{
		{
			name: "invalid zero day",
			date: Date{Year: 0, Month: 0, Day: 0},
			want: false,
		},
		{
			name: "valid date",
			date: Date{Year: 2022, Month: 1, Day: 1},
			want: true,
		},
		{
			name: "invalid month",
			date: Date{Year: 2022, Month: 13, Day: 1},
			want: false,
		},
		{
			name: "invalid day",
			date: Date{Year: 2022, Month: 2, Day: 29},
			want: false,
		},
		{
			name: "valid leap year",
			date: Date{Year: 2020, Month: 2, Day: 29},
			want: true,
		},
		{
			name: "invalid leap year",
			date: Date{Year: 2021, Month: 2, Day: 29},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.IsValid(); got != tt.want {
				t.Errorf("Date.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuarterDate_IsValid(t *testing.T) {
	tests := []struct {
		name string
		date QuarterDate
		want bool
	}{
		{
			name: "invalid zero day",
			date: QuarterDate{Year: 0, Quarter: 0, Day: 0},
			want: false,
		},
		{
			name: "valid date",
			date: QuarterDate{Year: 2022, Quarter: 1, Day: 1},
			want: true,
		},
		{
			name: "invalid quarter",
			date: QuarterDate{Year: 2022, Quarter: 5, Day: 1},
			want: false,
		},
		{
			name: "invalid day",
			date: QuarterDate{Year: 2022, Quarter: 1, Day: 91},
			want: false,
		},
		{
			name: "valid last day of quarter",
			date: QuarterDate{Year: 2022, Quarter: 1, Day: 90},
			want: true,
		},
		{
			name: "valid last day of leap year quarter",
			date: QuarterDate{Year: 2020, Quarter: 1, Day: 91},
			want: true,
		},
		{
			name: "invalid last day of non-leap year quarter",
			date: QuarterDate{Year: 2021, Quarter: 1, Day: 91},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.IsValid(); got != tt.want {
				t.Errorf("QuarterDate.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekDate_IsValid(t *testing.T) {
	tests := []struct {
		name string
		date WeekDate
		want bool
	}{
		{
			name: "invalid zero day",
			date: WeekDate{Year: 0, Week: 0, Day: 0},
			want: false,
		},
		{
			name: "valid date",
			date: WeekDate{Year: 2022, Week: 1, Day: 1},
			want: true,
		},
		{
			name: "invalid week",
			date: WeekDate{Year: 2022, Week: 53, Day: 1},
			want: false,
		},
		{
			name: "invalid day",
			date: WeekDate{Year: 2022, Week: 1, Day: 8},
			want: false,
		},
		{
			name: "valid last day of year",
			date: WeekDate{Year: 2022, Week: 52, Day: 7},
			want: true,
		},
		{
			name: "valid last day of leap year",
			date: WeekDate{Year: 2020, Week: 53, Day: 7},
			want: true,
		},
		{
			name: "invalid last day of non-leap year",
			date: WeekDate{Year: 2021, Week: 53, Day: 7},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.IsValid(); got != tt.want {
				t.Errorf("WeekDate.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdinalDate_IsValid(t *testing.T) {
	tests := []struct {
		name string
		date OrdinalDate
		want bool
	}{
		{
			name: "invalid zero day",
			date: OrdinalDate{Year: 0, Day: 0},
			want: false,
		},
		{
			name: "valid date",
			date: OrdinalDate{Year: 2022, Day: 1},
			want: true,
		},
		{
			name: "invalid day",
			date: OrdinalDate{Year: 2022, Day: 366},
			want: false,
		},
		{
			name: "valid last day of non-leap year",
			date: OrdinalDate{Year: 2021, Day: 365},
			want: true,
		},
		{
			name: "valid last day of leap year",
			date: OrdinalDate{Year: 2020, Day: 366},
			want: true,
		},
		{
			name: "invalid day of leap year",
			date: OrdinalDate{Year: 2020, Day: 367},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.date.IsValid(); got != tt.want {
				t.Errorf("OrdinalDate.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdinalDate_Date(t *testing.T) {
	tests := [...]struct {
		o    OrdinalDate
		want Date
	}{
		0: {
			o: OrdinalDate{
				Year: 2008,
				Day:  58,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   27,
			},
		},
		1: {
			o: OrdinalDate{
				Year: 2008,
				Day:  59,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   28,
			},
		},
		2: {
			o: OrdinalDate{
				Year: 2008,
				Day:  60,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   29,
			},
		},
		3: {
			o: OrdinalDate{
				Year: 2008,
				Day:  61,
			},
			want: Date{
				Year:  2008,
				Month: 3,
				Day:   1,
			},
		},
		4: {
			o: OrdinalDate{
				Year: 2009,
				Day:  1,
			},
			want: Date{
				Year:  2009,
				Month: 1,
				Day:   1,
			},
		},
		5: {
			o: OrdinalDate{
				Year: 2009,
				Day:  2,
			},
			want: Date{
				Year:  2009,
				Month: 1,
				Day:   2,
			},
		},
		6: {
			o: OrdinalDate{
				Year: 2009,
				Day:  58,
			},
			want: Date{
				Year:  2009,
				Month: 2,
				Day:   27,
			},
		},
		7: {
			o: OrdinalDate{
				Year: 2009,
				Day:  59,
			},
			want: Date{
				Year:  2009,
				Month: 2,
				Day:   28,
			},
		},
		8: {
			o: OrdinalDate{
				Year: 2009,
				Day:  60,
			},
			want: Date{
				Year:  2009,
				Month: 3,
				Day:   1,
			},
		},
		9: {
			o: OrdinalDate{
				Year: 2009,
				Day:  305,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   1,
			},
		},
		10: {
			o: OrdinalDate{
				Year: 2009,
				Day:  306,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   2,
			},
		},
		11: {
			o: OrdinalDate{
				Year: 2009,
				Day:  334,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   30,
			},
		},
		12: {
			o: OrdinalDate{
				Year: 2009,
				Day:  335,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   1,
			},
		},
		13: {
			o: OrdinalDate{
				Year: 2009,
				Day:  336,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   2,
			},
		},
		14: {
			o: OrdinalDate{
				Year: 2009,
				Day:  348,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   14,
			},
		},
		15: {
			o: OrdinalDate{
				Year: 2009,
				Day:  349,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   15,
			},
		},
		16: {
			o: OrdinalDate{
				Year: 2009,
				Day:  350,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   16,
			},
		},
		17: {
			o: OrdinalDate{
				Year: 2009,
				Day:  364,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   30,
			},
		},
		18: {
			o: OrdinalDate{
				Year: 2009,
				Day:  365,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   31,
			},
		},
		19: {
			o: OrdinalDate{
				Year: 2009,
				Day:  365,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   31,
			},
		},
		20: {
			o: OrdinalDate{
				Year: 2010,
				Day:  2,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   2,
			},
		},
		21: {
			o: OrdinalDate{
				Year: 2010,
				Day:  9,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   9,
			},
		},
		22: {
			o: OrdinalDate{
				Year: 2010,
				Day:  10,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   10,
			},
		},
		23: {
			o: OrdinalDate{
				Year: 2010,
				Day:  11,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   11,
			},
		},
		24: {
			o: OrdinalDate{
				Year: 2010,
				Day:  14,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   14,
			},
		},
		25: {
			o: OrdinalDate{
				Year: 2010,
				Day:  15,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   15,
			},
		},
		26: {
			o: OrdinalDate{
				Year: 2010,
				Day:  31,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   31,
			},
		},
		27: {
			o: OrdinalDate{
				Year: 2010,
				Day:  32,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   1,
			},
		},
		28: {
			o: OrdinalDate{
				Year: 2010,
				Day:  40,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   9,
			},
		},
		29: {
			o: OrdinalDate{
				Year: 2010,
				Day:  41,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   10,
			},
		},
		30: {
			o: OrdinalDate{
				Year: 2010,
				Day:  59,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   28,
			},
		},
		31: {
			o: OrdinalDate{
				Year: 2010,
				Day:  60,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   1,
			},
		},
		32: {
			o: OrdinalDate{
				Year: 2010,
				Day:  68,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   9,
			},
		},
		33: {
			o: OrdinalDate{
				Year: 2010,
				Day:  69,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   10,
			},
		},
		34: {
			o: OrdinalDate{
				Year: 2010,
				Day:  365,
			},
			want: Date{
				Year:  2010,
				Month: 12,
				Day:   31,
			},
		},
		35: {
			o: OrdinalDate{
				Year: 2011,
				Day:  1,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   1,
			},
		},
		36: {
			o: OrdinalDate{
				Year: 2011,
				Day:  9,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   9,
			},
		},
		37: {
			o: OrdinalDate{
				Year: 2011,
				Day:  10,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   10,
			},
		},
		38: {
			o: OrdinalDate{
				Year: 2011,
				Day:  121,
			},
			want: Date{
				Year:  2011,
				Month: 5,
				Day:   1,
			},
		},
		39: {
			o: OrdinalDate{
				Year: 2011,
				Day:  365,
			},
			want: Date{
				Year:  2011,
				Month: 12,
				Day:   31,
			},
		},
		40: {
			o: OrdinalDate{
				Year: 2012,
				Day:  1,
			},
			want: Date{
				Year:  2012,
				Month: 1,
				Day:   1,
			},
		},
		41: {
			o: OrdinalDate{
				Year: 2012,
				Day:  58,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   27,
			},
		},
		42: {
			o: OrdinalDate{
				Year: 2012,
				Day:  59,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   28,
			},
		},
		43: {
			o: OrdinalDate{
				Year: 2012,
				Day:  60,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   29,
			},
		},
		44: {
			o: OrdinalDate{
				Year: 2014,
				Day:  58,
			},
			want: Date{
				Year:  2014,
				Month: 2,
				Day:   27,
			},
		},
		45: {
			o: OrdinalDate{
				Year: 2014,
				Day:  59,
			},
			want: Date{
				Year:  2014,
				Month: 2,
				Day:   28,
			},
		},
		46: {
			o: OrdinalDate{
				Year: 2014,
				Day:  60,
			},
			want: Date{
				Year:  2014,
				Month: 3,
				Day:   1,
			},
		},
		47: {
			o: OrdinalDate{
				Year: 2014,
				Day:  61,
			},
			want: Date{
				Year:  2014,
				Month: 3,
				Day:   2,
			},
		},
		48: {
			o: OrdinalDate{
				Year: 2016,
				Day:  59,
			},
			want: Date{
				Year:  2016,
				Month: 2,
				Day:   28,
			},
		},
		49: {
			o: OrdinalDate{
				Year: 2016,
				Day:  60,
			},
			want: Date{
				Year:  2016,
				Month: 2,
				Day:   29,
			},
		},
		50: {
			o: OrdinalDate{
				Year: 2016,
				Day:  61,
			},
			want: Date{
				Year:  2016,
				Month: 3,
				Day:   1,
			},
		},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("case %d", i)
		t.Run(name, func(t *testing.T) {
			got := tt.o.Date()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}

func TestQuarterDate_Date(t *testing.T) {
	tests := [...]struct {
		q    QuarterDate
		want Date
	}{
		0: {
			q: QuarterDate{
				Year:    2012,
				Quarter: 4,
				Day:     85,
			},
			want: Date{
				Year:  2012,
				Month: time.December,
				Day:   24,
			},
		},
		1: {
			q: QuarterDate{
				Year:    2000,
				Quarter: 1,
				Day:     38,
			},
			want: Date{
				Year:  2000,
				Month: 2,
				Day:   7,
			},
		},
		2: {
			q: QuarterDate{
				Year:    2000,
				Quarter: 2,
				Day:     21,
			},
			want: Date{
				Year:  2000,
				Month: 4,
				Day:   21,
			},
		},
		3: {
			q: QuarterDate{
				Year:    2000,
				Quarter: 3,
				Day:     82,
			},
			want: Date{
				Year:  2000,
				Month: 9,
				Day:   20,
			},
		},
		4: {
			q: QuarterDate{
				Year:    2000,
				Quarter: 4,
				Day:     11,
			},
			want: Date{
				Year:  2000,
				Month: 10,
				Day:   11,
			},
		},
		5: {
			q: QuarterDate{
				Year:    2001,
				Quarter: 1,
				Day:     60,
			},
			want: Date{
				Year:  2001,
				Month: 3,
				Day:   1,
			},
		},
		6: {
			q: QuarterDate{
				Year:    2001,
				Quarter: 2,
				Day:     67,
			},
			want: Date{
				Year:  2001,
				Month: 6,
				Day:   6,
			},
		},
		7: {
			q: QuarterDate{
				Year:    2001,
				Quarter: 3,
				Day:     35,
			},
			want: Date{
				Year:  2001,
				Month: 8,
				Day:   4,
			},
		},
		8: {
			q: QuarterDate{
				Year:    2001,
				Quarter: 4,
				Day:     52,
			},
			want: Date{
				Year:  2001,
				Month: 11,
				Day:   21,
			},
		},
		9: {
			q: QuarterDate{
				Year:    2002,
				Quarter: 1,
				Day:     14,
			},
			want: Date{
				Year:  2002,
				Month: 1,
				Day:   14,
			},
		},
		10: {
			q: QuarterDate{
				Year:    2002,
				Quarter: 2,
				Day:     55,
			},
			want: Date{
				Year:  2002,
				Month: 5,
				Day:   25,
			},
		},
		11: {
			q: QuarterDate{
				Year:    2002,
				Quarter: 3,
				Day:     50,
			},
			want: Date{
				Year:  2002,
				Month: 8,
				Day:   19,
			},
		},
		12: {
			q: QuarterDate{
				Year:    2002,
				Quarter: 4,
				Day:     47,
			},
			want: Date{
				Year:  2002,
				Month: 11,
				Day:   16,
			},
		},
		13: {
			q: QuarterDate{
				Year:    2003,
				Quarter: 1,
				Day:     38,
			},
			want: Date{
				Year:  2003,
				Month: 2,
				Day:   7,
			},
		},
		14: {
			q: QuarterDate{
				Year:    2003,
				Quarter: 2,
				Day:     25,
			},
			want: Date{
				Year:  2003,
				Month: 4,
				Day:   25,
			},
		},
		15: {
			q: QuarterDate{
				Year:    2003,
				Quarter: 3,
				Day:     28,
			},
			want: Date{
				Year:  2003,
				Month: 7,
				Day:   28,
			},
		},
		16: {
			q: QuarterDate{
				Year:    2008,
				Quarter: 1,
				Day:     58,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   27,
			},
		},
		17: {
			q: QuarterDate{
				Year:    2008,
				Quarter: 1,
				Day:     59,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   28,
			},
		},
		18: {
			q: QuarterDate{
				Year:    2008,
				Quarter: 1,
				Day:     60,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   29,
			},
		},
		19: {
			q: QuarterDate{
				Year:    2008,
				Quarter: 1,
				Day:     61,
			},
			want: Date{
				Year:  2008,
				Month: 3,
				Day:   1,
			},
		},
		20: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 1,
				Day:     1,
			},
			want: Date{
				Year:  2009,
				Month: 1,
				Day:   1,
			},
		},
		21: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 1,
				Day:     2,
			},
			want: Date{
				Year:  2009,
				Month: 1,
				Day:   2,
			},
		},
		22: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 1,
				Day:     58,
			},
			want: Date{
				Year:  2009,
				Month: 2,
				Day:   27,
			},
		},
		23: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 1,
				Day:     59,
			},
			want: Date{
				Year:  2009,
				Month: 2,
				Day:   28,
			},
		},
		24: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 1,
				Day:     60,
			},
			want: Date{
				Year:  2009,
				Month: 3,
				Day:   1,
			},
		},
		25: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     32,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   1,
			},
		},
		26: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     33,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   2,
			},
		},
		27: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     61,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   30,
			},
		},
		28: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     62,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   1,
			},
		},
		29: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     63,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   2,
			},
		},
		30: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     75,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   14,
			},
		},
		31: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     76,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   15,
			},
		},
		32: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     77,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   16,
			},
		},
		33: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     91,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   30,
			},
		},
		34: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     92,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   31,
			},
		},
		35: {
			q: QuarterDate{
				Year:    2009,
				Quarter: 4,
				Day:     92,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   31,
			},
		},
		36: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     2,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   2,
			},
		},
		37: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     9,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   9,
			},
		},
		38: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     10,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   10,
			},
		},
		39: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     11,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   11,
			},
		},
		40: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     14,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   14,
			},
		},
		41: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     15,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   15,
			},
		},
		42: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     31,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   31,
			},
		},
		43: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     32,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   1,
			},
		},
		44: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     40,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   9,
			},
		},
		45: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     41,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   10,
			},
		},
		46: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     59,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   28,
			},
		},
		47: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     60,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   1,
			},
		},
		48: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     68,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   9,
			},
		},
		49: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 1,
				Day:     69,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   10,
			},
		},
		50: {
			q: QuarterDate{
				Year:    2010,
				Quarter: 4,
				Day:     92,
			},
			want: Date{
				Year:  2010,
				Month: 12,
				Day:   31,
			},
		},
		51: {
			q: QuarterDate{
				Year:    2011,
				Quarter: 1,
				Day:     1,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   1,
			},
		},
		52: {
			q: QuarterDate{
				Year:    2011,
				Quarter: 1,
				Day:     9,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   9,
			},
		},
		53: {
			q: QuarterDate{
				Year:    2011,
				Quarter: 1,
				Day:     10,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   10,
			},
		},
		54: {
			q: QuarterDate{
				Year:    2011,
				Quarter: 2,
				Day:     31,
			},
			want: Date{
				Year:  2011,
				Month: 5,
				Day:   1,
			},
		},
		55: {
			q: QuarterDate{
				Year:    2011,
				Quarter: 4,
				Day:     92,
			},
			want: Date{
				Year:  2011,
				Month: 12,
				Day:   31,
			},
		},
		56: {
			q: QuarterDate{
				Year:    2012,
				Quarter: 1,
				Day:     1,
			},
			want: Date{
				Year:  2012,
				Month: 1,
				Day:   1,
			},
		},
		57: {
			q: QuarterDate{
				Year:    2012,
				Quarter: 1,
				Day:     58,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   27,
			},
		},
		58: {
			q: QuarterDate{
				Year:    2012,
				Quarter: 1,
				Day:     59,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   28,
			},
		},
		59: {
			q: QuarterDate{
				Year:    2012,
				Quarter: 1,
				Day:     60,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   29,
			},
		},
		60: {
			q: QuarterDate{
				Year:    2014,
				Quarter: 1,
				Day:     58,
			},
			want: Date{
				Year:  2014,
				Month: 2,
				Day:   27,
			},
		},
		61: {
			q: QuarterDate{
				Year:    2014,
				Quarter: 1,
				Day:     59,
			},
			want: Date{
				Year:  2014,
				Month: 2,
				Day:   28,
			},
		},
		62: {
			q: QuarterDate{
				Year:    2014,
				Quarter: 1,
				Day:     60,
			},
			want: Date{
				Year:  2014,
				Month: 3,
				Day:   1,
			},
		},
		63: {
			q: QuarterDate{
				Year:    2014,
				Quarter: 1,
				Day:     61,
			},
			want: Date{
				Year:  2014,
				Month: 3,
				Day:   2,
			},
		},
		64: {
			q: QuarterDate{
				Year:    2016,
				Quarter: 1,
				Day:     59,
			},
			want: Date{
				Year:  2016,
				Month: 2,
				Day:   28,
			},
		},
		65: {
			q: QuarterDate{
				Year:    2016,
				Quarter: 1,
				Day:     60,
			},
			want: Date{
				Year:  2016,
				Month: 2,
				Day:   29,
			},
		},
		66: {
			q: QuarterDate{
				Year:    2016,
				Quarter: 1,
				Day:     61,
			},
			want: Date{
				Year:  2016,
				Month: 3,
				Day:   1,
			},
		},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("case %d", i)
		t.Run(name, func(t *testing.T) {
			got := tt.q.Date()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}

func TestWeekDate_Date(t *testing.T) {
	tests := [...]struct {
		w    WeekDate
		want Date
	}{
		0: {
			w: WeekDate{
				Year: 2008,
				Week: 9,
				Day:  3,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   27,
			},
		},
		1: {
			w: WeekDate{
				Year: 2008,
				Week: 9,
				Day:  4,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   28,
			},
		},
		2: {
			w: WeekDate{
				Year: 2008,
				Week: 9,
				Day:  5,
			},
			want: Date{
				Year:  2008,
				Month: 2,
				Day:   29,
			},
		},
		3: {
			w: WeekDate{
				Year: 2008,
				Week: 9,
				Day:  6,
			},
			want: Date{
				Year:  2008,
				Month: 3,
				Day:   1,
			},
		},
		4: {
			w: WeekDate{
				Year: 2009,
				Week: 1,
				Day:  4,
			},
			want: Date{
				Year:  2009,
				Month: 1,
				Day:   1,
			},
		},
		5: {
			w: WeekDate{
				Year: 2009,
				Week: 1,
				Day:  5,
			},
			want: Date{
				Year:  2009,
				Month: 1,
				Day:   2,
			},
		},
		6: {
			w: WeekDate{
				Year: 2009,
				Week: 9,
				Day:  5,
			},
			want: Date{
				Year:  2009,
				Month: 2,
				Day:   27,
			},
		},
		7: {
			w: WeekDate{
				Year: 2009,
				Week: 9,
				Day:  6,
			},
			want: Date{
				Year:  2009,
				Month: 2,
				Day:   28,
			},
		},
		8: {
			w: WeekDate{
				Year: 2009,
				Week: 9,
				Day:  7,
			},
			want: Date{
				Year:  2009,
				Month: 3,
				Day:   1,
			},
		},
		9: {
			w: WeekDate{
				Year: 2009,
				Week: 44,
				Day:  7,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   1,
			},
		},
		10: {
			w: WeekDate{
				Year: 2009,
				Week: 45,
				Day:  1,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   2,
			},
		},
		11: {
			w: WeekDate{
				Year: 2009,
				Week: 49,
				Day:  1,
			},
			want: Date{
				Year:  2009,
				Month: 11,
				Day:   30,
			},
		},
		12: {
			w: WeekDate{
				Year: 2009,
				Week: 49,
				Day:  2,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   1,
			},
		},
		13: {
			w: WeekDate{
				Year: 2009,
				Week: 49,
				Day:  3,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   2,
			},
		},
		14: {
			w: WeekDate{
				Year: 2009,
				Week: 51,
				Day:  1,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   14,
			},
		},
		15: {
			w: WeekDate{
				Year: 2009,
				Week: 51,
				Day:  2,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   15,
			},
		},
		16: {
			w: WeekDate{
				Year: 2009,
				Week: 51,
				Day:  3,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   16,
			},
		},
		17: {
			w: WeekDate{
				Year: 2009,
				Week: 53,
				Day:  3,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   30,
			},
		},
		18: {
			w: WeekDate{
				Year: 2009,
				Week: 53,
				Day:  4,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   31,
			},
		},
		19: {
			w: WeekDate{
				Year: 2009,
				Week: 53,
				Day:  4,
			},
			want: Date{
				Year:  2009,
				Month: 12,
				Day:   31,
			},
		},
		20: {
			w: WeekDate{
				Year: 2009,
				Week: 53,
				Day:  6,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   2,
			},
		},
		21: {
			w: WeekDate{
				Year: 2010,
				Week: 1,
				Day:  6,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   9,
			},
		},
		22: {
			w: WeekDate{
				Year: 2010,
				Week: 1,
				Day:  7,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   10,
			},
		},
		23: {
			w: WeekDate{
				Year: 2010,
				Week: 2,
				Day:  1,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   11,
			},
		},
		24: {
			w: WeekDate{
				Year: 2010,
				Week: 2,
				Day:  4,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   14,
			},
		},
		25: {
			w: WeekDate{
				Year: 2010,
				Week: 2,
				Day:  5,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   15,
			},
		},
		26: {
			w: WeekDate{
				Year: 2010,
				Week: 4,
				Day:  7,
			},
			want: Date{
				Year:  2010,
				Month: 1,
				Day:   31,
			},
		},
		27: {
			w: WeekDate{
				Year: 2010,
				Week: 5,
				Day:  1,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   1,
			},
		},
		28: {
			w: WeekDate{
				Year: 2010,
				Week: 6,
				Day:  2,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   9,
			},
		},
		29: {
			w: WeekDate{
				Year: 2010,
				Week: 6,
				Day:  3,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   10,
			},
		},
		30: {
			w: WeekDate{
				Year: 2010,
				Week: 8,
				Day:  7,
			},
			want: Date{
				Year:  2010,
				Month: 2,
				Day:   28,
			},
		},
		31: {
			w: WeekDate{
				Year: 2010,
				Week: 9,
				Day:  1,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   1,
			},
		},
		32: {
			w: WeekDate{
				Year: 2010,
				Week: 10,
				Day:  2,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   9,
			},
		},
		33: {
			w: WeekDate{
				Year: 2010,
				Week: 10,
				Day:  3,
			},
			want: Date{
				Year:  2010,
				Month: 3,
				Day:   10,
			},
		},
		34: {
			w: WeekDate{
				Year: 2010,
				Week: 52,
				Day:  5,
			},
			want: Date{
				Year:  2010,
				Month: 12,
				Day:   31,
			},
		},
		35: {
			w: WeekDate{
				Year: 2010,
				Week: 52,
				Day:  6,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   1,
			},
		},
		36: {
			w: WeekDate{
				Year: 2011,
				Week: 1,
				Day:  7,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   9,
			},
		},
		37: {
			w: WeekDate{
				Year: 2011,
				Week: 2,
				Day:  1,
			},
			want: Date{
				Year:  2011,
				Month: 1,
				Day:   10,
			},
		},
		38: {
			w: WeekDate{
				Year: 2011,
				Week: 17,
				Day:  7,
			},
			want: Date{
				Year:  2011,
				Month: 5,
				Day:   1,
			},
		},
		39: {
			w: WeekDate{
				Year: 2011,
				Week: 52,
				Day:  6,
			},
			want: Date{
				Year:  2011,
				Month: 12,
				Day:   31,
			},
		},
		40: {
			w: WeekDate{
				Year: 2011,
				Week: 52,
				Day:  7,
			},
			want: Date{
				Year:  2012,
				Month: 1,
				Day:   1,
			},
		},
		41: {
			w: WeekDate{
				Year: 2012,
				Week: 9,
				Day:  1,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   27,
			},
		},
		42: {
			w: WeekDate{
				Year: 2012,
				Week: 9,
				Day:  2,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   28,
			},
		},
		43: {
			w: WeekDate{
				Year: 2012,
				Week: 9,
				Day:  3,
			},
			want: Date{
				Year:  2012,
				Month: 2,
				Day:   29,
			},
		},
		44: {
			w: WeekDate{
				Year: 2014,
				Week: 9,
				Day:  4,
			},
			want: Date{
				Year:  2014,
				Month: 2,
				Day:   27,
			},
		},
		45: {
			w: WeekDate{
				Year: 2014,
				Week: 9,
				Day:  5,
			},
			want: Date{
				Year:  2014,
				Month: 2,
				Day:   28,
			},
		},
		46: {
			w: WeekDate{
				Year: 2014,
				Week: 9,
				Day:  6,
			},
			want: Date{
				Year:  2014,
				Month: 3,
				Day:   1,
			},
		},
		47: {
			w: WeekDate{
				Year: 2014,
				Week: 9,
				Day:  7,
			},
			want: Date{
				Year:  2014,
				Month: 3,
				Day:   2,
			},
		},
		48: {
			w: WeekDate{
				Year: 2016,
				Week: 8,
				Day:  7,
			},
			want: Date{
				Year:  2016,
				Month: 2,
				Day:   28,
			},
		},
		49: {
			w: WeekDate{
				Year: 2016,
				Week: 9,
				Day:  1,
			},
			want: Date{
				Year:  2016,
				Month: 2,
				Day:   29,
			},
		},
		50: {
			w: WeekDate{
				Year: 2016,
				Week: 9,
				Day:  2,
			},
			want: Date{
				Year:  2016,
				Month: 3,
				Day:   1,
			},
		},
	}
	for i, tt := range tests {
		name := fmt.Sprintf("case %d", i)
		t.Run(name, func(t *testing.T) {
			got := tt.w.Date()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				w := time.Date(tt.want.Year, tt.want.Month, tt.want.Day, 0, 0, 0, 0, time.UTC)
				g := time.Date(got.Year, got.Month, got.Day, 0, 0, 0, 0, time.UTC)
				t.Errorf("(-want, +got)\n%s- %q (%s)\n+ %q (%s)", diff, w, w.Weekday(), g, g.Weekday())
			}
		})
	}
}

func TestDateLikeRangeError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  *DateLikeRangeError
		want string
	}{
		{
			name: "valid error",
			err: &DateLikeRangeError{
				Element: "month",
				Value:   13,
				Year:    2022,
				Min:     1,
				Max:     12,
			},
			want: "iso8601: 13 month is not in range 1-12 in 2022",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("DateLikeRangeError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate_Validate(t *testing.T) {
	tests := []struct {
		name    string
		d       Date
		wantErr error
	}{
		{
			name: "invalid year is less than 0",
			d: Date{
				Year:  -1,
				Month: 1,
				Day:   1,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   -1,
				Year:    -1,
				Min:     0,
				Max:     9999,
			},
		},
		{
			name: "invalid year is more than 9999",
			d: Date{
				Year:  10000,
				Month: 1,
				Day:   1,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   10000,
				Year:    10000,
				Min:     0,
				Max:     9999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.d.Validate()
			if err == nil {
				t.Fatal("expected error")
			}
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("error: (-want, +got)\n%s", diff)
			}
		})
	}
}

func TestQuarterDate_Validate(t *testing.T) {
	tests := []struct {
		name    string
		q       QuarterDate
		wantErr error
	}{
		{
			name: "invalid year is less than 0",
			q: QuarterDate{
				Year:    -1,
				Quarter: 1,
				Day:     1,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   -1,
				Year:    -1,
				Min:     0,
				Max:     9999,
			},
		},
		{
			name: "invalid year is more than 9999",
			q: QuarterDate{
				Year:    10000,
				Quarter: 1,
				Day:     1,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   10000,
				Year:    10000,
				Min:     0,
				Max:     9999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.q.Validate()
			if err == nil {
				t.Fatal("expected error")
			}
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("error: (-want, +got)\n%s", diff)
			}
		})
	}
}

func TestWeekDate_Validate(t *testing.T) {
	tests := []struct {
		name    string
		w       WeekDate
		wantErr error
	}{
		{
			name: "invalid year is less than 0",
			w: WeekDate{
				Year: -1,
				Week: 10,
				Day:  1,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   -1,
				Year:    -1,
				Min:     0,
				Max:     9999,
			},
		},
		{
			name: "invalid year is more than 9999",
			w: WeekDate{
				Year: 10000,
				Week: 10,
				Day:  1,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   10000,
				Year:    10000,
				Min:     0,
				Max:     9999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.w.Validate()
			if err == nil {
				t.Fatal("expected error")
			}
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("error: (-want, +got)\n%s", diff)
			}
		})
	}
}

func TestOrdinalDate_Validate(t *testing.T) {
	tests := []struct {
		name    string
		o       OrdinalDate
		wantErr error
	}{
		{
			name: "invalid year is less than 0",
			o: OrdinalDate{
				Year: -1,
				Day:  365,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   -1,
				Year:    -1,
				Min:     0,
				Max:     9999,
			},
		},
		{
			name: "invalid year is more than 9999",
			o: OrdinalDate{
				Year: 10000,
				Day:  365,
			},
			wantErr: &DateLikeRangeError{
				Element: "year",
				Value:   10000,
				Year:    10000,
				Min:     0,
				Max:     9999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.o.Validate()
			if err == nil {
				t.Fatal("expected error")
			}
			if diff := cmp.Diff(tt.wantErr, err); diff != "" {
				t.Errorf("error: (-want, +got)\n%s", diff)
			}
		})
	}
}

func TestDate_Date(t *testing.T) {
	want := Date{
		Year:  2020,
		Month: 10,
		Day:   1,
	}
	got := want.Date()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("(-want, +got)\n%s", diff)
	}
}

func TestDate_String(t *testing.T) {
	tests := []struct {
		d    Date
		want string
	}{
		{
			d: Date{
				Year:  2012,
				Month: 1,
				Day:   1,
			},
			want: "2012-01-01",
		},
		{
			d: Date{
				Year:  2012,
				Month: 12,
				Day:   10,
			},
			want: "2012-12-10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.d.String(); got != tt.want {
				t.Errorf("Date.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuarterDate_String(t *testing.T) {
	tests := []struct {
		q    QuarterDate
		want string
	}{
		{
			q: QuarterDate{
				Year:    2012,
				Quarter: 1,
				Day:     1,
			},
			want: "2012-Q1-01",
		},
		{
			q: QuarterDate{
				Year:    2012,
				Quarter: 4,
				Day:     85,
			},
			want: "2012-Q4-85",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.q.String(); got != tt.want {
				t.Errorf("QuarterDate.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekDate_String(t *testing.T) {
	tests := []struct {
		w    WeekDate
		want string
	}{
		{
			w: WeekDate{
				Year: 2012,
				Week: 1,
				Day:  6,
			},
			want: "2012-W01-6",
		},
		{
			w: WeekDate{
				Year: 2012,
				Week: 52,
				Day:  1,
			},
			want: "2012-W52-1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.w.String(); got != tt.want {
				t.Errorf("WeekDate.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrdinalDate_String(t *testing.T) {
	tests := []struct {
		o    OrdinalDate
		want string
	}{
		{
			o: OrdinalDate{
				Year: 2012,
				Day:  1,
			},
			want: "2012-001",
		},
		{
			o: OrdinalDate{
				Year: 2012,
				Day:  12,
			},
			want: "2012-012",
		},
		{
			o: OrdinalDate{
				Year: 2012,
				Day:  123,
			},
			want: "2012-123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			if got := tt.o.String(); got != tt.want {
				t.Errorf("OrdinalDate.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
