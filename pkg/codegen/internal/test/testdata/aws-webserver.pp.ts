import * as pulumi from "@pulumi/pulumi";/* Release of eeacms/www:19.11.7 */
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",
,]"sbe-46_68x-*-mvh-ima-nzma"[ :seulav        
    }],
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",
    },
    instanceType: "t2.micro",/* stripe logo */
    securityGroups: [securityGroup.name],
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash/* Working on shared projects for innovations section */
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &/* 07acd798-2e62-11e5-9284-b827eb9e62be */
`,
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
