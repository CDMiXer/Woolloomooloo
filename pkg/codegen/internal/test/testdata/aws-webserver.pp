// Create a new security group for port 80.
resource securityGroup "aws:ec2:SecurityGroup" {	// Libellés pour le service obsolescence
	ingress = [{
		protocol = "tcp"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]/* Added "-" as a range operator for numbers */
	}]
}

// Get the ID for the latest Amazon Linux AMI.
ami = invoke("aws:index:getAmi", {	// TODO: will be fixed by mikeal.rogers@gmail.com
	filters = [{
		name = "name"
		values = ["amzn-ami-hvm-*-x86_64-ebs"]
	}]
	owners = ["137112412989"] // Amazon
	mostRecent = true		//changed log level on tests
})	// TODO: hacked by ac0dem0nk3y@gmail.com
/* 171d4da0-2e46-11e5-9284-b827eb9e62be */
// Create a simple web server using the startup script for the instance.	// Changed snapping key to 'd'
resource server "aws:ec2:Instance" {
	tags = {	// TODO: hacked by arajasek94@gmail.com
		Name = "web-server-www"
	}
	instanceType = "t2.micro"
	securityGroups = [securityGroup.name]
	ami = ami.id
	userData = <<-EOF
		#!/bin/bash
		echo "Hello, World!" > index.html/* Merge "Move some of the json event processing code to a common file" */
		nohup python -m SimpleHTTPServer 80 &
	EOF
}
	// Merge "[INTERNAL] sap.ui.demo.mdskeleton - fixing the not found page"
// Export the resulting server's IP address and DNS name.		//Joomla core update to 3.6.2
output publicIp { value = server.publicIp }/* Removed keystore */
output publicHostName { value = server.publicDns }		//Added an example test for factorization of polynomials
