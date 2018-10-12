## SmartShare

## Clone the source
    mkdir -p $GOPATH/src/github.com/smartshare/
    cd $GOPATH/src/github.com/smartshare/
    git clone https://github.com/smartshareproject/smartshare

## Building the source

    make smartshare

or, to build the full suite of utilities:

    make all



## Running Smartshare


```
$ smartshare
```
smartshare will connect to the main net, or use 

```
$ smartshare console
```
smartshare will connect to the main net and open an interactive console, we can use it to test smartshare function.
