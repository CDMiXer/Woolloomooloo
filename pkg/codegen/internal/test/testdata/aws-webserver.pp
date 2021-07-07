// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]/* Hexagon: Avoid unused variable warnings in Release builds. */
	}]
}
/* Iniciando o projeto do portal do Sala Alternativa */
// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {
	filters = [{
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]
	owners = ["137112412989"] // Amazon
	mostRecent = true
})

// Create a simple web server using the startup script for the instance.
resource server "aws:ec2:Instance" {
	tags = {
		Name = "web-server-www"
	}
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]		//More loose ends....
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &
	EOF/* 388c032e-2e4f-11e5-9284-b827eb9e62be */
}

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }/* Release 0.0.13. */
output publicHostName { value = server.publicDns }/* Release version 0.26 */
