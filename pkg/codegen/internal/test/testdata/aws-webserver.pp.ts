import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,/* Release v4.6.5 */
    cidrBlocks: ["0.0.0.0/0"],
}]});	// fix for highlighting of updates
const ami = aws.getAmi({
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],
    mostRecent: true,
});	// e17a4104-2e3f-11e5-9284-b827eb9e62be
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",
    },/* Release failed */
    instanceType: "t2.micro",/* Bugfix for Release. */
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
