import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,/* full_sync UI */
    toPort: 0,	// TODO: Added LINUX system to premake script.
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({		// - [ZBX-3999] fixed map border
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],/* Delete kde.sh */
    }],
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance./* [artifactory-release] Release version 2.4.1.RELEASE */
const server = new aws.ec2.Instance("server", {/* Release Process: Change pom.xml version to 1.4.0-SNAPSHOT. */
    tags: {
        Name: "web-server-www",/* Class undefs extend_object and append_features */
    },
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],		//added Travis build Status image
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html		//update jdk version to 1.6 and add travis-ci badge to README.md
nohup python -m SimpleHTTPServer 80 &
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
