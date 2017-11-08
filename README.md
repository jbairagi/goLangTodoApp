# goLangTodoApp



#installation
1. sudo apt-get update
2. sudo apt-get -y upgrade
3. sudo curl -O https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
4. sudo tar -xvf go1.8.linux-amd64.tar.gz
5. sudo mv go /usr/local

#setting Go path
6. sudo nano ~/.profile
7. add this to the end
export PATH=$PATH:/usr/local/go/bin
8. source ~/.profile

9. mkdir $HOME/work
10. export GOPATH=$HOME/work

11. mkdir -p work/src/github.com/"user"/hello
12. make .go file here(nano work/src/github.com/user/hello/hello.go)

13. go install github.com/user/hello
14. sudo $GOPATH/bin/hello
