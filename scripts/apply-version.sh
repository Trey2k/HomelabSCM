#!/bin/bash

cd "$(dirname "$0")/.."
set -e

version=$(cat version.txt)

file_to_apply=$1
if [ -z "$file_to_apply" ]; then
    echo "Usage: $0 <file>"
    exit 1
fi

echo "Applying version=$version to $file_to_apply"

sed -i "s/@@VERSION@@/$version/g" $file_to_apply