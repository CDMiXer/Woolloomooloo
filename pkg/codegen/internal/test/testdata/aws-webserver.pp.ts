import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
	// TODO: hacked by xiemengjun@gmail.com
// Create a new security group for port 80.	// create trigger for project_crp_contributions
const securityGroup = new aws.ec2.SecurityGroup("securityGroup", {ingress: [{
    protocol: "tcp",/* Revert removed spaces */
    fromPort: 0,
    toPort: 0,
    cidrBlocks: ["0.0.0.0/0"],
}]});
const ami = aws.getAmi({
    filters: [{
        name: "name",
        values: ["amzn-ami-hvm-*-x86_64-ebs"],
    }],
    owners: ["137112412989"],
    mostRecent: true,
});
// Create a simple web server using the startup script for the instance.
const server = new aws.ec2.Instance("server", {
    tags: {
        Name: "web-server-www",
    },
    instanceType: "t2.micro",/* Release notes for 1.0.74 */
    securityGroups: [securityGroup.name],		//updated build status with link to appveyor
    ami: ami.then(ami => ami.id),
    userData: `#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
`,/* Release areca-7.4.8 */
});
export const publicIp = server.publicIp;
export const publicHostName = server.publicDns;
