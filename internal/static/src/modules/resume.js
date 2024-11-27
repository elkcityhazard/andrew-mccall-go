class Resume {
  constructor(formID = "") {
    this.formID = formID || null;
    this.form = document.getElementById(formID) || null;
    this.btnIds = [
      "#socialMediaBtn",
      "#skillListBtn",
      "#employmentListBtn",
      "#educationListBtn",
      "#awardListBtn",
      "#referenceListBtn",
    ];
    if (!this.formID || !this.form) return null;

    this.events();
  }

  events() {
    for (let btn of this.btnIds) {
      this.form
        .querySelector(btn)
        .addEventListener("click", this.addNewRow.bind(this));
    }
  }

  addNewRow(e) {
    e.preventDefault();
    var formRow = e.target?.previousElementSibling;
    var clonedRow = formRow.cloneNode(true);

    // reset values on new row
    clonedRow.querySelectorAll("input")?.forEach((input) => (input.value = ""));

    var fieldset = e.target.closest("fieldset");
    // button is last child so subtract 1
    return fieldset.insertBefore(
      clonedRow,
      fieldset.children[fieldset.children.length - 1],
    );
  }

  /**
   * *
   * Creates an HTML input element along with an optional label and appends it to the specified container.
   * @function htmlInputFactory
   * @param {Object} options - Options for configuring the input element.
   * @param {string} [options.tagName=""] - The type of the HTML element to be created (e.g., 'input', 'select').
   * @param {string} [options.labelText=""] - The text for the label associated with the input element.
   * @param {Object} [options.attributes] - Attributes to set on the HTML input element.
   * @param {string} [options.attributes.name=""] - The 'name' attribute for the input element.
   * @param {string} [options.attributes.value=""] - The 'value' attribute for the input element.
   * @param {string} [options.attributes.id=""] - The 'id' attribute for the input element.
   * @param {string} [options.attributes.type=""] - The 'type' attribute for the input element.
   * @param {string} [options.attributes.placeholder="text"] - The 'placeholder' attribute for the input element.
   * @param {boolean} [options.attributes.required=true] - Specifies whether the input element is required.
   * @param {Object} [options.attributes.rest] - Additional attributes that can be set on the input element.
   * @param {string} [options.containerClass="form-control"] - The CSS class to apply to the container div element.
   * @param {HTMLElement} [options.appendTo=this.form] - The DOM element to which the constructed input element will be appended.
   * @returns {Function} A function that constructs and appends the configured input and label elements to the DOM.
   */

  htmlInputFactory({
    tagName = "",
    labelText = "",
    attributes = {
      name: "",
      value: "",
      id: "",
      type: "",
      placeholder: "text",
      required: true,
      ...rest,
    },
    containerClass = "form-control",
    appendTo = this.form,
  }) {
    return function () {
      let formControlEl = document.createElement("div");
      formControlEl.classList.add(containerClass);
      let label = document.createElement("label");
      label.setAttribute("for", attributes.name);
      label.textContent = labelText;
      formControlEl.appendChild(label);
      let el = document.createElement(tagName);
      for (let [k, v] of Object.entries(attributes)) {
        switch (k) {
          case "disabled":
            if (v == false) break;
            el.setAttribute("disabled", "");
            break;
          case "required":
            if (v == false) break;
            el.setAttribute(k, "");
            break;
          default:
            el.setAttribute(k, v);
        }
      }
      formControlEl.appendChild(el);

      appendTo.appendChild(formControlEl);
    }.bind(this);
  }
}

export { Resume };
