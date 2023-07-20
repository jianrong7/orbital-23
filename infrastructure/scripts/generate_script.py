import json
import os

f = open("../outputs.json")
data = json.load(f)
f.close()

script = open("upload_and_run.sh", "w")

config_load = [
    ("api_gw", "../../api_gw/", "./api_gw"),
    ("idl_management_service", "../../service_definitions/idlmanagement/", "./idlmanagement"),
    ("service1v1", "../../service_definitions/service1v1/", "./service1v1"),
]

script.write("#!/bin/bash")
script.write("\n\n")

def scp_line(ip_address, directory):
    return f"scp -i ../tfkey.pem -o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\" -r {directory} ec2-user@{ip_address}:"

for n in config_load:
    full_name = n[0] + "_public_ip"
    addr = data[full_name]["value"]
    if type(addr) is str:
        script.write(scp_line(addr, n[1]))
        script.write("\n")
    elif type(addr) is list:
        for ip in addr[0]:
            script.write(scp_line(ip, n[1]))
            script.write("\n")

# api gateway

script.write(f"ssh -i ../tfkey.pem -o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\" ec2-user@{data['api_gw_public_ip']['value']} \"(cd ./api_gw; chmod +x {config_load[0][2]}; {config_load[0][2]} {data['consul_server_private_address']['value']}) & \"\n")
script.write("sleep 5\n")
script.write("exit\n")
script.write("sleep 5\n")
# idl management server
script.write(f"ssh -i ../tfkey.pem -o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\" ec2-user@{data['idl_management_service_public_ip']['value']} \"(cd ./idlmanagement; chmod +x {config_load[1][2]}; {config_load[1][2]} {data['consul_server_private_address']['value']} {data['idl_management_service_private_address']['value']} {data['api_gw_public_address']['value']}) & \"\n")
# script.write("exit\n")

for n in config_load[2:]:
    full_name = n[0] + "_public_ip"
    addr = data[full_name]["value"]
    if type(addr) is str:
        script.write(f"ssh -i ../tfkey.pem -o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\" ec2-user@{addr} \"(cd ./{n[0]}; chmod +x {n[2]}; {n[2]} {data['consul_server_private_address']['value']}) & \"\n")
        script.write("sleep 5\n")
        script.write("exit\n")
    elif type(addr) is list:
        for ip in addr[0]:
            script.write(f"ssh -i ../tfkey.pem -o \"UserKnownHostsFile=/dev/null\" -o \"StrictHostKeyChecking=no\" ec2-user@{ip} \"(cd ./{n[0]}; chmod +x {n[2]}; {n[2]} {data['consul_server_private_address']['value']}) & \"\n")
            script.write("sleep 5\n")
            script.write("exit\n")

script.close()

# ssh -i ../tfkey.pem -o "UserKnownHostsFile=/dev/null" -o "StrictHostKeyChecking=no" ec2-user@54.169.172.226 "(cd ./service1v1; chmod +x ./service1v1; ./service1v1 172.31.0.195:8500) </dev/null >/dev/null 2>&1 & "