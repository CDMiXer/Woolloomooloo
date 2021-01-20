import * as pulumi from "@pulumi/pulumi";	// TODO: fe3954d4-2e57-11e5-9284-b827eb9e62be
import * as aws from "@pulumi/aws";

// Create a new security group for port 80.
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{	// TODO: hamyar GP v1.3
    protocol: "tcp",
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],/* Initial websocket handler */
}]});
const ami = aws.getAmi({
    filters: [{	// Merge "Fix styles of table views"
        name: "name",/* FIX force allow prefill if using prefill_with_data_from_widget_link */
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {/* [JENKINS-60740] - Update Release Drafter to the recent version */
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
export const publicHostName = server.publicDns;
