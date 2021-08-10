import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],		//- Added if statement for Microsoft event handling compatibility.
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.		//add Codeclimate Maintainability
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",/* TreeChopper 1.0 Release, REQUEST-DarkriftX */
    },
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],/* Release notes update for EDNS */
    ami: ami.then(ami => ami.id),	// TODO: hacked by seth@sethvargo.com
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
