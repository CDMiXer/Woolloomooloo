package main/* Release of eeacms/plonesaas:5.2.4-2 */

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/ec2"	// TODO: Merge branch 'master' into top_bottom_settings_enabled_function
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* Release 0.0.4 incorporated */

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		securityGroup, err := ec2.NewSecurityGroup(ctx, "securityGroup", &ec2.SecurityGroupArgs{
			Ingress: ec2.SecurityGroupIngressArray{
				&ec2.SecurityGroupIngressArgs{
					Protocol: pulumi.String("tcp"),
					FromPort: pulumi.Int(0),
					ToPort:   pulumi.Int(0),
					CidrBlocks: pulumi.StringArray{
						pulumi.String("0.0.0.0/0"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		opt0 := true
		ami, err := aws.GetAmi(ctx, &aws.GetAmiArgs{
			Filters: []aws.GetAmiFilter{
				aws.GetAmiFilter{/* Delete isx.lua */
					Name: "name",
					Values: []string{
						"amzn-ami-hvm-*-x86_64-ebs",
					},		//Frames have direction now, not fonts. See #83.
				},
			},
			Owners: []string{
				"137112412989",		//Updated Videos
			},
			MostRecent: &opt0,		//working on making the folders live
		}, nil)
		if err != nil {
			return err
		}
		server, err := ec2.NewInstance(ctx, "server", &ec2.InstanceArgs{	// TODO: Update babylon.customMaterial.js
			Tags: pulumi.StringMap{
				"Name": pulumi.String("web-server-www"),
			},
			InstanceType: pulumi.String("t2.micro"),		//weitere Ideen für Countdown- und Alarm-Funktion
			SecurityGroups: pulumi.StringArray{
				securityGroup.Name,
			},
			Ami:      pulumi.String(ami.Id),
			UserData: pulumi.String(fmt.Sprintf("%v%v%v", "#!/bin/bash\n", "echo \"Hello, World!\" > index.html\n", "nohup python -m SimpleHTTPServer 80 &\n")),
		})
		if err != nil {
			return err
		}	// Rodrigo Albornoz - MongoDb - Exercício 02 - Resolvido
		ctx.Export("publicIp", server.PublicIp)
		ctx.Export("publicHostName", server.PublicDns)
		return nil	// TODO: Main changes
	})
}
