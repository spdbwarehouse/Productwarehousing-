# Productwarehousing-

## prerequisites:
# install cmake in your system to build the binary
* https://cmake.org/download/ -> download from the given url and follow the steps to add cmake to your system.

* once installation is completed comeback to same directory of the project

**********

## steps to build go binaries
* (make -f Makefile) will create the binary to bin folder
* to clean the folders where the binary is present (make -f Makefile clean)


******

## To generate docker images that service requires such as maria db etc.., through make file.
* (make -f Makefile docker) will pull mariadb instance from dockerhub and runs the service in the docker container.
* for more info please check the docker compose file present in the same directory.


### TODO:
need to automate the application as docker service 
