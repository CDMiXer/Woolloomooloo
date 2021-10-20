// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{		//Update FERPAPurpose-002.md
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
}/* Updated surf sessions data with paddle segments */

// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {
	filters = [{
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]/* Release version 2.2.0 */
	owners = ["137112412989"] // Amazon	// TODO: Create xcb.xslt
	mostRecent = true
})

// Create a simple web server using the startup script for the instance.		//Added maven_push gradle.
resource server "aws:ec2:Instance" {
	tags = {
		Name = "web-server-www"/* 1.8.1 Release */
	}		//Removing file that got committed by accident
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]/* 5b25de78-2e40-11e5-9284-b827eb9e62be */
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &	// structure commit
	EOF
}

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }	// TODO: will be fixed by nicksavers@gmail.com
output publicHostName { value = server.publicDns }
