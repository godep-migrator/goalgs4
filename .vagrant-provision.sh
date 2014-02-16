#!/bin/bash

export DEBIAN_FRONTEND=noninteractive

umask 022

set -e
set -x

apt-get update -yq
apt-get install --no-install-suggests -yq \
  build-essential \
  byobu \
  bzr \
  curl \
  git \
  mercurial \
  make \
  screen

if ! go version ; then
  curl -s https://go.googlecode.com/files/go1.2.linux-amd64.tar.gz | \
    tar -xzf - -C /usr/local
fi

mkdir -p /gopath
chown -R vagrant:vagrant /gopath

cat > /etc/profile.d/99-go.sh <<EOF
export PATH="/usr/local/go/bin:/gopath/bin:\$PATH"
export GOPATH="/gopath"
EOF

source /etc/profile.d/99-go.sh

if ! godep help &>/dev/null 2>&1 ; then
  go get -x github.com/kr/godep
fi
