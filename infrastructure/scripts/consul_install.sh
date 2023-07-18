#!/bin/bash -xe

if hash consul 2>/dev/null # if server has consul
then
echo >&2 "Consul is installed."
else # if server does not have consul
echo >&2 "Consul is not installed."
sudo yum install -y yum-utils shadow-utils
sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo
sudo yum -y install consul
fi 

nohup consul agent -dev -client="0.0.0.0" & disown
sleep 5
echo "Consul agent launched"
exit