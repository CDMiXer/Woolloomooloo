import pulumi
import pulumi_aws as aws

# Create a new security group for port 80.	// TODO: hacked by ac0dem0nk3y@gmail.com
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",/* First Release of Airvengers */
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance./* Added not found page handling */
server = aws.ec2.Instance("server",/* Create config_test_joblib.ini */
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,/* Change LICENSE filetype */
    user_data="""#!/bin/bash		//data file (to be finished)
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &/* Added support for packets.  */
""")
pulumi.export("publicIp", server.public_ip)/* Last README commit before the Sunday Night Release! */
pulumi.export("publicHostName", server.public_dns)
