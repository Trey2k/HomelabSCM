#!/bin/bash

cd "$(dirname "$0")/.."
set -e

HOMELAB_SCM_BASE_PATH=$(pwd) HOMELAB_SCM_PORT=5080 HOMELAB_SCM_DEV_MODE=true ./build/bin/homelab-scm