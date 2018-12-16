#!/bin/bash

flagNumber=$#

if [ ${flagNumber} -gt 0 ]
then
 	flag1=$1
	
	if [ ${flag1} == "install" ]
	then
		echo "install golang"
		curl -O https://storage.googleapis.com/golang/go1.11.4.linux-amd64.tar.gz
		tar -xvf go1.11.4.linux-amd64.tar.gz
		sudo mv go /usr/local
		mkdir $HOME/GoProjects
		echo "export GOROOT=/usr/local/go" >> ~/.profile
		echo "export GOPATH=$HOME/GoProjects" >> ~/.profile
		echo "export PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> ~/.profile
		source ~/.profile
		echo "Go Version is :"
		go version
		echo "Go Installation script complete"
	fi
fi

## To add command line tools for golang, type the following command
## go get github.org/x/tools/cmd..

