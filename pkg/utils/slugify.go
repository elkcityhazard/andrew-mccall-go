package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Slugify takes in a string and attempts to convert it to a slug
// it trims spaces from the end, sets to all lowecase,
// and attemps to replace special characters, duplicate spaces, and duplicates dashes
// it also attempts to ensure that there are no trailing or leading dashes
func (u *Util) Slugify(slug string) string {

	if strings.HasPrefix(slug, "/") {
		slug = slug[1:]
	}

	slug = strings.TrimSpace(slug)
	slug = strings.ToLower(slug)

	spaceRe := regexp.MustCompile(`\s+`)

	slug = spaceRe.ReplaceAllString(slug, " ")

	specialRe := regexp.MustCompile(`[^a-zA-Z0-9-\/]`)
	slug = specialRe.ReplaceAllString(slug, "-")

	dashRe := regexp.MustCompile(`\W+`)
	slug = dashRe.ReplaceAllString(slug, "-")

	if strings.HasSuffix(slug, "-") {
		slug = slug[:len(slug)-1]
	}

	if strings.HasPrefix(slug, "-") {
		slug = slug[1:]
	}

	return fmt.Sprintf("%s/%s", u.GenerateDateSlug(), slug)
}

func (u *Util) GenerateDateSlug() string {
	var dateEls []any

	dateEls = append(dateEls, time.Now().Year(), int(time.Now().Month()), time.Now().Day())
	return fmt.Sprintf("/%04d/%02d/%02d", dateEls...)

}
