// accessServer ...サーバーへの通信を行う。
function accessServer(path, Callback = () => { }, method = "GET") {
    let httpObj = new XMLHttpRequest();
    httpObj.onreadystatechange = function () {
        if (httpObj.readyState === 4 && httpObj.status === 200) Callback(httpObj.responseText);
    }
    httpObj.open(method, path, true);
    httpObj.send(null);
}

// getParams ...パラメータをすべて取得する。
function getParams(params = location.href) {
    const regex = /[?&]([^=#]+)=([^&#]*)/g;
    const params_obj = {};
    let match;
    while (match = regex.exec(params)) {
        params_obj[match[1]] = match[2];
    }
    return params_obj;
}