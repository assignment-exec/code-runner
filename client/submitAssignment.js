let validSubmission = false;
let validRunCmd = false;
let validCompileCmd = false;
let hostUrl = config.hostname + ":" + config.port;
let supportedLanguage = "";

window.onload = function() {
    $.ajax({
        url: hostUrl + getLangHandle,
        type: 'GET',
        cache: false,
        success: function (data) {
            console.log(data);
            supportedLanguage = data;
        },
        error: function (data) {
            console.log(data);
        }
    });
};
function uploadForm() {
    validateSubmission();
    if(validSubmission) {
        let formData = new FormData(document.getElementById(formId));
        let buildButton = document.getElementById(buildButtonId);
        let runButton = document.getElementById(runButtonId);
        buildButton.disabled = true;
        $.ajax({
            url: hostUrl + uploadHandle,
            data: formData,
            processData: false,
            contentType: false,
            type: 'POST',
            cache: false,
            success: function (data) {
                console.log(data);
                let output = document.getElementById(outputId);
                output.innerHTML = data;
                if(interpretedLanguages.find(x=>x.localeCompare(supportedLanguage) === 0))
                    runButton.disabled = false;
                else
                    buildButton.disabled = false;
                resetValidations();
            },
            error: function (data) {
                console.log(data);
                let output = document.getElementById(outputId);
                output.innerHTML = data;
            }
        });
    }
}

function buildAssignment() {
    validateCompileCmd();
    if(validCompileCmd) {
        let formData = new FormData(document.getElementById(formId));
        let runButton = document.getElementById(runButtonId);
        runButton.disabled = true;
        $.ajax({
            url: hostUrl + buildHandle,
            data: formData,
            processData: false,
            contentType: false,
            type: 'POST',
            cache: false,
            success: function (data) {
                console.log(data);
                let output = document.getElementById(outputId);
                output.innerHTML = data;

                runButton.disabled = false;
            },
            error: function (data) {
                console.log(data);
                let output = document.getElementById(outputId);
                output.innerHTML = data;
            }
        });
    }
}

function runAssignment() {
    validateRunCmd();
    if(validRunCmd) {
        let formData = new FormData(document.getElementById(formId));
        $.ajax({
            url: hostUrl + runHandle,
            data: formData,
            processData: false,
            contentType: false,
            type: 'POST',
            cache: false,
            success: function (data) {
                console.log(data);
                let output = document.getElementById(outputId);
                output.innerHTML = data;
            },
            error: function (data) {
                console.log(data);
                let output = document.getElementById(outputId);
                output.innerHTML = data;
            }
        });
    }
}

function resetValidations() {
    validSubmission = false;
    validCompileCmd = false;
    validRunCmd = false;
    let formError = document.getElementById(formErrorId);
    formError.innerHTML = "";
}

// Verifies the file type (should be .zip, .tar or .tar.gz)
function verifyFiles(fileInput) {
    let files = fileInput.files;
    let fileType = files[0].type;
    let output = document.getElementById(formErrorId);
	if (!allowedFiles.includes(fileType)) {
        output.innerHTML = "Please upload files having extensions: " + allowedExtensions.join(', ') + " only.";
        return false;
    }
	output.innerHTML = "";
    return true;
}
// Validates the file upload and command line args
function validateSubmission() {
  let fileName = document.forms[formId][fileId].value;


  if (fileName === "") {
	let output = document.getElementById(formErrorId);
	output.innerHTML = "Please select a file to upload";
    validSubmission = false;
    return;
  }
  validSubmission = true;
}

function validateCompileCmd() {
    let compileCmd = document.forms[formId][compileCmdId].value;
    if (compileCmd.trim() === "") {
        let output = document.getElementById(formErrorId);
        output.innerHTML = "Command to compile cannot be empty";
        validSubmission = false;
        return;
    }
    validCompileCmd = true;
}
// Validates the file upload and command line args
function validateRunCmd() {
    let runCmd = document.forms[formId][runCmdId].value;

    if (runCmd.trim() === "") {
        let output = document.getElementById(formErrorId);
        output.innerHTML = "Command to run cannot be empty";
        validRunCmd = false;
        return;
    }
    validRunCmd = true;
}

let argCount = 1;

// Adds key and argument text boxes on click of "Add" button
function appendRow()
{
   let d = document.getElementById(cmdArgsId);
   d.insertAdjacentHTML('beforeend',"<div class='form-group row'>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='key"+ argCount +"' name='key"+ argCount +"'>\
   </div>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='arg"+ argCount +"' name='arg"+ argCount++ +"'>\
   </div>\
   </div>");
}
