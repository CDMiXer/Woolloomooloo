import * as pulumi from "@pulumi/pulumi";/* Release for 2.20.0 */
import * as aws from "@pulumi/aws";

// Create a new security group for port 80./* Update PayrollReleaseNotes.md */
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{/* Support Promise cancellation for SecureConnector */
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],	// Fix linter error in Kick
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],/* Change user class name and debug install mode */
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {	// Slight update for adding tooltips and direct ResourceBundle access
        Name: "web-server-www",
    },
    instanceType: "t2.micro",/* Forgot NDEBUG in the Release config. */
    securityGroups: [securityGroup.name],		//Correct lock wording
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
,`
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
