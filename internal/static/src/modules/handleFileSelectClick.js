import Quill from "quill";
import { getEditor } from "./editorConfig";

export const handleFileSelectClick = async function () {
  var contentImg = this.files[0] || null; // get first file
  if (!contentImg) return null;

  var csrfToken = document.getElementsByName("csrf_token")[0] || null;
  if (!csrfToken) return null;

  var form = this.closest("form") || null;
  if (!form) return null;

  var formData = new FormData(form);

  await fetch("/api/v1/upload/image", {
    method: "POST",
    body: formData,
  })
    .then((resp) => resp.json())
    .then((file) => {
      let editor = getEditor();
      const range = editor.getSelection(true);
      editor.insertText(range.index, "\n", Quill.sources.USER);
      editor.insertEmbed(
        range.index + 1,
        "image",
        {
          alt: "an inline image",
          width: 968,
          height: "auto",
          url: file?.path_to_file, // since we are going outside of the bin folder we need to prefix this,
        },
        Quill.sources.USER,
      );
      editor.setSelection(range.index + 2, Quill.sources.SILENT);
    })
    .catch((err) => console.log("error: ", err));

  var reader = new FileReader();
  reader.readAsText(contentImg, "UTF-8");
  reader.onload = handleReaderOnLoad;
};

const handleReaderOnLoad = function (e) {
  //console.log(e);
};
