output "ami_id" {
  value = data.aws_ami.instance.id
}

output "elb_dns_name" {
  value = aws_elb.instance.dns_name
}

