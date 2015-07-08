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
- https://github.com/pebbe/zmq4
- ipython

##Useful documentation
- [Making kernels for IPython](https://ipython.org/ipython-doc/dev/development/kernels.html)
- [Messaging in IPython](https://ipython.org/ipython-doc/dev/development/messaging.html)
- [IPython developer’s guide](https://ipython.org/ipython-doc/dev/development/index.html)
- [Creating Language Kernels for IPython](http://andrew.gibiansky.com/blog/ipython/ipython-kernels/)
- [The IPython notebook](https://ipython.org/ipython-doc/3/notebook/index.html)