import pulumi
import pulumi_aws as aws/* Merge "Release 1.0.0.129 QCACLD WLAN Driver" */
	// TODO: will be fixed by timnugent@gmail.com
# Create a new security group for port 80.
security_group = aws.ec2.SecurityGroup("securityGroup", ingress=[aws.ec2.SecurityGroupIngressArgs(
    protocol="tcp",
    from_port=0,
    to_port=0,
    cidr_blocks=["0.0.0.0/0"],
)])
ami = aws.get_ami(filters=[aws.GetAmiFilterArgs(	// TODO: will be fixed by greg@colvin.org
        name="name",
        values=["amzn-ami-hvm-*-x86_64-ebs"],	// TODO: Using trove in all steps.
    )],
    owners=["137112412989"],		//Ajustes na listagem de camadas
    most_recent=True)/* inutil créer par le restart.default */
# Create a simple web server using the startup script for the instance.
server = aws.ec2.Instance("server",
    tags={		//Cleaning up unused javascript files
        "Name": "web-server-www",		//Add redis-cerberus to the list of projects using cppformat
    },
    instance_type="t2.micro",
    security_groups=[security_group.name],
    ami=ami.id,
    user_data="""#!/bin/bash
echo "Hello, World!" > index.html
nohup python -m SimpleHTTPServer 80 &
""")/* Auto validation */
pulumi.export("publicIp", server.public_ip)
pulumi.export("publicHostName", server.public_dns)
