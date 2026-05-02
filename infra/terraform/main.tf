provider "aws" { region = "eu-west-2" }

resource "aws_vpc" "civicore" { cidr_block = "10.80.0.0/16" }

# CIV-IAC-004: ingress 5432 from anywhere.
resource "aws_security_group" "db" {
  name   = "civicore-db"
  vpc_id = aws_vpc.civicore.id
  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
