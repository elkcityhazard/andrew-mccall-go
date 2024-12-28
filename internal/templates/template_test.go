package templates

import (
	"strings"
	"testing"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/internal/repository/sqldbconn"
)

func Test_SetTemplateSQLDBRepo(t *testing.T) {
	var mockConnRepo *sqldbconn.SQLDbConn

	SetTemplateSQLDbRepo(mockConnRepo)
}

func Test_formatPlurarlYear(t *testing.T) {

	type formatTest struct {
		name    string
		testVal int
		want    string
		got     string
	}

	tests := []formatTest{
		{
			name:    "test 1",
			testVal: 1,
			want:    "1 year",
		},
		{
			name:    "test 0",
			testVal: 0,
			want:    "0 years",
		},
		{
			name:    "test negative year",
			testVal: -1,
			want:    "1 year",
		},
		{
			name:    "test many years",
			testVal: 5,
			want:    "5 years",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			v.got = formatPluralYear(v.testVal)

			if v.got != v.want {
				t.Fatalf("Expected %s but got %s", v.want, v.got)
			}
		})
	}

}

func Test_calculateLimit(t *testing.T) {

	type calculateLimitTest struct {
		name      string
		limit     int
		offset    int
		count     int
		increment bool
		want      int
		got       int
	}

	tt := []calculateLimitTest{
		{
			name:      "base case",
			limit:     10,
			offset:    0,
			count:     10,
			increment: true,
			want:      10,
		},
		{
			name:      "limit < 0",
			limit:     -1,
			offset:    0,
			count:     10,
			increment: true,
			want:      10,
		}, {
			name:      "limit greater than count",
			limit:     20,
			offset:    0,
			count:     10,
			increment: true,
			want:      10,
		},
		{
			name:      "limit is within parameters",
			limit:     20,
			offset:    10,
			count:     30,
			increment: true,
			want:      20,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			limit := calculateLimit(v.limit, v.offset, v.count, v.increment)

			if limit != v.want {
				t.Fatalf("Exepected %d but got %d", v.want, v.got)
			}
		})
	}
}

func Test_calculateOffset(t *testing.T) {

	type calculateOffsetTest struct {
		name      string
		limit     int
		offset    int
		count     int
		increment bool
		want      int
		got       int
	}

	tt := []calculateOffsetTest{
		{
			name:      "base case",
			limit:     10,
			offset:    10,
			count:     20,
			increment: false,
			want:      0,
		},
		{
			name:      "increment offset",
			limit:     10,
			offset:    10,
			count:     30,
			increment: true,
			want:      20,
		},
		{
			name:      "test offset + limit is greater than count",
			limit:     10,
			offset:    20,
			count:     20,
			increment: true,
			want:      20,
		},
		{
			name:      "test offset + limit is less than 0",
			limit:     10,
			offset:    -20,
			count:     20,
			increment: true,
			want:      0,
		},
		{
			name:      "test decrement and offset - limit < 0",
			limit:     20,
			offset:    10,
			count:     30,
			increment: false,
			want:      0,
		},
		{
			name:      "test decrement and offset - limit >= count",
			limit:     10,
			offset:    30,
			count:     10,
			increment: false,
			want:      30,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			v.got = calculateOffset(v.limit, v.offset, v.count, v.increment)

			if v.want != v.got {
				t.Fatalf("Expected %d, but got %d", v.want, v.got)
			}
		})
	}
}

func Test_humanDate(t *testing.T) {

	currentTimestamp := time.Now()
	expectedResult := currentTimestamp.Format("Jan 02, 2006")
	humanReadable := humanDate(currentTimestamp)
	if !strings.EqualFold(expectedResult, humanReadable) {
		t.Fatalf("Expected %s, but got %s", expectedResult, humanReadable)
	}

}

func Test_unsafeHTML(t *testing.T) {

	testHTML := `<script>alert("hello world");</script>`

	out := unsafeHTML(testHTML)

	if !strings.Contains(out, "&lt;script&gt;") {
		t.Fatal("Expected out to contain &lt;script&gt;")
	}

}

func Test_parseQuillContent(t *testing.T) {
	testContent := `<p>hello world</p>`

	out := parseQuillContent(testContent)

	if !strings.EqualFold(string(out), testContent) {
		t.Fatal("out should equal test content")
	}

}

func Test_GetTEmplateDir(t *testing.T) {
	templateDir := GetTemplateDir()

	if templateDir == nil {
		t.Fatal("Expected templateDir to not be nil but got nil")
	}
}

func Test_BuildTemplateCache(t *testing.T) {
	mockTC := BuildTemplateCache()

	if mockTC == nil {
		t.Fatal("Expected tc to not be nil, but got nil")
	}
}
