let isValidSubmission = false;
let isValidRunCmd = false;
let isValidCompileCmd = false;
let supportedLanguage = "";
let isCmdArgsEnabled = false;
let cmdArgsCount = 1;

/**
 * Sends a GET request to the server for getting the supported language for current environment.
 * This function is invoked on load of the web page.
 * It stores the supported language, which is required for allowing users to either compile or directly run the assignment.
 */
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

/**
 * Validates assignment tarball file and sends a POST request to server.
 * It logs the success and error status of the operation.
 */
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
                // Enable 'Run' button if the supported language is one of the interpreted languages.
                // Else enable 'Build' button for all other languages.
                // This is done because interpreted languages don't need to be compiled.
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

/**
 * Validates the command to compile and sends POST request to server.
 * It logs and displays the response from server and also logs any
 * error encountered.
 */
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
                // Enable the 'Run' button after successful compilation.
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

/**
 * Validates the command to run and sends POST request to server.
 * It logs and displays the response from serve and also logs
 * any error encountered.
 */
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
                // Reset all validation flags after successful execution.
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

/**
 * Resets all validation flags and clears the 'formError' element.
 */
function resetValidations() {
    isValidSubmission = false;
    isValidCompileCmd = false;
    isValidRunCmd = false;
    let formError = document.getElementById(formErrorId);
    formError.innerHTML = '';
}

/**
 * Verifies the selected file input for valid file types and extensions.
 * Displays error if a file is not selected within allowed types.
 * @param fileInput - The selected file to upload
 * @returns {boolean} - A boolean flag indicating whether correct file is uploaded.
 */
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

/**
 * Validates file selection. Displays error if a file is not selected.
 */
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

/**
 * Validates command to compile. Displays error if command to compile is empty.
 */
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

/**
 * Validates command to run. Displays error if command to run is empty.
 */
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

let argCount = 1;

/**
 * Appends a row with two text boxes for adding command line arguments in key-value form.
 */
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
