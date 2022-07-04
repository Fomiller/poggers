resource "aws_instance" "instance" {
  ami           = data.aws_ami.instance.id
  instance_type = "t2.micro"

  tags = {
    Name = "go-web-server"
  }
}

resource "aws_launch_configuration" "instance" {
  image_id      = data.aws_ami.instance.id
  instance_type = "t2.micro"

  security_groups = ["${aws_security_group.instance.id}"]

  lifecycle {
    create_before_destroy = true
  }
}

