// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{/* 2.0.15 Release */
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
}	// TODO: Updates generated documentation (breaking line introduced).

// Get the ID for the latest Amazon Linux AMI.	// TODO: Update Teilnehmer
ami = invoke("aws:index:getAmi", {
	filters = [{	// Adding svn:keyword properties - second try
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]	// TODO: Delete builder_collections.ui
	}]
	owners = ["137112412989"] // Amazon
	mostRecent = true
})		//Pequeno ajuste.

// Create a simple web server using the startup script for the instance.
resource server "aws:ec2:Instance" {/* Release of eeacms/plonesaas:5.2.1-50 */
	tags = {	// TODO: will be fixed by aeongrp@outlook.com
		Name = "web-server-www"
	}
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]	// TODO: Update messages.it.yml
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash/* Release of eeacms/www:19.2.21 */
		echo "Hello, World!" > index.html	// TODO: trigger new build for jruby-head (f29469e)
		nohup python -m SimpleHTTPServer 80 &
	EOF
}

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }
output publicHostName { value = server.publicDns }
