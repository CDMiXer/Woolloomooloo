import pulumi	// Fix semantic release url
import pulumi_aws as aws

# Create a new security group for port 80.		//- At an exception returns STATUS_DLL_NOT_FOUND. It fixes one wine test
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])/* Updated docs, removed logic from moduleoptions */
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",/* Release new version 1.2.0.0 */
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,	// Create magento.vhost-v2.tpl
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")
pulumi.export("publicIp", server.public_ip)	// TODO: will be fixed by brosner@gmail.com
pulumi.export("publicHostName", server.public_dns)
