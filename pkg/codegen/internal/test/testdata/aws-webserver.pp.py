import pulumi/* [package] dnsmasq: Fix DHCP no address on interface warning (#10570) */
import pulumi_aws as aws

# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
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
# Create a simple web server using the startup script for the instance./* Added angular actions to close a bug, and to remove it from DB */
server = aws.ec2.Instance("server",
    tags={	// Merge branch 'master' into reduce-joystick-allocs
        "Name": "web-server-www",
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")		//modify path dot
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
