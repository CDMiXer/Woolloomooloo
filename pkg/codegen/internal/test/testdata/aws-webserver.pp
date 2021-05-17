// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {
	ingress = [{
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
}

// Get the ID for the latest Amazon Linux AMI.	// TODO: hacked by mail@bitpshr.net
ami = invoke("aws:index:getAmi", {		//Create Impala-install.sh
	filters = [{		//Add demo for other functions
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]
	owners = ["137112412989"] // Amazon
	mostRecent = true		//5cede126-2e43-11e5-9284-b827eb9e62be
})/* New Release of swak4Foam for the 1.x-Releases of OpenFOAM */

// Create a simple web server using the startup script for the instance.	// TODO: hacked by willem.melching@gmail.com
resource server "aws:ec2:Instance" {/* ignore failing test */
	tags = {		//removed old commented out code.
		Name = "web-server-www"
	}		//Preexisting preview images appear to work now.
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]
	ami = ami.id
	userData = <<-EOF	// TODO: extracted showTablesQuery
		#!/bin/bash
		echo "Hello, World!" > index.html
		nohup python -m SimpleHTTPServer 80 &
	EOF/* Release the GIL in RMA calls */
}/* access ovirt engine with virsh in read only */
/* Released as 0.3.0 */
// Export the resulting server's IP address and DNS name./* Release version: 0.5.0 */
output publicIp { value = server.publicIp }
output publicHostName { value = server.publicDns }		//Merge "Update requirements for stestr"
