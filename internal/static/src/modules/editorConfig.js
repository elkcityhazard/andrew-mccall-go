import Quill from "quill";
import "../../node_modules/quill/dist/quill.snow.css";
import { imageHandler, toolbarOptions } from "./editorToolbarOptions";
import "./formats/imageBlot.js";
import { registerEditorFormSubmit } from "./registerEditorFormSubmit.js";
import { handleEditorLoadDelta } from "./handleEditorLoadDelta.js";
import { handlePutEditorContent } from "./handleEditPutUpdateContent.js";
let editor;

export const editorInit = function () {
  try {
    if (!document.getElementById("editor")) return null;

    const quill = new Quill("#editor", {
      modules: {
        toolbar: toolbarOptions,
      },
      theme: "snow",
    });
    registerImageHandler(quill);

    // reassignment to get from outside of module
    editor = quill;

    // register the form data handler
    registerEditorFormSubmit();
    handleEditorLoadDelta();
    handlePutEditorContent();
  } catch (err) {
    throw new Error(err);
  }
};

export const getEditor = function () {
  return editor;
};

export const registerImageHandler = function (editor = {}) {
  let toolbar = editor.getModule("toolbar");
  toolbar.addHandler("image", imageHandler);
};
