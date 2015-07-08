#!/usr/bin/env sh

set -e

if [ ! -d "~/.ipython/kernels/golang" ]; then
    mkdir -p ~/.ipython/kernels/golang
fi

cat > ~/.ipython/kernels/golang/kernel.json << EOF
{
    "argv": ["${GOPATH}/bin/igolang", "{connection_file}"],
    "display_name": "Go",
    "language": "Go"
}
EOF
