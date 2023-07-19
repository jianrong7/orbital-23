import json
import os

dirname = os.path.dirname(__file__)
filename = os.path.join(dirname, "../outputs.json")
f = open(filename)
data = json.load(f)

config_load = [
    ("api_gw", "../api_gw"),
    ("idl_management_service", "../service_definitions/idlmanagement/"),
    ("service1v1", "../service_definitions/service1v1/"),
]

def scp_line(ip_address, directory):
    return f"scp -i ..tfkey.pem -r {directory} ec2-user@{ip_address}"

for n in config_load:
    full_name = n[0] + "_public_ip"
    addr = data[full_name]["value"]
    if type(addr) is str:
        scp_line(addr, n[1])
        print(full_name + " " + addr)
    elif type(addr) is list:
        for ip in addr[0]:
            print(full_name + " " + ip)



# scp_line()