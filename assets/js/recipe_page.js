// switchIngredientList ...
function switchIngredientList() {
    let elm = document.getElementById("recipe_ingredients");
    let sty = elm.style.display === "block" ? "none" : "block";
    elm.style.display = sty;
    document.getElementById("recipe_dialog_back").style.display = sty;
}
