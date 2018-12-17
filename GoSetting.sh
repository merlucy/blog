#!/bin/bash

flagNumber=$#

if [ ${flagNumber} -gt 0 ]
then
 	flag1=$1
	
	## Install only if there is a flag named "install"
	if [ ${flag1} == "install" ]
	then
		echo "install golang"
		## Installing golang version 1.11.4
		curl -O https://storage.googleapis.com/golang/go1.11.4.linux-amd64.tar.gz
		tar -xvf go1.11.4.linux-amd64.tar.gz
		sudo mv go /usr/local
		mkdir $HOME/GoProjects

		## Modify GOROOT or GOPATH here if you wish to change the path settings
		echo "export GOROOT=/usr/local/go" >> ~/.profile
		echo "export GOPATH=$HOME/GoProjects" >> ~/.profile
		echo "export PATH=$PATH:$GOROOT/bin:$GOPATH/bin" >> ~/.profile
		
		## Source profile
		source ~/.profile
		
		
		## Check go version
		echo "Go Version is :"
		go version
		echo "Go Installation script complete"
	fi
fi

## To add command line tools for golang, type the following command
## go get github.org/x/tools/cmd..

## Install latest version of vim

## Install vim 8.0
## sudo add-apt-repository ppa:jonathonf/vim
## sudo apt update
## sudo apt install vim

## Plug vim-go

## sudo apt-get install mysql-client-5.7 mysql-server-5.7

## go get github.com/jinzhu/gorm
