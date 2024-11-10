import { getEditor } from "./editorConfig";

export const handleEditorLoadDelta = function () {
  var editorForm = document.getElementById("editor") || null;
  let deltaContent = editorForm.dataset.prevContent || null;
  if (!deltaContent || !editorForm) return null;
  var editor = getEditor();
  editor.setContents(JSON.parse(deltaContent));
};
