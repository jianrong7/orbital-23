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
  # AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY env variables must be set before running terraform plan / apply!
}

resource "aws_vpc" "orbital-23" {
  cidr_block = var.cidr_block
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

# can add more subnets to look cooler with the different network addresses :)
# or rather have the different servers (API Gateway, IDL Management, Services, Consul) reside in different subnets

resource "aws_key_pair" "tfkey" {
  key_name   = "tfkey"
  public_key = file("./tfkey.pub")
}

resource "aws_instance" "consul_server" {
  ami                    = var.image_id
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.consul_server_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  # install and run consul on the EC2 instance using the provisioners
  provisioner "file" {
    source      = "./scripts/consul_install.sh"
    destination = "/tmp/consul_install.sh"

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = self.public_ip
    }
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/consul_install.sh",
      "/tmp/consul_install.sh"
    ]

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = self.public_ip
    }
  }

  tags = {
    Name = "Consul_Server"
  }
}

resource "aws_instance" "api_gateway" {
  ami                    = var.image_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.api_gw_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  tags = {
    Name = "API_Gateway"
  }
}

resource "aws_instance" "idl_management" {
  ami                    = var.image_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  tags = {
    Name = "IDL_Management_Service"
  }
}

resource "aws_instance" "service1v1" {
  count                  = var.service1v1_count
  ami                    = var.image_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  tags = {
    Name = "service1v1-${count.index}"
  }
}

resource "aws_instance" "service1v2" {
  count                  = var.service1v2_count
  ami                    = var.image_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  tags = {
    Name = "service1v2-${count.index}"
  }
}

resource "aws_instance" "service2v1" {
  count                  = var.service2v1_count
  ami                    = var.image_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  tags = {
    Name = "service2v1-${count.index}"
  }
}

resource "aws_internet_gateway" "my_internet_gateway" {
  vpc_id = aws_vpc.orbital-23.id

  tags = {
    Name = "dev-igw"
  }
}

resource "aws_route_table" "my_public_rt" {
  vpc_id = aws_vpc.orbital-23.id

  tags = {
    Name = "dev_public_rt"
  }
}

resource "aws_route" "default_route" {
  route_table_id         = aws_route_table.my_public_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.my_internet_gateway.id
}

resource "aws_route_table_association" "my_public_assoc" {
  subnet_id      = aws_subnet.sg_a.id
  route_table_id = aws_route_table.my_public_rt.id
}

resource "aws_security_group" "vpc_sg" { # allow all tcp traffic within the VPC
  name        = "vpc_sg"
  description = "server tcp security group"
  vpc_id      = aws_vpc.orbital-23.id

  ingress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = [aws_vpc.orbital-23.cidr_block]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "ssh_sg" { # expose ssh port to all addresses
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

resource "aws_security_group" "consul_server_sg" { # expose 8500 port to all addresses
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

resource "aws_security_group" "api_gw_sg" { # expose 8888 port to all addresses
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
  type        = string
  description = "The IPv4 address of the VPC of the EC2 instances."
  default     = "172.31.0.0/16"
}

variable "image_id" {
  type        = string
  description = "The AMI id you want to launch the EC2 instance in."
  default     = "ami-0b1217c6bff20e276"
}

variable "instance_type" {
  type        = string
  description = "The type of EC2 instance you want to launch."
  default     = "t2.micro"
}

variable "service1v1_count" {
  type        = number
  description = "The number of identical EC2 instances to be deployed for service1v1."
  default     = 1
}

variable "service1v2_count" {
  type        = number
  description = "The number of identical EC2 instances to be deployed for service1v2."
  default     = 1
}

variable "service2v1_count" {
  type        = number
  description = "The number of identical EC2 instances to be deployed for service2v1."
  default     = 1
}