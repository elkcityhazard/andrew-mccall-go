import { handleFileSelectClick } from "./handleFileSelectClick";

export const toolbarOptions = [
  ["bold", "italic", "underline", "strike"],
  ["blockquote", "code-block"],
  ["link", "image", "video", "formula"],
  [{ header: [1, 2, 3, 4, 5, 6, false] }],
];

export const imageHandler = function () {
  var fileSelect = document.getElementById("fileSelect") || null;
  if (!fileSelect) return null;

  fileSelect.addEventListener("change", handleFileSelectClick);

  fileSelect.click();
};
