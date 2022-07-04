resource "aws_autoscaling_group" "instance" {
  launch_configuration = aws_launch_configuration.instance.id
  load_balancers       = ["${aws_elb.instance.name}"]
  availability_zones   = ["eu-west-1b", "eu-west-1a"]
  min_size             = 2
  max_size             = 5

  tag {
    key                 = "Name"
    value               = "terraform-go-api"
    propagate_at_launch = true
  }
}

