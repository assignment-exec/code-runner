
function verifyFiles() {
    let allowedFiles = [".zip", ".tar", ".tar.gz"];
    let fileUpload = document.getElementById("file");
    let output = document.getElementById("output");
    let regex = new RegExp("([a-zA-Z0-9\s_\\.\-:()])+(" + allowedFiles.join('|') + ")$");
    if (!regex.test(fileUpload.value.toLowerCase())) {
        output.innerHTML = "Please upload files having extensions: x`" + allowedFiles.join(', ') + " only.";
        return false;
    }
    output.innerHTML = "";
    return true;
}

var x = 1

function appendRow()
{
   var d = document.getElementById('cmdArgs');
   d.innerHTML += "<div><input type='text' id='key"+ x +"' name='key"+ x +"'><input type='text' id='arg"+ x +"' name='arg"+ x++ +"'></div>";
}
