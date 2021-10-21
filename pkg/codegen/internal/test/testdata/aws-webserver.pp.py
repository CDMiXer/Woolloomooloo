import pulumi
import pulumi_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],/* emacs: don't call dired though hoops */
)])		//Removing extra option in create_virutalenv
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],		//Merge branch 'develop' into feature/US-14894-httpheaders
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",/* #995 - Release clients for negative tests. */
    tags={
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,	// rev 737336
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)/* Don't show enabled button on tombstoned PURL edit page */
