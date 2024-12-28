package utils_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
)

func TestSlugify(t *testing.T) {
	u := utils.NewUtil()
	type slugifyTests struct {
		name string
		have string
		want string
		got  string
	}

	tests := []slugifyTests{}

	tests = append(tests, slugifyTests{
		name: "no slash prefix or dashes",
		have: "no slash prefix or dashes",
		want: u.GenerateDateSlug() + "/no-slash-prefix-or-dashes",
	})

	tests = append(tests, slugifyTests{
		name: "has a number in the title",
		have: "99 ways to do stuff",
		want: u.GenerateDateSlug() + "/99-ways-to-do-stuff",
	})

	tests = append(tests, slugifyTests{
		name: "has multiple spaces in title",
		have: "This      Is a spaced title",
		want: u.GenerateDateSlug() + "/this-is-a-spaced-title",
	})

	tests = append(tests, slugifyTests{
		name: "Has Special Characters",
		have: "This is a question? But Also An Answer!",
		want: u.GenerateDateSlug() + "/this-is-a-question-but-also-an-answer",
	})

	tests = append(tests, slugifyTests{
		name: "Has Special Character Prefix",
		have: "!This is Cool!",
		want: u.GenerateDateSlug() + "/this-is-cool",
	})

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			v.got = u.Slugify(v.have)
			if v.got != v.want {
				t.Fatalf("Expected to be %s but got %s", v.want, v.got)
			}
		})
	}

}

func Test_GenerateDateSlug(t *testing.T) {

	u := utils.NewUtil()
	year := time.Now().Year()
	month := int(time.Now().Month())
	day := time.Now().Day()

	dateString := u.GenerateDateSlug()
	expected := fmt.Sprintf("/%04d/%02d/%02d", year, month, day)

	if !strings.EqualFold(dateString, expected) {
		t.Fatalf("dateString did not match, got %s but wanted %s", dateString, expected)
	}
}
