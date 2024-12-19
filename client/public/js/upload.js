let imageBase64 = null;

document.addEventListener("DOMContentLoaded", function () {
    const form = document.getElementById("documentLoader");
    const uploadInput = document.getElementById("upload");

    // Image yükleme ve base64 dönüşümü
    uploadInput.addEventListener('change', function() {
        const file = this.files[0];
        const reader = new FileReader();

        reader.onload = function(event) {
            imageBase64 = event.target.result;
            document.getElementById('imageResult').src = imageBase64; // Önyükleme resmi güncelleme
        };

        if (file) {
            reader.readAsDataURL(file);
            document.getElementById('upload-label').textContent = 'File name: ' + file.name; // Dosya adını göster
        } else {
            imageBase64 = null;
            document.getElementById('imageResult').src = "#";
            document.getElementById('upload-label').textContent = 'Choose file'; // Dosya yoksa varsayılan label
        }
    });

    form.addEventListener("submit", async function (event) {
        event.preventDefault();

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
            showModal("error", "Hata!", "Ana belge oluşturulurken bir hata oluştu!");
            return;
        }

        const mainData = await mainResponse.json();
        //console.log(JSON.stringify(mainData)); This is a control for mainResponse
        // Step 2: Create Sub Document
        const subTitle = document.querySelector('input[name="sub_title"]').value;
        const productCode = document.querySelector('input[name="product_code"]').value;
        const subMessage = document.querySelector('textarea[name="sub_message"]').value;
        const colText = document.querySelector('textarea[name="about_collection"]').value;
        const jewCare = document.querySelector('textarea[name="jewellery_care"]').value;

        const subDocumentData = {
            main_id: mainData.data,
            sub_title: subTitle,
            product_code: productCode,
            sub_message: subMessage,
        };

        if (imageBase64) {
            const cleanBase64 = imageBase64.split(",")[1];
            subDocumentData.asset = [cleanBase64];
        }
        //console.log(JSON.stringify(subDocumentData)); this is control for subDocumentData

       const subResponse = await fetch("http://127.0.0.1:8083/documenter/sub", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(subDocumentData),
        });

        if (!subResponse.ok) {
            const errorText = await subResponse.text();
            showModal("error", "Hata!", `Alt belge oluşturulurken bir hata oluştu! Hata: ${errorText}`);
            console.log("Sub Document Error: ", errorText);
            return;
        }
        const subData = await subResponse.json();
        const contentResponse = await fetch("http://127.0.0.1:8083/documenter/content", {
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

        if(!contentResponse.ok){
          showModal("error", "Hata!", `İçerik oluşturulurken bir hata oluştu!`);
          return;
        }

        showModal("success", "Başarılı!", "Tüm belgeler başarıyla oluşturuldu!");

        // Modal kapandıktan sonra sayfayı yenile
        const modalElement = document.getElementById("kt_modal_1");
        const modal = bootstrap.Modal.getInstance(modalElement);

        modalElement.addEventListener("hidden.bs.modal", function () {
            window.location.reload(); // Sayfayı yenile
        });

        modal.hide();
    });
});

function showModal(type, title, message) {
    const modalTitle = document.getElementById("kt_modal_1").querySelector(".modal-title");
    const modalBody = document.getElementById("kt_modal_1").querySelector(".modal-body");
    const modalFooter = document.getElementById("kt_modal_1").querySelector(".modal-footer");

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