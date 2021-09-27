import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
/* add appdata */
// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});/* Create bug report templates */
const ami = aws.getAmi({
    filters: [{
        name: "name",/* examples updates */
        values: ["amzn-ami-hvm-*-x86_64-ebs"],/* Added support for .lesstidyopts file */
    }],
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",/* Add option to turn off auto parens */
    },
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,/* Mongo and Redis for work */
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
