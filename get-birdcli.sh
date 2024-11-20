#!/bin/sh

set -u

latest_url="https://api.github.com/repos/birdcorp/cli/releases/latest"
download_url="https://github.com/birdcorp/cli/releases/download"
out_prefix="/usr/local/bin"
ask="y"
quiet="n"
version=""
os=$(uname -s | awk '{print tolower($0)}')
arch=$(uname -m)

usage() {
  echo "get-birdcli.sh - A script to help fetch the Bird CLI"
  echo ""
  echo "Flags"
  echo "-----"
  echo "  -h: print this usage message"
  echo "  -q: quiet mode, silence updates to stdout"
  echo "  -a: set the target machine architecture (amd64, arm64)"
  echo "  -s: set the target operating system (linux, macos)"
  echo "  -o: installation prefix (default: /usr/local/bin)"
  echo "  -y: accept defaults, don't ask before executing commands" 
  echo "  -v: specify the version to install"
}

latest_tag() {
  curl -s $latest_url | grep tag_name | awk '{ print $2 }' | sed 's/[",]//g'
}

untar() {
  print "Extracting release to $out_prefix/birdcli"
  $_sudo sh -c "tar -xzO birdcli > \"$out_prefix/birdcli\""
  $_sudo chmod +x "$out_prefix/birdcli"
}

print() {
  if [ "$quiet" = "n" ]; then
    echo "$@"
  fi
}

err() {
  echo "$@" >&2
  exit 1
}

while getopts "hyqa:s:o:v:" arg "$@"; do
    case "$arg" in
        h)
            usage
            exit 0
            ;;
        y)
            ask="n"
            ;;
        q)
            quiet="y"
            ;;
        a)
            arch=$OPTARG
            ;;
        s)
            os=$OPTARG
            ;;
        o)
            out_prefix=$OPTARG
            ;;
        v)
            version=$OPTARG
            ;;
        *)
            ;;
    esac
done

_user=$(whoami)
if [ "$_user" = "root" ]; then
  _sudo=""
else
  _sudo=$(which sudo)
  case $out_prefix in
  $HOME/*)
    _sudo=""
    ;;
  *)
    if [ "$_sudo" = "" ]; then
      echo "No sudo installation found, but needed to install into $out_prefix" && exit 1
    fi
    ;;
  esac
fi

case "$arch" in
x86_64)
  arch=amd64 ;;
aarch64)
  arch=arm64 ;;
*)
  ;;
esac

case "$os" in
darwin)
  os=Darwin ;;
linux)
  os=Linux ;;
*)
  err "Unsupported operating system: $os"
  ;;
esac

if test -z "$version"; then
  version=$(latest_tag)
fi

if [ "$ask" = "y" ]; then
  echo "Confirm installation:"
  echo "  Version: $version"
  echo "  OS: $os"
  echo "  Arch: $arch"
  if [ "$_sudo" != "" ]; then
    echo "  Sudo: $_sudo"
  fi
  echo "  Destination: $out_prefix/birdcli"
  echo "Proceed? [y/N]:"
  read -r reply < /dev/tty
else
  reply="y"
fi

if [ "$reply" = "y" ] || [ "$reply" = "Y" ] || [ "$reply" = "yes" ]; then
  if [ ! -d "$out_prefix" ]; then
    print "Creating directory $out_prefix"
    $_sudo mkdir -p "$out_prefix"
  fi
  curl -L -s "$download_url/$version/cli_${os}_${arch}.tar.gz" | untar
  print "birdcli executable installed to $out_prefix/birdcli"
else
  err "Installation canceled."
fi
