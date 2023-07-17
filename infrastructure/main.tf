terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = "ap-southeast-1"
}

resource "aws_vpc" "orbital-23" {
  cidr_block    = "172.31.0.0/16"
}

resource "aws_instance" "api_gateway" {
  ami           = "ami-0b1217c6bff20e276"
  instance_type = "t2.micro"

  user_data = <<-EOL
  #!/bin/bash -xe

  yum update -y
  yum install git -y
  
  

  EOL

  tags = {
    Name = "API_Gateway"
  }
}

resource "aws_instance" "consul_server" {
  ami           = aws_ami_from_instance.example.id
  instance_type = "t2.micro"
  key_name = aws_key_pair.aws_ec2.id
  vpc_security_group_ids = [aws_security_group.my_sg.id]
  subnet_id              = aws_subnet.my_public_subnet.id

  user_data = <<-EOL
  #!/bin/bash -xe

  hash foo 2>/dev/null || { echo >&2 "I require foo but it's not installed.  Aborting."; exit 1; }

  sudo yum install -y yum-utils shadow-utils
  sudo yum-config-manager --add-repo https://rpm.releases.hashicorp.com/AmazonLinux/hashicorp.repo
  sudo yum -y install consul
  consul -v
  consul agent -dev -client="0.0.0.0"
  EOL

  tags = {
    Name = "Consul_Server"
  }
}

# NOTHING WORKS HERE.
# resource "aws_ami_from_instance" "example" {
#   name               = "terraform-example"
#   source_instance_id = "i-0b19474a09b5b1d11"
# }

# resource "aws_key_pair" "aws_ec2" {
#   key_name = "aws-ec2"
#   public_key = file("~/.ssh/id_ed25519.pub")
# }

# resource "aws_vpc" "my_vpc" {
#   cidr_block           = "10.123.0.0/16"
#   enable_dns_hostnames = true
#   enable_dns_support   = true

#   tags = {
#     Name = "ec2_vpc"
#   }
# }

# resource "aws_subnet" "my_public_subnet" {
#   vpc_id                  = aws_vpc.my_vpc.id
#   cidr_block              = "10.123.1.0/24"
#   map_public_ip_on_launch = true
#   availability_zone       = "ap-southeast-1"

#   tags = {
#     Name = "dev-public"
#   }
# }

# resource "aws_internet_gateway" "my_internet_gateway" {
#   vpc_id = aws_vpc.my_vpc.id

#   tags = {
#     Name = "dev-igw"
#   }
# }

# resource "aws_route_table" "my_public_rt" {
#   vpc_id = aws_vpc.my_vpc.id

#   tags = {
#     Name = "dev_public_rt"
#   }
# }

# resource "aws_route" "default_route" {
#   route_table_id         = aws_route_table.my_public_rt.id
#   destination_cidr_block = "0.0.0.0/0"
#   gateway_id             = aws_internet_gateway.my_internet_gateway.id
# }

# resource "aws_route_table_association" "my_public_assoc" {
#   subnet_id      = aws_subnet.my_public_subnet.id
#   route_table_id = aws_route_table.my_public_rt.id
# }


# resource "aws_security_group" "my_sg" {
#   name        = "dev_sg"
#   description = "dev security group"
#   vpc_id      = aws_vpc.my_vpc.id

#   ingress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"]
#   }
#   egress {
#     from_port   = 0
#     to_port     = 0
#     protocol    = "-1"
#     cidr_blocks = ["0.0.0.0/0"] 
#   }
# }

