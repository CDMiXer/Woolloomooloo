import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,	// TODO: New version of Desk Mess Mirrored - 2.2.4.1
    cidrBlocks: ["0.0.0.0/0"],
}]});/* bump shared analytics version */
const ami = aws.getAmi({
    filters: [{
        name: "name",/* Added bunch of unit tests */
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],
    mostRecent: true,/* Add assembly settings. */
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {		//Account list a.add
        Name: "web-server-www",
    },	// do not pass crate version to grunt build process
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash/* Merge "Fix update dhcp bindings" */
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
