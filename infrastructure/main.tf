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
  # AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY env variables must be set!
}

resource "aws_vpc" "orbital-23" {
  cidr_block    = var.cidr_block
}

resource "aws_subnet" "sg_a" {
  vpc_id                  = aws_vpc.orbital-23.id
  cidr_block              = "172.31.0.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-southeast-1a"

  tags = {
    Name = "sg_a"
  }
}

# can add more subnets to look cooler with the different network addresses

# resource "aws_key_pair" "orbital-23" {
#   key_name = "orbital-23"
#   public_key = var.orbital_public_key
# }

resource "aws_instance" "api_gateway" {
  ami           = var.image_id
  instance_type = var.instance_type
  key_name      = "orbital-23"
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.api_gw_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  user_data = <<-EOL
  #!/bin/bash -xe

  yum update -y
  
  EOL

  tags = {
    Name = "API_Gateway"
  }
}

resource "aws_instance" "consul_server" {
  ami           = var.image_id
  instance_type = "t2.micro"
  key_name      = "orbital-23"
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.consul_server_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  user_data = <<-EOL
  #!/bin/bash -xe

  if hash consul 2>/dev/null # if server has consul
  then
    echo >&2 "Consul is installed."
  else # if server does not have consul
    echo >&2 "Consul is not installed."
    curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
    sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
    sudo apt-get update && sudo apt-get install consul
  fi 
  consul -v
  consul agent -dev -client="0.0.0.0" &
  EOL

  tags = {
    Name = "Consul_Server"
  }
}

resource "aws_instance" "idl_management" {
  ami           = var.image_id
  instance_type = var.instance_type
  key_name      = "orbital-23"
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  user_data = <<-EOL
  #!/bin/bash -xe

  yum update -y
  
  EOL

  tags = {
    Name = "IDL_Management_Service"
  }
}

resource "aws_instance" "service1v1" {
  ami           = var.image_id
  instance_type = var.instance_type
  key_name      = "orbital-23"
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  user_data = <<-EOL
  #!/bin/bash -xe

  yum update -y

  EOL

  tags = {
    Name = "service1v1"
  }
}

resource "aws_security_group" "vpc_sg" {
  name        = "vpc_sg"
  description = "server tcp security group"
  vpc_id      = aws_vpc.orbital-23.id

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.orbital-23.cidr_block, var.my_ip]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] 
  }
}

resource "aws_security_group" "ssh_sg" {
  name        = "ssh_sg"
  description = "server ssh security group"
  vpc_id      = aws_vpc.orbital-23.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] 
  }
}

resource "aws_security_group" "consul_server_sg" {
  name        = "consul_server_sg"
  description = "consul server exposed port 8500 security group"
  vpc_id      = aws_vpc.orbital-23.id

  ingress {
    from_port   = 8500
    to_port     = 8500
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] 
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] 
  }
}

resource "aws_security_group" "api_gw_sg" {
  name        = "api_gw_sg"
  description = "api gateway exposed port 8888 security group"
  vpc_id      = aws_vpc.orbital-23.id

  ingress {
    from_port   = 8888
    to_port     = 8888
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] 
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"] 
  }
}

variable "cidr_block" {
  type = string
  description = "The IPv4 address of the VPC of the EC2 instances."
}

variable "image_id" {
  type = string
  description = "The AMI id you want to launch the EC2 instance in."
}

variable "instance_type" {
  type = string
  description = "The type of EC2 instance you want to launch."
}

variable "key_pair_id" {
  type = string
  description = "The id of the key_pair used for the EC2 deployment."
}

variable "orbital_public_key" {
  type = string
  description = "The public key used for the EC2 deployment."
}

variable "my_ip" {
  type = string
  description = "My Public IP address."
}
