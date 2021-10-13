import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";/* Release 2.7.1 */

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,/* 1.1.1 Release */
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{/* Add More Details to Release Branches Section */
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],	// TODO: Upgrade to peep 2.0.
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",
    },
    instanceType: "t2.micro",/* rev 876025 */
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
