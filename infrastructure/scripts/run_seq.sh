#!/bin/bash

terraform output -json > outputs.json
../build.sh
upload_and_run.sh
