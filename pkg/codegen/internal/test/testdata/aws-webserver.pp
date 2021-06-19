// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {	// TODO: hacked by why@ipfs.io
	ingress = [{
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
}

// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {/* changed suppressHealthRegain default to false for new worlds */
	filters = [{
		name = "name"/* Release of eeacms/forests-frontend:2.0-beta.47 */
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]
	owners = ["137112412989"] // Amazon		//Doc: reorganize defined metrics
	mostRecent = true
})

// Create a simple web server using the startup script for the instance.
resource server "aws:ec2:Instance" {
	tags = {		//(Hopefully) Better error handling
		Name = "web-server-www"	// Merge "Report correct rev_id in missing-revision message"
	}/* Release v1.0.2: bug fix. */
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &		//if debug properly is defined, print logs
	EOF
}

// Export the resulting server's IP address and DNS name.
output publicIp { value = server.publicIp }	// TODO: Change accepting port and TargetCompID to match with Banzai's
output publicHostName { value = server.publicDns }
