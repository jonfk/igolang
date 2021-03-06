igolang
=======

Golang kernel for ipython notebook.

##Installation

To install the kernel run:
```bash
$ ./install.sh

# to verify that it was installed run
$ ipython kernelspec list
Available kernels:
  python3
  golang
```

##Dependencies
###Go Dependencies
- https://github.com/pebbe/zmq4

###ipython
```bash
$ pip install ipython[all]
```
###zeromq 4.x (using 4.0.4)
On osx:
```bash
$ brew install zeromq
```
on ubuntu 14.04:
```bash
$ sudo apt-get install libzmq3-dev
```


##Useful documentation
- [Making kernels for IPython](https://ipython.org/ipython-doc/dev/development/kernels.html)
- [Messaging in IPython](https://ipython.org/ipython-doc/dev/development/messaging.html)
- [IPython developer’s guide](https://ipython.org/ipython-doc/dev/development/index.html)
- [Creating Language Kernels for IPython](http://andrew.gibiansky.com/blog/ipython/ipython-kernels/)
- [The IPython notebook](https://ipython.org/ipython-doc/3/notebook/index.html)