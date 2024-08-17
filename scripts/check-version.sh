#!/bin/bash

cd "$(dirname "$0")/.."
set -e

# Read version and if it ends with -devX increment X by 1 so if its v0.0.1-dev1 it will become v0.0.1-dev2
version=$(cat version.txt)
if [[ $version == *-dev* ]]; then
    dev_version=$(echo $version | grep -oP 'dev\K[0-9]+')
    dev_version=$((dev_version + 1))
    version=$(echo $version | sed "s/dev[0-9]\+$/dev$dev_version/")
    echo $version > version.txt
fi