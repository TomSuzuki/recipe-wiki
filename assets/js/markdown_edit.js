// markdownEdit ...マークダウン編集ページへのリンク。
function markdownEdit() {
    let url = new URL(window.location.href);
    let name = url.searchParams.get("name");
    location.href = `/markdown?name=${name}&edit=1`;
}

// markdownSave ...マークダウンを保存する。
function markdownSave() {
    const url = new URL(window.location.href);
    const name = url.searchParams.get("name")
    const request = `/markdown/data?name=${name}`;
    const form = new FormData(document.getElementById("frame_markdown"));
    const xhr = new XMLHttpRequest();

    // event
    xhr.onreadystatechange = () => {
        if (xhr.readyState === 4) if (xhr.status === 200) {
            location.href = `/markdown?name=${name}`;
        } else {
            alert("送信エラーが発生しました。")
        }
    }

    // send
    xhr.open("POST", request);
    xhr.send(form);
}