let mainId = null;
let imageBase64 = null;
document.addEventListener("DOMContentLoaded", function () {
  const form = document.getElementById("documentLoader");

  form.addEventListener("submit", async function (event) {
    event.preventDefault();
    //#region Place for Main Document
    // Step 1: Create Main Document
    const mainTitle = document.querySelector('input[name="main_title"]').value;
    const mainResponse = await fetch("http://127.0.0.1:8083/documenter/main", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ main_title: mainTitle }),
    });

    if (!mainResponse.ok) {
      alert("Failed to create main document");
      return;
    }

    const mainData = await mainResponse.json();
    //#region Place for Sub Document
    // Step 2: Create Sub Document
    const subTitle = document.querySelector(
      'input[name="sub_title"]'
    ).value;
    const productCode = document.querySelector(
        'input[name="product_code"]'
    ).value;
    const subMessage = document.querySelector(
      'textarea[name="sub_message"]'
    ).value;
    const colText = document.querySelector(
      'textarea[name="about_collection"]'
    ).value;
    const jewCare = document.querySelector(
      'textarea[name="jewellery_care"]'
    ).value;

    const subResponse = await fetch("http://127.0.0.1:8083/documenter/sub", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        main_id: mainData,
        sub_title: subTitle,
        product_code: productCode,
        sub_message: subMessage,
        asset: imageBase64,
      }),
    });
    if (!subResponse.ok) {
      alert("Failed to create sub document");
      const errorText = await subResponse.text();
      console.log("Sub Document Error: ", errorText);
      return;
    }
    const subData = await subResponse.json();
    await fetch("http://127.0.0.1:8083/documenter/content", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        sub_id: subData.data,
        about_collection: colText,
        jewellery_care: jewCare,
      }),
    });

    showModal("success", "Başarılı!", "Tüm belgeler başarıyla oluşturuldu!");

     // Modal kapandıktan sonra sayfayı yenile
     const modalElement = document.getElementById("kt_modal_1");
     const modal = bootstrap.Modal.getInstance(modalElement);
 
     modalElement.addEventListener("hidden.bs.modal", function () {
       window.location.reload(); // Sayfayı yenile
     });
 
     modal.hide(); // Modalı manuel olarak kapat
    
    // Process Dropzone uploads
    //myDropzone.processQueue();
    //#endregion
  });
});

// Initialize Dropzone
const myDropzone = new Dropzone("#upload", {
    url: "http://127.0.0.1:8083/documenter/sub",
    paramName: "file",
    maxFiles: 10,
    maxFilesize: 10, // MB
    addRemoveLinks: true,
    autoProcessQueue: false, // We'll process the queue manually
    accept: function (file, done) {
      var reader = new FileReader();
      reader.onload = function (event) {
        var base64String = event.target.result.split(",")[1];
        file.base64Data = base64String;
        file.imageName = file.name.split(".")[0];
        imageBase64 = base64String;
        done();
      };
      reader.readAsDataURL(file);
    },
  });
  

function showModal(type, title, message) {
  const modalTitle = document.getElementById("kt_modal_1").querySelector(".modal-title");
  const modalBody = document.getElementById("kt_modal_1").querySelector(".modal-body");
  const modalFooter = document.getElementById("kt_modal_1").querySelector(".modal-footer");

  // Başarı durumunda yeşil renk, hata durumunda kırmızı
  if (type === 'success') {
    modalTitle.textContent = title;
    modalBody.innerHTML = `<p class="text-success">${message}</p>`;
    modalFooter.innerHTML = `<button type="button" class="btn btn-light" data-bs-dismiss="modal">Kapat</button>`;
  } else if (type === 'error') {
    modalTitle.textContent = title;
    modalBody.innerHTML = `<p class="text-danger">${message}</p>`;
    modalFooter.innerHTML = `<button type="button" class="btn btn-light" data-bs-dismiss="modal">Kapat</button>`;
  }

  const modal = new bootstrap.Modal(document.getElementById("kt_modal_1"));
  modal.show();
}
