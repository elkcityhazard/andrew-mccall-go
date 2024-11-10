import { getEditor } from "./editorConfig";

let handleEditorFormFormData = function (e) {
  e.formData.append(
    "editorDelta",
    JSON.stringify(getEditor().getContents().ops),
  );

  e.formData.append(
    "editorContent",
    getEditor().getSemanticHTML(0, getEditor().getLength() - 1),
  );
};

export const registerEditorFormSubmit = function () {
  var form = document.getElementById("editorForm") || null;
  if (!form) return null;

  form.addEventListener("formdata", handleEditorFormFormData);
};
