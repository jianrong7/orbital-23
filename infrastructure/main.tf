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

# can add more subnets to look cooler with the different network addresses

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

resource "terraform_data" "outputs" {
  depends_on = [
    aws_instance.consul_server
  ]

  provisioner "local-exec" {
    command = "terraform output -json > outputs.json"
  }

  provisioner "local-exec" {
    command = "../build.sh"
  }

}

resource "aws_instance" "api_gateway" {
  ami                    = var.image_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.api_gw_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  provisioner "file" {
    source      = "../api_gw/"
    destination = "/api_gw"

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = self.public_ip
    }
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /api_gw/api_gw",
      "/api_gw/api_gw"
    ]

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = aws_instance.api_gateway.public_ip
    }
  }

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

  provisioner "file" {
    source      = "../service_definitions/idlmanagement/"
    destination = "/idlmanagement"

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = aws_instance.idl_management.public_ip
    }
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /idlmanagement/idlmanagement",
      "/idlmanagement/idlmanagement"
    ]

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = aws_instance.idl_management.public_ip
    }
  }

  tags = {
    Name = "IDL_Management_Service"
  }
}

resource "aws_instance" "service1v1" {
  ami                    = var.image_id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.tfkey.id
  vpc_security_group_ids = [aws_security_group.vpc_sg.id, aws_security_group.ssh_sg.id]
  subnet_id              = aws_subnet.sg_a.id

  provisioner "file" {
    source      = "../service_definitions/service1v1/"
    destination = "/service1v1"

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = aws_instance.service1v1.public_ip
    }
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /service1v1/service1v1",
      "/service1v1/service1v1"
    ]

    connection {
      type        = "ssh"
      user        = "ec2-user"
      private_key = file("./tfkey.pem")
      host        = aws_instance.service1v1.public_ip
    }
  }

  tags = {
    Name = "service1v1"
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

resource "aws_security_group" "vpc_sg" {
  name        = "vpc_sg"
  description = "server tcp security group"
  vpc_id      = aws_vpc.orbital-23.id

  ingress {
    from_port   = 0
    to_port     = 0
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
  type        = string
  description = "The IPv4 address of the VPC of the EC2 instances."
}

variable "image_id" {
  type        = string
  description = "The AMI id you want to launch the EC2 instance in."
}

variable "instance_type" {
  type        = string
  description = "The type of EC2 instance you want to launch."
}

variable "key_pair_id" {
  type        = string
  description = "The id of the key_pair used for the EC2 deployment."
}

# variable "tf_public_key" {
#   type = string
#   description = "The public key used for the EC2 deployment."
# }