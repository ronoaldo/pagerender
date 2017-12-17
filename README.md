# pagerender
Simple microservice to render web pages or web page
fragments as images.

## Requirements

To run Page Render you need Docker and the Go programming
language installed.

Docker is used in order to encapsulate the same version of
dependencies and build this project as a container.
In particular, this project works well with the Debian version
of the PhantomJS program.

Go is my favorite programming language and is used to expose
phantomjs as a web handler. One can easily replace Go with a
Node web app and reduce this dependency.

## Running from Docker Registry

This project can also be executed from the Docker public registry:

    docker run ronoaldo/pagerender:latest

## Building and running locally

If you are on Linux, you probably have GNU Make (or similar)
installed and can use the handy shortcuts:

    make run

If you don't have Make or don't want to install it, just execute
the same commands described in the Makefile:

    go build -o pagerender
    docker build -t ronoaldo/pagerender:latest .
    docker run -it ronoaldo/pagerender:latest
