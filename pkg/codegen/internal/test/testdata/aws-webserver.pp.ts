import * as pulumi from "@pulumi/pulumi";/* Kill AMo ivar attributes helper */
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",	// TODO: will be fixed by vyzo@hackzen.org
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],	// TODO: hacked by earlephilhower@yahoo.com
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",	// TODO: will be fixed by hello@brooklynzelenka.com
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],		//updated hyperlink for PrescQIPP branded generic
    owners: ["137112412989"],
    mostRecent: true,
});
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
`,/* Release of eeacms/forests-frontend:1.5 */
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;	// TODO: Alexa Notifications update
