
// Creates a http POST request to the server and sends assignment tar ball and cmg args 
$("#mainForm").submit(function(event) {
    var formData = new FormData(this);
    event.stopPropagation();
    event.preventDefault();
    $.ajax({
      url: "http://localhost:8083/upload",
      data: formData,
      processData: false,
      contentType: false,
      type: 'POST',
	  cache: false,
      success: function(data) {
       console.log(data);
	   let output = document.getElementById("output");
	   output.innerHTML = data;
      },
	  error: function(data) {
       console.log(data);
	   let output = document.getElementById("output");
	   output.innerHTML = data;
      }
    });
    return false;
  });

// Verifys the file type (should be .zip, .tar or .tar.gz)
function verifyFiles(fileInput) {
    var files = fileInput.files; 
    let allowedFiles = ["application/zip", "application/x-tar","application/gzip"];
    let allowedExtensions = [".zip", ".tar", ".tar.gz"];

    var fileType = files[0].type; 
    let output = document.getElementById("fileError");
	if (!allowedFiles.includes(fileType)) {
        output.innerHTML = "Please upload files having extensions: " + allowedExtensions.join(', ') + " only.";
        return false;
    }
	output.innerHTML = "";
    return true;
}
// Validates the file upload and command line args
function validateForm() {
  var fileName = document.forms["mainForm"]["file"].value;
  if (fileName == "") {
	let output = document.getElementById("fileError");
	output.innerHTML = "Please select a file to upload";
    return false;
  }
}

var argCount = 1

// Adds key and argument text boxes on click of "Add" button
function appendRow()
{
   var d = document.getElementById('cmdArgs');
   d.insertAdjacentHTML('beforeend',"<div class='form-group row'>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='key"+ argCount +"' name='key"+ argCount +"'>\
   </div>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='arg"+ argCount +"' name='arg"+ argCount++ +"'>\
   </div>\
   </div>");
}
