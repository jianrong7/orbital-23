#!/bin/bash

scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../api_gw/ ec2-user@54.255.81.16:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/idlmanagement/ ec2-user@18.139.219.73:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v1/ ec2-user@18.136.203.250:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v1/ ec2-user@13.215.163.243:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v2/ ec2-user@54.169.202.40:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v2/ ec2-user@18.136.199.28:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service2v1/ ec2-user@13.229.250.43:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service2v1/ ec2-user@13.229.97.119:
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@54.255.81.16 "(cd ./api_gw; chmod +x ./api_gw; ./api_gw 172.31.0.26:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@18.139.219.73 "(cd ./idlmanagement; chmod +x ./idlmanagement; ./idlmanagement 172.31.0.26:8500 172.31.0.71:9999 54.255.81.16:8888) </dev/null >/dev/null 2>&1 & "
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@18.136.203.250 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.26:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.215.163.243 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.26:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@54.169.202.40 "(cd ./service1v2; chmod +x ./service1v2; ./service1v2 172.31.0.26:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@18.136.199.28 "(cd ./service1v2; chmod +x ./service1v2; ./service1v2 172.31.0.26:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.229.250.43 "(cd ./service2v1; chmod +x ./service2v1; ./service2v1 172.31.0.26:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@13.229.97.119 "(cd ./service2v1; chmod +x ./service2v1; ./service2v1 172.31.0.26:8500) </dev/null >/dev/null 2>&1 & "
sleep 5
