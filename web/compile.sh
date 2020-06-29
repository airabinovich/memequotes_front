#!/bin/bash

set -e

root_dir=$(dirname $0)

cd $root_dir

npm run build-dev
rm -r ../api/assets
cp -r build/ ../api/assets

cd -
