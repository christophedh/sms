let form = document.getElementById("form-sms");
let serializeForm = function (form) {
  let obj = {};
  let formData = new FormData(form);
  for (var key of formData.keys()) {
    obj[key] = formData.get(key);
  }
  return obj;
};

let sendMessage = function (target) {
  var init = {
    headers: {
      "Content-Type": "application/json",
    },
    method: "POST",
    body: JSON.stringify(serializeForm(target)),
  };

  fetch("/send-sms/", init)
    .then((resp) => {
      if (resp.ok) {
        alert("The message has been sent!");
      } else {
        resp.json().then((data) => {
          // console.log(data);
          alert("There was a problem...\n" + data.title);
        });
      }
    })
    .catch((err) => {
      alert(err);
    });
};
//
form.addEventListener("submit", function (e) {
  e.preventDefault();
  sendMessage(e.target);
});
