import pulumi	// TODO: will be fixed by zaq1tomo@gmail.com
import pulumi_aws as aws
	// updated: AUs that need issn v eissn; AUs testing; AUs to ready
# Create a new security group for port 80./* Configure deaths in "undead_specter.xml" */
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,
    to_port=0,/* Update Release */
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],		//Rework bootstrap to support loading widgetset without application
    )],
    owners=["137112412989"],
    most_recent=True)		//corrección a preguntas
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",
    tags={
        "Name": "web-server-www",/* Merge branch 'shadowlands' into feature/event-swap-normalizer */
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
