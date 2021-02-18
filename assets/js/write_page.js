// saveRecipe ...レシピデータをサーバーへ送る。
function saveRecipe() {

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