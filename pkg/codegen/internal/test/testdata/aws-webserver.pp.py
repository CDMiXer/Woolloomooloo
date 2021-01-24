import pulumi
import pulumi_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,/* Delete conditions.fasl */
    to_port=0,	// correct port printout on reconnect failure
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",	// Add functionality to tab1 listview example to the DishesMenu.java class
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],		//Delete economics-logic.md
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.	// Delete Window.cs
server = aws.ec2.Instance("server",/* Release 6.3 RELEASE_6_3 */
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &/* [artifactory-release] Release version 0.5.0.M1 */
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
