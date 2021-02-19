// saveRecipe ...レシピデータをサーバーへ送る。
function saveRecipe() {
    // create json
    let json = {};
    json["id"] = null;
    json["name"] = document.getElementById("write_name").value;
    json["ingredients"] = [];
    json["steps"] = [];
    json["evaluation"] = Number(document.getElementById("write_evaluation").value);
    json["image_base64"] = "";

    // error check
    if (json["name"] === "") {
        alert("レシピ名は必ず入力してください。");
        return;
    }

    // id
    let url = new URL(window.location.href);
    let id = url.searchParams.get("id");
    if (id !== null) json["id"] = Number(id);

    // ingredients
    let ingredients = document.getElementById("write_ingredients").value.split("\n");
    for (let i in ingredients) {
        let ing = ingredients[i].split(",");
        if (ing.length !== 2) continue;
        json["ingredients"].push({
            "ingredient_name": ing[0],
            "ingredient_quantity": ing[1],
        });
    }

    // steps
    let steps = document.getElementById("write_steps").value.split("\n");
    for (let i in steps) if (steps[i] !== "") json["steps"].push(steps[i]);

    // image_base64
    let fr = new FileReader();
    fr.onload = () => {
        json["image_base64"] = encodeURIComponent(fr.result);
        send(json);
    };
    let img = document.getElementById("write_image").files;
    if (img.length > 0) fr.readAsDataURL(img[0]);
    else send(json);

    return;

    // send json
    function send(json) {
        let text = JSON.stringify(json);
        let xhr = new XMLHttpRequest;
        xhr.onload = () => {
            let json = JSON.parse(xhr.responseText);
            location.href = `/recipe?id=${json["id"]}`
        };
        xhr.onerror = () => {
            alert("送信エラーが発生しました。");
        }
        xhr.open('POST', "/recipe/data", true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.send(text);
    }
}

// imagePreview ...画像のプレビュー
function imagePreview() {
    let fr = new FileReader();
    let elm = document.getElementById("write_image");
    fr.onload = () => {
        document.getElementById("write_preview").src = fr.result
    }
    fr.readAsDataURL(elm.files[0]);
}