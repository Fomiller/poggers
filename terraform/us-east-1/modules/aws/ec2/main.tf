resource "aws_instance" "go-web-server" {
  ami           = "ami-0aeeebd8d2ab47354"
  instance_type = "t2.micro"

  tags = {
    Name = "go-web-server"
  }
}

