(function (window) {
    ajax();
    // 发送请求到后台
    // 绑定事件
    document.querySelector('#copy').onclick = function () {
        var seq = getTextValue();
        superClipBoard.copy(seq);
        showHelpBlock(seq);
    }

    function ajax() {
        var xmlhttp;
        if (window.XMLHttpRequest) {// code for IE7+, Firefox, Chrome, Opera, Safari
            xmlhttp = new XMLHttpRequest();
        }
        else {// code for IE6, IE5
            xmlhttp = new ActiveXObject('Microsoft.XMLHTTP');
        }

        xmlhttp.open('GET', '/seq', true);
        xmlhttp.send();

        xmlhttp.onreadystatechange = function () {
            if (xmlhttp.readyState == 4 && xmlhttp.status == 200) {
                var response = JSON.parse(xmlhttp.responseText);
                console.log(response);
                setTextValue(response.seq);
            }
        }
    }

    function getTextValue() {
        return document.querySelector('#text').getAttribute('value');
    }

    function setTextValue(txt) {
        document.querySelector('#text').setAttribute('value', '【' + txt + '】');
    }

    function showHelpBlock(seq) {
        var help = document.querySelector('#helpBlock');
        help.innerText = seq + '已拷贝到剪切板';
        help.style.display = 'block';
    }
})(window);