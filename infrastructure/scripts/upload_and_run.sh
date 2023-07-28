#!/bin/bash

scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../api_gw/ ec2-user@13.212.254.114:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/idlmanagement/ ec2-user@13.229.135.254:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v1/ ec2-user@3.0.52.17:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v2/ ec2-user@54.179.111.243:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service2v1/ ec2-user@13.213.3.7:
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.212.254.114 "(cd ./api_gw; chmod +x ./api_gw; ./api_gw 172.31.0.245:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.229.135.254 "(cd ./idlmanagement; chmod +x ./idlmanagement; ./idlmanagement 172.31.0.245:8500 172.31.0.237:9999 13.212.254.114:8888) </dev/null >/dev/null 2>&1 & "
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@3.0.52.17 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.245:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@54.179.111.243 "(cd ./service1v2; chmod +x ./service1v2; ./service1v2 172.31.0.245:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.213.3.7 "(cd ./service2v1; chmod +x ./service2v1; ./service2v1 172.31.0.245:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
