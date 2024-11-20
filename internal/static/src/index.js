import { CategoryForm } from "./modules/categoryFormHandler.js";
import { editorInit } from "./modules/editorConfig.js";
import { initNotificationRemoval } from "./modules/removeNotification.js";
import { Resume } from "./modules/resume.js";

async function init() {
  try {
    editorInit();
    initNotificationRemoval();
    new CategoryForm("addCategory");
    let resume = new Resume("resume");

    let employerTitleFactory = resume.htmlInputFactory({
      tagName: "input",
      labelText: "Employer Title: ",
      containerClass: "form-control",
      attributes: {
        name: "resume_employment_list_item_title",
        id: "resume_employment_list_item_title",
        value: "",
        type: "input",
        placeholder: "Title",
        required: true,
        disabled: false,
      },
      appendTo: resume.form,
    });

    let employerTitle = employerTitleFactory();

    let employerFromFactory = resume.htmlInputFactory({
      tagName: "input",
      labelText: "From: ",
      containerClass: "form-control",
      attributes: {
        name: "resume_employment_list_item_from_date",
        id: "resume_employment_list_item_from_date",
        value: "",
        type: "text",
        placeholder: "1999",
        required: true,
        disabled: false,
        pattern: "[0-9]{4}",
      },
      appendTo: resume.form,
    });

    let employerFromDate = employerFromFactory();
    let employerToFactory = resume.htmlInputFactory({
      tagName: "input",
      labelText: "To: ",
      containerClass: "form-control",
      attributes: {
        name: "resume_employment_list_item_to_date",
        id: "resume_employment_list_item_to_date",
        value: "",
        type: "text",
        placeholder: "1999",
        required: true,
        disabled: false,
        pattern: "[0-9]{4}",
      },
      appendTo: resume.form,
    });

    let employerToDate = employerToFactory();

    let employerJobTitleFactory = resume.htmlInputFactory({
      tagName: "input",
      labelText: "Job Title: ",
      containerClass: "form-control",
      attributes: {
        name: "resume_employment_list_item_job_title",
        id: "resume_employment_list_item_job_title",
        value: "",
        type: "text",
        placeholder: "job title",
        required: true,
        disabled: false,
        pattern: "[a-zA-Z]+",
      },
      appendTo: resume.form,
    });

    let employerJobTitle = employerJobTitleFactory();
  } catch (err) {
    throw new Error(err);
  }
}

document.addEventListener("DOMContentLoaded", init);
