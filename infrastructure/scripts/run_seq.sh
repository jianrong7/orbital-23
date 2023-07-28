#!/bin/bash

cd ..
terraform output -json > outputs.json
cd ..

./build.sh

cd infrastructure/scripts

python generate_script.py
./upload_and_run.sh
