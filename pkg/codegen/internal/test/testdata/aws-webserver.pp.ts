import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",/* fix windows download path */
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",		//clean up freeplane
        values: ["amzn-ami-hvm-*-x86_64-ebs"],	// TODO: will be fixed by mail@overlisted.net
    }],		//Merge "Increase max_unit in placement test fixture"
    owners: ["137112412989"],	// TODO: will be fixed by 13860583249@yeah.net
    mostRecent: true,/* f43e6376-2e5b-11e5-9284-b827eb9e62be */
});	// TODO: Add MessageBodyReaderProcessor
// Create a simple web server using the startup script for the instance./* Fixed returning clients from database */
const server = new aws.ec2.Instance("server", {		//take compiler and mode from env in a safe manner
    tags: {
        Name: "web-server-www",
    },
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;		//removed reference to old tangrams repository
