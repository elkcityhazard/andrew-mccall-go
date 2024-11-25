import { Resume } from "./resume.js";

const InitResumeFormFields = function (data = {}, form = new Resume("resume")) {
  var fieldset, legend;

  for (let [key, val] of Object.entries(data)) {
    if (key == "fieldset") {
      const { id, legend, fields } = val;

      fields.forEach((item) => console.log(item));

      return;
    }
    if (typeof val === "object") {
      return InitResumeFormFields(val, form);
    }
    console.log(key);
  }
};

export { InitResumeFormFields };
