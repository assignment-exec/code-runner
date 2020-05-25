let isValidSubmission = false;
let isValidRunCmd = false;
let isValidCompileCmd = false;
let supportedLanguage = "";
let isCmdArgsEnabled = false;
let cmdArgsCount = 1;

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
    if(isValidSubmission) {
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
    if(isValidCompileCmd) {
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
    if(isValidRunCmd) {
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
    isValidSubmission = false;
    isValidCompileCmd = false;
    isValidRunCmd = false;
    let formError = document.getElementById(formErrorId);
    formError.innerHTML = '';
}

// Verifies the file type (should be .zip, .tar or .tar.gz)
function verifyFiles(fileInput) {
    let files = fileInput.files;
    let fileType = files[0].type;
    let formError = document.getElementById(formErrorId);
	if (!allowedFiles.includes(fileType)) {
        formError.innerHTML = "Please upload files having extensions: " + allowedExtensions.join(', ') + " only.";
        return false;
    }
	formError.innerHTML = '';
    return true;
}
// Validates the file upload.
function validateSubmission() {
  let fileName = document.forms[formId][fileId].value;

  if (fileName === "") {
	let formError = document.getElementById(formErrorId);
	formError.innerHTML = "Please select a file to upload";
    isValidSubmission = false;
    return;
  }
  isValidSubmission = true;
}

// Validates the compile command.
function validateCompileCmd() {
    let compileCmd = document.forms[formId][compileCmdId].value;
    if (compileCmd.trim() === "") {
        let output = document.getElementById(formErrorId);
        output.innerHTML = "Command to compile cannot be empty";
        isValidSubmission = false;
        return;
    }
    isValidCompileCmd = true;
}

// Validates the run command.
function validateRunCmd() {
    let runCmd = document.forms[formId][runCmdId].value;

    if (runCmd.trim() === "") {
        let formError = document.getElementById(formErrorId);
        formError.innerHTML = "Command to run cannot be empty";
        isValidRunCmd = false;
        return;
    }
    isValidRunCmd = true;
}

// Adds key and argument text boxes on click of "Add" button
function appendRow()
{
    isCmdArgsEnabled = true;
   let cmdArgs = document.getElementById(cmdArgsId);
   cmdArgs.insertAdjacentHTML('beforeend',"<div class='form-group row'>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='key"+ cmdArgsCount +"' name='key"+ cmdArgsCount +"'>\
   </div>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='arg"+ cmdArgsCount +"' name='arg"+ cmdArgsCount++ +"'>\
   </div>\
   </div>");
}
