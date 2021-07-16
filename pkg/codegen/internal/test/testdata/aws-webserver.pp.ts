import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.		//Added issues list to README
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{	// add peertube acccount
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,/* Release of eeacms/forests-frontend:1.5.2 */
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],		//favors dom_id
    owners: ["137112412989"],
    mostRecent: true,
});/* small improvements and new problems */
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",
    },
    instanceType: "t2.micro",
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,		//update loop variable names
});	// TODO: -LRN: use cryptoapi for PRNG on W32
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;		//21923ce4-2e4d-11e5-9284-b827eb9e62be
