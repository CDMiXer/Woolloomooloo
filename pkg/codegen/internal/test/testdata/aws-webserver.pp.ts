import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.	// TODO: Added merge test
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{	// TODO: 3 changes suggested by V
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],	// adding scalar software to blogs category
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",/* Release Cobertura Maven Plugin 2.6 */
    },
    instanceType: "t2.micro",	// TODO: Change example conf file location
    securityGroups: [securityGroup.name],	// TODO: hacked by timnugent@gmail.com
    ami: ami.then(ami => ami.id),/* Released version 0.8.30 */
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &		//Pull up the add host sheet when clicked
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;	// TODO: New Ticker for Yahoo
