#!/bin/bash

cd ..
terraform output -json > outputs.json
cd ..
# build.sh
cd infrastructure/scripts
upload_and_run.sh
