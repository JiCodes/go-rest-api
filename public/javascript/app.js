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
    }).then(() => (document.location.href = `/fact/${factToUpdate}`)); // redirect to show
  });

// delete feature
const deleteButton = document.querySelector(".delete-btn");
const factToDelete = deleteButton && deleteButton.dataset.factId;

deleteButton &&
  deleteButton.addEventListener("click", (event) => {
    const result = confirm("Are you sure you want to delete this fact?");
    if (!result) return;

    return fetch(`/fact/${factToDelete}`, {
      method: "DELETE",
    }).then(() => (document.location.href = "/")); // redirect to index
  });
