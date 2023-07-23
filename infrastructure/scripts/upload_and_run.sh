#!/bin/bash

scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../api_gw/ ec2-user@18.139.224.140:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/idlmanagement/ ec2-user@18.140.56.168:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v1/ ec2-user@13.229.85.143:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v1/ ec2-user@13.250.45.115:
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@18.139.224.140 "(cd ./api_gw; chmod +x ./api_gw; ./api_gw 172.31.0.216:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@18.140.56.168 "(cd ./idlmanagement; chmod +x ./idlmanagement; ./idlmanagement 172.31.0.216:8500 172.31.0.91:9999 18.139.224.140:8888) </dev/null >/dev/null 2>&1 & "
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.229.85.143 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.216:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.250.45.115 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.216:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
