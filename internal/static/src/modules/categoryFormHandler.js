import { Notyf } from "notyf";

class CategoryForm {
  constructor(id) {
    this.id = id;
    this.form = document.getElementById(this.id) || null;

    if (!this.form) {
      return null;
    }

    this.form.addEventListener("submit", this.handleFormSubmit.bind(this));
  }

  async handleFormSubmit(e) {
    try {
      e.preventDefault();
      var data = new FormData(this.form);

      let cat = {
        category_name: data.get("name"),
      };

      let catResp = await fetch("/api/v1/category", {
        method: "POST",
        headers: {
          "X-CSRF-Token": data.get("csrf_token"),
          "Content-Type": "application/json",
        },
        body: JSON.stringify(cat),
      })
        .then((resp) => resp.json())
        .then((data) => {
          if (data.error) {
          }
          return data;
        })
        .catch((err) => err);

      if (catResp.category_name != null) {
        location.assign("/admin/categories");
      }
    } catch (err) {
      throw new Error(err);
    }
  }
}

export { CategoryForm };
