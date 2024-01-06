// use js fetch for edit feature

const editForm = document.querySelector(".form-update-fact");
const factToUpdate = editForm && editForm.dataset.factId;

editForm &&
  editForm.addEventListener("submit", (event) => {
    event.preventDefault();
    const formData = Object.fromEntries(new FormData(editForm));

    fetch(`/fact/${factToUpdate}`, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(formData),
    })
      .then(() => document.location.href = `/fact/${factToUpdate}`) // redirect to show 
  });
