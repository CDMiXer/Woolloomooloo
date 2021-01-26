import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,	// TODO: will be fixed by mikeal.rogers@gmail.com
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",	// TODO: Update build-filmography.sh
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],		//Work in progress: New hub as a webservice in tomcat
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {/* Merge "Release green threads properly" */
        Name: "web-server-www",
    },/* Merge "Rename 'history' -> 'Release notes'" */
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &/* Flat LazyImage */
`,		//Added grunt-json dependency
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
