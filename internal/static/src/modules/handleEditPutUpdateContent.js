export const handlePutEditorContent = function () {
  let forms = document.getElementsByTagName("form");

  let putForm = Array.from(forms).filter(
    (el) => el.getAttribute("method").toLowerCase() == "put",
  );

  if (!putForm.length) return null;

  putForm = putForm[0];
  putForm.addEventListener("submit", handleFormData);
};

const handleFormData = async function (e) {
  try {
    e.preventDefault();

    let formData = new FormData(e.target);

    var payload = {};

    formData.forEach(function (val, key) {
      payload[key] = val;
    });

    payload.ID = parseInt(e.target.dataset.id);

    delete payload.csrf_token;
    delete payload.file;
    delete payload.featuredImage;
    console.log(payload);
    await fetch("/admin/compose/edit/" + e.target.dataset.id, {
      method: "PUT",
      headers: {
        "X-CSRF-TOKEN": formData.get("csrf_token"),
        "Content-Type": "application/json",
      },
      body: JSON.stringify(payload),
    })
      .then((resp) => {
        if (!resp.ok) {
          throw new Error(resp.status);
        }
        return resp.json();
      })
      .then((data) => {
        alert("post successfully updated");
        return;
      })
      .catch((err) => {
        throw new Error(err);
      });
  } catch (err) {
    throw new Error(err);
  }
};
