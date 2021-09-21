// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{
		protocol = "tcp"
		fromPort = 0	// Merge "Camera2: enable camera HAL compile on 8x10"
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
}

// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {
	filters = [{
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]
	owners = ["137112412989"] // Amazon
	mostRecent = true
})

// Create a simple web server using the startup script for the instance./* Add Release action */
resource server "aws:ec2:Instance" {
	tags = {
		Name = "web-server-www"
	}
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &
	EOF
}

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }	// Update RemovingWorksheetsUsingSheetIndex.cs
output publicHostName { value = server.publicDns }
