#!/bin/bash

cd "$(dirname "$0")/.."
set -e

rm -rf build/tmp

mkdir -p build/tmp/debian/bin
mkdir -p build/tmp/debian/etc/systemd/system
mkdir -p build/tmp/debian/DEBIAN
mkdir -p build/tmp/debian/var/opt/homelab-scm/bin
mkdir -p build/tmp/debian/var/opt/homelab-scm/git-data
mkdir -p build/tmp/debian/var/opt/homelab-scm/configs

cp init/homelab-scm.service build/tmp/debian/etc/systemd/system
cp build/package/debian/control build/tmp/debian/DEBIAN
cp build/package/debian/postinst build/tmp/debian/DEBIAN

cp build/bin/homelab-scm build/tmp/debian/var/opt/homelab-scm/bin
cp build/bin/homelab-scm-shell build/tmp/debian/var/opt/homelab-scm/bin

./scripts/apply-version.sh build/tmp/debian/DEBIAN/control

chmod +x build/tmp/debian/DEBIAN/postinst
chmod +x build/tmp/debian/var/opt/homelab-scm/bin/homelab-scm
chmod +x build/tmp/debian/var/opt/homelab-scm/bin/homelab-scm-shell
chmod -R 755 build/tmp/debian


dpkg-deb --build build/tmp/debian build/bin/homelab-scm.deb