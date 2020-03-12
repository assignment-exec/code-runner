
$("#myForm").submit(function(event) {
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

// function to validate the file upload and command line args
function validateForm() {
  var fileName = document.forms["myForm"]["file"].value;
  if (fileName == "") {
	let output = document.getElementById("fileError");
	output.innerHTML = "Please select a file to upload";
    return false;
  }
}

// function to verify the file type (should be .zip, .tar or .tar.gz)
function verifyFiles() {
    let allowedFiles = [".zip", ".tar", ".tar.gz"];
    let fileUpload = document.getElementById("file");
    let output = document.getElementById("fileError");
    let regex = new RegExp("([a-zA-Z0-9\s_\\.\-:()])+(" + allowedFiles.join('|') + ")$");
    if (!regex.test(fileUpload.value.toLowerCase())) {
        output.innerHTML = "Please upload files having extensions: " + allowedFiles.join(', ') + " only.";
        return false;
    }
    output.innerHTML = "";
    return true;
}

var x = 1

// function to add key and argument text boxes on click of "Add" button
function appendRow()
{
   var d = document.getElementById('cmdArgs');
   d.innerHTML += "<div class='form-group row'>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='key"+ x +"' name='key"+ x +"'>\
   </div>\
   <div class='col-sm-2'>\
   <input type='text' class='form-control' id='arg"+ x +"' name='arg"+ x++ +"'>\
   </div>\
   </div>";
}
