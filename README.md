Code Runner: A tool to help run your assignments
===================================================
Code Runner is a web application to run code pertaining to your assignment using user provided configurations.

## Features
There is one role in the application - Student.
#### Student
- Students can submit their assignments and view the output through a web page.
- They can enter the assignment working directory, commands to compile and run the assignment.
- They can also enter the optional command line arguments required to run their assignment.
- Web page accepts assignment files in compressed format. Allowed formats are .zip, .tar and .tar.gz.
      
Currently the application supports following server features:  
- Run the server.
- Accept assignment tarball, working directory, commands to compile and run, and command line arguments.
- Extract the files, compile and run the assignment and send the respective outputs to web page.

## Build and Run Server on local machine
Compile the source code using the `make` tool as shown below.
```commandline
make
```
Use the `-h` option to get information about other command-line options.

#### Port number
Use the `-port` option to specify the port number for the server to listen requests on. Below is an example.
```commandline
./code-runner-server -port <port>
```

## Code Runner as a docker image
- A docker image of code runner is used as a base image for different language specific images.
- Code Runner is not directly run, instead use the desired language specific image to run the assignments.
- Prerequisite for this is that docker engine should be installed.
- See [instructions](https://docs.docker.com/engine/installation/) for installing docker engine on different supported platforms.

Use the following docker command to run the language specific image.
```commandline
docker run --publish <port_to_expose>:<port_to_run> <language_image_tag> -port <port_to_run>
```

## Supported Languages
A fixed set of languages and their versions are currently supported.
- gcc 7
- g++ 7
- python 3.7
- java 8 & 11

## Web page to Submit Assignment
Use the exposed port number to open the web page. Example - for exposed port number as 52453, hit `localhost:52453` on browser.

![](docs/webPage.png)

Submit your assignment through the web page as shown in the figure above.
- Initially `Build` and `Run` buttons are disabled. Submit the assignment details to enable them.
- Choose the assignment tar ball.
- Enter commands to compile and run the code.
- Enter working directory to run the commands. If not provided then commands will be run at the assignment root directory.
- Add command line arguments if required to run the assignment using `Add` button.
- Click on `Submit` button to upload the details to the server.
- The status of upload will be displayed in the `Output` section.
- Click on `Build` button to compile the code.
- The status of build will be displayed in the `Output` section.
- Click on `Run` button to run the compiled code.
- The output will be displayed in the `Output` section as shown above.

#### Add Command Line Arguments
- Command Line Arguments can be added as key-value pairs in the text boxes.
- First text box is for key and second is for the value. 
- If the arguments have only value then use second text box for the value. First can be left empty.


#### Run Assignment for Interpreted Languages
- The only supported interpreted language is python.
- `Build` button remains disabled if the application being run is for any of the interpreted languages.
- Hence instead of entering command to compile, enter only command to run.
- After submitting the assignment details click on `Run` button directly to run the assignment.