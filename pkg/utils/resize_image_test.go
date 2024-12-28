package utils_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
)

func Test_ResizeImage(t *testing.T) {

	u := utils.NewUtil()

	type imageTest struct {
		name                 string
		inFilepath           string
		outFilepath          string
		mockDetectedFiletype string
		expect               error
	}

	tt := []imageTest{{
		name:                 "test png",
		inFilepath:           "./test.png",
		outFilepath:          "./out.png",
		mockDetectedFiletype: "image/png",
		expect:               nil,
	},
		{
			name:                 "test jpeg",
			inFilepath:           "./test.jpg",
			outFilepath:          "./out.jpg",
			mockDetectedFiletype: "image/jpeg",
			expect:               nil,
		},
		{
			name:                 "webp image",
			inFilepath:           "./test.webp",
			outFilepath:          "./out.jpg",
			mockDetectedFiletype: "image/jpeg",
			expect:               errors.New("invalid"),
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			infile, err := os.Open(v.inFilepath)

			if err != nil {
				t.Fatal("test png does not exist")
			}
			defer infile.Close()

			outfile, err := os.Create(v.outFilepath)

			if err != nil {
				t.Fatal("could not create outfile")
			}

			defer outfile.Close()

			err = u.ResizeImage(infile, outfile, v.mockDetectedFiletype, 250)

			if v.expect != err {
				if strings.Contains(err.Error(), v.expect.Error()) {
					return
				}
				t.Fatalf("expected %s but got %s", v.expect.Error(), err.Error())
			}

		})
	}

}
