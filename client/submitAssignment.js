let validSubmission = false;
let validRunCmd = false;
let validCompileCmd = false;
let supportedLanguage = "";
let cmdArgsEnabled = false;

// Gets the supported language for current environment from the server.
window.onload = function() {
    $.ajax({
        url: getLangHandle,
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

// Uploads the form data to server.
function uploadForm() {
    validateSubmission();
    if(validSubmission) {
        let formData = new FormData(document.getElementById(formId));
        let buildButton = document.getElementById(buildButtonId);
        let runButton = document.getElementById(runButtonId);
        buildButton.disabled = true;
        $.ajax({
            url: uploadHandle,
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

// Send the assignment using the given command.
function buildAssignment() {
    validateCompileCmd();
    if(validCompileCmd) {
        let formData = new FormData(document.getElementById(formId));
        let runButton = document.getElementById(runButtonId);
        runButton.disabled = true;
        $.ajax({
            url: buildHandle,
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

// Runs the assignment with the given command.
function runAssignment() {
    validateRunCmd();
    if(validRunCmd) {
        let formData = new FormData(document.getElementById(formId));
        $.ajax({
            url: runHandle,
            data: formData,
            processData: false,
            contentType: false,
            type: 'POST',
            cache: false,
            success: function (data) {
                console.log(data);
                let output = document.getElementById(outputId);
                output.innerHTML = data;
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

// Resets validation flags.
function resetValidations() {
    validSubmission = false;
    validCompileCmd = false;
    validRunCmd = false;
    let formError = document.getElementById(formErrorId);
    formError.innerHTML = '';
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
	output.innerHTML = '';
    return true;
}
// Validates the file upload.
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

// Validates the compile command.
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

// Validates the run command.
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
    cmdArgsEnabled = true;
   let d = document.getElementById(cmdArgsId);
   let keyName = 'key'+ argCount;
   let argName = 'arg'+ argCount;
   d.insertAdjacentHTML('beforeend',"<div class='form-group row'>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='key"+ argCount +"' name='key"+ argCount +"'>\
   </div>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='arg"+ argCount +"' name='arg"+ argCount++ +"'>\
   </div>\
   </div>");
}
