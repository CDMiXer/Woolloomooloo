import pulumi
import pulumi_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,/* Release notes moved on top + link to the 0.1.0 branch */
    to_port=0,/* When rolling back, just set the Formation to the old Release's formation. */
    cidr_blocks=["0.0.0.0/0"],/* Technical 2 blog Draft */
)])		//Fix #925 (Bulk convert (2 files) fails with the errors below.)
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],/* Merge "Release 5.3.0 (RC3)" */
    most_recent=True)
# Create a simple web server using the startup script for the instance./* Release of eeacms/jenkins-slave-eea:3.21 */
server = aws.ec2.Instance("server",
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html/* Release 1.10.7 */
nohup python -m SimpleHTTPServer 80 &
""")		//Update hypothesis from 3.65.3 to 3.66.1
pulumi.export("publicIp", server.public_ip)	// TODO: dashboard style updates #227
pulumi.export("publicHostName", server.public_dns)
