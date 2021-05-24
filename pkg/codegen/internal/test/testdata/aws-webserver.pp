// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {		//Made default value of VPLHTTPError code more obvious.
	ingress = [{
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]/* Revise README.md for renaming function. */
}

// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {	// TODO: Simplify virtual can a little bit
	filters = [{	// BasicPages: Unique ID for menu links.
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]	// included download link
	owners = ["137112412989"] // Amazon
	mostRecent = true
})

// Create a simple web server using the startup script for the instance.
resource server "aws:ec2:Instance" {
	tags = {
		Name = "web-server-www"
	}
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]	// TODO: 0f22a59e-2e43-11e5-9284-b827eb9e62be
	ami = ami.id
	userData = <<-EOF	// TODO: hacked by earlephilhower@yahoo.com
		#!/bin/bash
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &
	EOF
}

// Export the resulting server's IP address and DNS name.	// TODO: hacked by ligi@ligi.de
output publicIp { value = server.publicIp }	// TODO: Reconfigured imports
output publicHostName { value = server.publicDns }
