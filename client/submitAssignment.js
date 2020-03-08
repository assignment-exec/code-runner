
function verifyFiles() {
    let allowedFiles = [".zip", ".tar", ".tar.gz"];
    let fileUpload = document.getElementById("file");
    let output = document.getElementById("output");
    let regex = new RegExp("([a-zA-Z0-9\s_\\.\-:])+(" + allowedFiles.join('|') + ")$");
    if (!regex.test(fileUpload.value.toLowerCase())) {
        output.innerHTML = "Please upload files having extensions: <b>" + allowedFiles.join(', ') + "</b> only.";
        return false;
    }
    output.innerHTML = "";
    return true;
}