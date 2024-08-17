#!/bin/bash

cd "$(dirname "$0")/.."
set -e

scp build/bin/homelab-scm.deb homelabscm.com:~/
ssh homelabscm.com "sudo apt install ./homelab-scm.deb"