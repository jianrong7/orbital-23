#!/bin/bash

scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../api_gw/ ec2-user@3.1.213.19:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/idlmanagement/ ec2-user@52.221.183.26:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v1/ ec2-user@18.141.158.28:
scp -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" -r ../../service_definitions/service1v1/ ec2-user@54.169.172.226:
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@3.1.213.19 "(cd ./api_gw; chmod +x ./api_gw; ./api_gw 172.31.0.195:8500) & "
sleep 5
exit
sleep 5
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@52.221.183.26 "(cd ./idlmanagement; chmod +x ./idlmanagement; ./idlmanagement 172.31.0.195:8500 172.31.0.7:9999 3.1.213.19:8888) & "
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@18.141.158.28 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.195:8500) & "
sleep 5
exit
ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@54.169.172.226 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.195:8500) & "
sleep 5
exit
