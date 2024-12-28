package utils_test

import (
	"errors"
	"testing"

	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
)

func Test_CreateJSONEnvelope(t *testing.T) {
	u := utils.NewUtil()

	type testCase struct {
		name     string
		haveKey  string
		haveData any
		want     any
	}

	tt := []testCase{
		{
			name:     "name and data provided",
			haveKey:  "msg",
			haveData: errors.New("some message").Error(),
			want:     "some message",
		},
		{
			name:     "undefined key name passed",
			haveKey:  "",
			haveData: errors.New("some message").Error(),
			want:     "some message",
		},
		{
			name:     "nil passed as data",
			haveKey:  "noData",
			haveData: nil,
			want:     make(map[string]any),
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {

			data := u.CreateJSONEnvelope(v.haveKey, v.haveData)

			if len(v.haveKey) < 1 {
				if data["data"] != v.want {
					t.Errorf("Expected %s with value %s to equal %s but got %s", data[v.haveKey], data[v.haveKey], v.want, data[v.haveKey])
				}
				return
			}

			if v.haveData == nil {

				if data[v.haveKey] == nil {
					t.Errorf("Expected %s with value %s to equal %s but got %s", data[v.haveKey], data[v.haveKey], v.want, data[v.haveKey])
				}

				return
			}

			if data[v.haveKey] != v.want {
				t.Errorf("Expected %s with value %s to equal %s but got %s", data[v.haveKey], data[v.haveKey], v.want, data[v.haveKey])
			}

		})
	}

}
