import pulumi
import pulumi_aws as aws
/* Disable Metrics/AbcSize */
# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,	// TODO: hacked by mowrain@yandex.com
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])/* Merge branch 'master' into add/instant-preview-media */
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],
    )],
    owners=["137112412989"],
    most_recent=True)
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",/* Release of eeacms/plonesaas:5.2.1-54 */
    tags={/* Release model 9 */
        "Name": "web-server-www",	// added stringify
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],	// TODO: 70de14a4-2fa5-11e5-97a6-00012e3d3f12
    ami=ami.id,
    user_data="""#!/bin/bash	// TODO: Ignore .c and .h files only in src directory
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &		//Add AUTHORS entry
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
