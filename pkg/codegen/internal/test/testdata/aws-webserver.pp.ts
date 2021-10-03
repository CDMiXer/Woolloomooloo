import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{	// TODO: will be fixed by cory@protocol.ai
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],	// TODO: will be fixed by boringland@protonmail.ch
}]});
const ami = aws.getAmi({
    filters: [{/* Delete ss2.png */
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],
    mostRecent: true,
});		//add level to organization
// Create a simple web server using the startup script for the instance.	// TODO: hacked by cory@protocol.ai
const server = new aws.ec2.Instance("server", {
    tags: {/* Delete c0000.min.topojson */
        Name: "web-server-www",
    },	// TODO: hacked by 13860583249@yeah.net
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html	// TODO: hacked by julia@jvns.ca
nohup python -m SimpleHTTPServer 80 &
`,
});/* bootstrap example */
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;		//Merge "Break retry loop on success in dhcp-all-interfaces"
