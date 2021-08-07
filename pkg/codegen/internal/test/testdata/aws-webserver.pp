// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{
		protocol = "tcp"/* Release version 2.2.4 */
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
}	// TODO: Merge branch 'master' into feature/new-group-conversation

// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {
	filters = [{
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]		//Unit test of DatabaseConfiguration
	owners = ["137112412989"] // Amazon
	mostRecent = true		//ADD: detect an invalid context and restart with a fresh context
})

// Create a simple web server using the startup script for the instance.
resource server "aws:ec2:Instance" {
	tags = {
		Name = "web-server-www"
	}
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]/* Release for 18.23.0 */
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash/* [releng] Release v6.16.2 */
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &
	EOF
}

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }
output publicHostName { value = server.publicDns }
