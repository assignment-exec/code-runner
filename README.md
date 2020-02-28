Code Runner: An Environment to run your code
===================================================
Code Runner is a web application to run code snippet or code files on a docker environment with code specific configurations.

This application is mainly designed for running students' assignments for different courses.

## Features
- It has two ends for users - 
    - One end from where professors can create course specific docker image.
    - Second where students can submit code and view the output.
- Currently it supports only server features - 
    - Run the server
    - Accept code files in compressed format
    - Accept optional command line argument required to run code
- Allows .zip, .tar and .tar.gz compression formats.
- Server decompresses the files, reads the code and sends status response to client.

## Build and Run Server
Compile the source code using the `make` tool as shown below.
```commandline
make
```
Use the `-h` option to get information about other command-line options.

#### Port number
Use the `-port` option to specify the port number for the server to listen requests on. Below is an example.
```commandline
./code-runner-server -port "8083"
```

## Send code files and arguments to server
Use curl command to connect to the server and send request.
Use `-F` i.e. multipart/form data option to provide file and argument(s).
Below is an example.
```commandline
curl <server_ip_address:port_number> -F <compressed_filepath> -F <arg1> -F <arg2> ...
```