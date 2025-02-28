#!/usr/bin/env bash

# Installs Darktable on Linux
# bash <(curl -s https://raw.githubusercontent.com/photoprism/photoprism/develop/scripts/dist/install-darktable.sh)

PATH="/usr/local/sbin:/usr/sbin:/sbin:/usr/local/bin:/usr/bin:/bin:/scripts:$PATH"

# abort if not executed as root
if [[ $(id -u) != "0" ]]; then
  echo "Usage: run ${0##*/} as root" 1>&2
  exit 1
fi

if [[ $PHOTOPRISM_ARCH ]]; then
  SYSTEM_ARCH=$PHOTOPRISM_ARCH
else
  SYSTEM_ARCH=$(uname -m)
fi

DESTARCH=${2:-$SYSTEM_ARCH}

set -e

. /etc/os-release

echo "Installing Darktable for ${DESTARCH^^}..."

case $DESTARCH in
  amd64 | AMD64 | x86_64 | x86-64)
    if [[ $VERSION_CODENAME == "bullseye" ]]; then
      apt-get update
      apt-get -qq install -t bullseye-backports darktable
    elif [[ $VERSION_CODENAME == "buster" ]]; then
      apt-get update
      apt-get -qq install -t buster-backports darktable
    else
      echo "install-darktable: installing standard amd64 (Intel 64-bit) package"
      apt-get -qq install darktable
    fi
    ;;

  arm64 | ARM64 | aarch64)
    if [[ $VERSION_CODENAME == "bullseye" ]]; then
      apt-get update
      apt-get -qq install -t bullseye-backports darktable
    elif [[ $VERSION_CODENAME == "buster" ]]; then
      apt-get update
      apt-get -qq install -t buster-backports darktable
    else
      echo "install-darktable: installing standard arm64 (ARM 64-bit) package"
      apt-get -qq install darktable
    fi
    ;;

  *)
    echo "Unsupported Machine Architecture: \"$BUILD_ARCH\"" 1>&2
    exit 0
    ;;
esac

echo "Done."
