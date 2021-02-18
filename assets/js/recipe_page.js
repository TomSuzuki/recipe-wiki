window.addEventListener("DOMContentLoaded", () => {
    if (document.getElementById("frame_recipe") !== undefined) {
        let id = getParams()["id"];
        if (id !== undefined) accessServer(`/recipe/data?id=${id}`, (result) => {
            setRecipiPage(JSON.parse(result));
        });
    }
});

// setRecipiPage ...レシピページを表示する。
function setRecipiPage(json) {
    console.log(json);
}