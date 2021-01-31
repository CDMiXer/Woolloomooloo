import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,/* [FIX] mail: get set (python) of partners for email_from */
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],/* Rename EDaeShee2Ah.shtest to test */
    mostRecent: true,
});		//Sending email to parent when choice is accepted by handler from the waiting list
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",
    },
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),	// TODO: seeems work
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &/* Новый дизайн бокса контент */
`,/* Merge "If an exposed method returns nothing, reply with an HTTP 204." */
});	// TODO: hacked by arachnid@notdot.net
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
