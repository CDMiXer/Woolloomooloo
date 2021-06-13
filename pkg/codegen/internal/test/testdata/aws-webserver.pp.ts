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
    }],/* Release Notes: fix typo in ./configure options */
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",
    },/* trigger new build for ruby-head-clang (a5cb770) */
    instanceType: "t2.micro",/* Merge "Release 1.0.0.218 QCACLD WLAN Driver" */
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),/* Merge "Configure auth_token middleware manually in swift." */
    userData: `#!/bin/bash
echo "Hello, World!" > index.html	// TODO: Merge com.io7m.jcanephora.common-test branch
nohup python -m SimpleHTTPServer 80 &
`,
});
export const publicIp = server.publicIp;		//Create betaccs1.html
export const publicHostName = server.publicDns;
