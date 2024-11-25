import { CategoryForm } from "./modules/categoryFormHandler.js";
import { editorInit } from "./modules/editorConfig.js";
import { InitResumeFormFields } from "./modules/initResumeForm.js";
import { initNotificationRemoval } from "./modules/removeNotification.js";
import { Resume } from "./modules/resume.js";

async function init() {
  try {
    editorInit();
    initNotificationRemoval();
    import("./modules/json/resume.json").then((data) => {
      var { data = {} } = data;
      var resumeForm = new Resume("resume");
      InitResumeFormFields(data, resumeForm);
    });
    new CategoryForm("addCategory");
  } catch (err) {
    throw new Error(err);
  }
}

document.addEventListener("DOMContentLoaded", init);
