export const ComposeSlugGeneration = function () {
  var titleInput = document.getElementById("title") || null;

  if (!titleInput) return null;

  async function handleTitleInputUpdate(e) {
    if (!e.target.value) return null;

    try {
      const { value } = e.target;

      var slug = {
        value: value,
      };
      fetch("/api/v1/generate-slug", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": e.target
            .closest("form")
            .querySelector('input[name="csrf_token"]').value,
        },
        body: JSON.stringify(slug),
      })
        .then((resp) => resp.json())
        .then((data) => {
          let form = e.target.closest("form");

          if (!form) return null;

          form.querySelector('input[name="slug"]').value = data?.slug;
        });
    } catch (err) {
      console.error(err);
      throw new Error(err);
    }
  }

  titleInput.addEventListener("change", handleTitleInputUpdate);
};
