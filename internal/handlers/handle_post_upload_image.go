package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/elkcityhazard/andrew-mccall-go/pkg/utils"
)

type FileUpload struct {
	OriginalFileName string `json:"original_filename"`
	NewFileName      string `json:"new_filename"`
	FileSize         int64  `json:"filesize"`
	PathToFile       string `json:"path_to_file"`
}

func returnErr(w http.ResponseWriter, err error) {

	type error struct {
		Code  int    `json:"code"`
		Error string `json:"error"`
	}
	var e error
	e.Code = 400
	e.Error = err.Error()
	w.WriteHeader(400)
	if err = json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (hr *HandlerRepo) HandlePostUploadImage(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(2 >> 30)

	if err != nil {
		fmt.Println(err)
		returnErr(w, err)
		return
	}

	for _, fheader := range r.MultipartForm.File {
		for _, hdr := range fheader {
			infile, err := hdr.Open()

			if err != nil {
				fmt.Println(err)
				returnErr(w, err)
				return
			}

			defer infile.Close()
			f := FileUpload{}
			f.OriginalFileName = hdr.Filename
			f.FileSize = hdr.Size

			var checkFileBuf = make([]byte, 512)

			infile.Read(checkFileBuf)

			infile.Seek(0, 0)

			fileType := http.DetectContentType(checkFileBuf)

			var allowedFileTypes = []string{"image/png", "image/jpeg", "image/gif", "image/svg+xml"}
			var isValidFile = false
			for _, v := range allowedFileTypes {
				if fileType == v {
					isValidFile = true
					break
				}
				isValidFile = false
			}

			if !isValidFile {
				returnErr(w, errors.New("invalid filetype"))
				return
			}

			fp := fmt.Sprintf("uploads/%d/%s", hr.app.SessionManager.GetInt64(r.Context(), "id"), f.OriginalFileName)

			// clean to prevent any nefarious doing...
			fp = filepath.Clean(fp)

			_, err = os.Stat(filepath.Dir(fp))

			if err != nil {
				err = os.MkdirAll(filepath.Dir(fp), 0755)

				if err != nil {
					returnErr(w, err)
					return
				}
			}

			out, err := os.Create(fp)

			if err != nil {
				returnErr(w, err)
				return
			}
			defer out.Close()

			util := utils.NewUtil()

			err = util.ResizeImage(infile, out, fileType, 968)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			newInfo, err := os.Stat(fp)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			f.FileSize = newInfo.Size()
			f.NewFileName = f.OriginalFileName
			f.PathToFile = "/" + fp // the context of the directory is upload, so we are prefixing a slash for the iamge path

			if err = json.NewEncoder(w).Encode(f); err != nil {
				returnErr(w, err)
				return
			}

		}
	}

}
