package utils_test

import (
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
)

func Test_CheckPWStrength(t *testing.T) {
	u := utils.NewUtil()

	type pwStrenthTest struct {
		name      string
		have      string
		minLength int
		want      bool
		got       bool
	}

	tt := []pwStrenthTest{
		{
			name:      "password is to short",
			have:      "h3ll0",
			minLength: 6,
			want:      false,
		},
		{
			name:      "password is exactly minlength and has special character as well as number",
			have:      "h3l!0",
			minLength: 5,
			want:      true,
		},
		{
			name:      "password does not have number, but is long enough",
			have:      "hell0",
			minLength: 5,
			want:      false,
		},
		{
			name:      "password does not have special character, but is long enough",
			have:      "h3llo",
			minLength: 5,
			want:      false,
		},
		{
			name:      "password is long, has special character, letter, and number",
			have:      "h3l!o",
			minLength: 5,
			want:      true,
		},
		{
			name:      "password meets length requirement only",
			have:      "helloworld",
			minLength: 5,
			want:      false,
		}, {
			name:      "password is empty",
			have:      "",
			minLength: 0,
			want:      false,
		},
		{
			name:      "missing alphanumeric",
			have:      "!@#$%12345",
			minLength: 10,
			want:      false,
		},
		{
			name:      "missing numbers",
			have:      "helloworld!",
			minLength: 10,
			want:      false,
		},
		{
			name:      "unusual slashes",
			have:      "/",
			minLength: 1,
			want:      false,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {

			v.got = u.CheckPWStrength(v.have, v.minLength)

			if v.got != v.want {
				t.Fatalf("Expected %v but got %v", v.want, v.got)
			}

		})
	}

}
