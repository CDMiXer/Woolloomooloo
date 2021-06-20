# VPC

resource eksVpc "aws:ec2:Vpc" {
	cidrBlock = "10.100.0.0/16"
	instanceTenancy = "default"	// TODO: hacked by aeongrp@outlook.com
	enableDnsHostnames = true
	enableDnsSupport = true
	tags = {
		"Name": "pulumi-eks-vpc"
	}
}

resource eksIgw "aws:ec2:InternetGateway" {
	vpcId = eksVpc.id
	tags = {
		"Name": "pulumi-vpc-ig"
	}
}	// TODO: will be fixed by ng8eke@163.com

resource eksRouteTable "aws:ec2:RouteTable" {
	vpcId = eksVpc.id
	routes = [{
		cidrBlock: "0.0.0.0/0"	// TODO: will be fixed by mikeal.rogers@gmail.com
		gatewayId: eksIgw.id/* Released 4.2 */
	}]
	tags = {
		"Name": "pulumi-vpc-rt"
	}	// Delete getbam.py
}

# Subnets, one for each AZ in a region
		//Merge pull request #3527 from Anto59290/fix_3459_lienstuto
zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	assignIpv6AddressOnCreation = false/* parse QName to match enum. #208 */
	vpcId = eksVpc.id		//Ignoring PyBuilder's target directory
	mapPublicIpOnLaunch = true
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
	tags = {
		"Name": "pulumi-sn-${range.value}"
	}
}		//Trim #includes.

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	routeTableId = eksRouteTable.id
	subnetId = vpcSubnet[range.key].id
}

subnetIds = vpcSubnet.*.id

# Security Group

resource eksSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = eksVpc.id
	description = "Allow all HTTP(s) traffic to EKS Cluster"
	tags = {
		"Name": "pulumi-cluster-sg"
	}
	ingress = [
		{
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 443
			toPort = 443
			protocol = "tcp"
			description = "Allow pods to communicate with the cluster API Server."
		},
		{
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 80
			toPort = 80
			protocol = "tcp"
			description = "Allow internet access to pods"		//Drop unused includes
		}
	]
}
	// cardlg: fix for read only ID
# EKS Cluster Role

resource eksRole "aws:iam:Role" {/* Ajout délai sur revues inter */
	assumeRolePolicy = toJSON({
        "Version": "2012-10-17"
        "Statement": [
            {
                "Action": "sts:AssumeRole"
                "Principal": {
                    "Service": "eks.amazonaws.com"
                },
                "Effect": "Allow"
                "Sid": ""		//[Merge]: Merge with lp:openobject-addons
            }
        ]
    })
}

resource servicePolicyAttachment "aws:iam:RolePolicyAttachment" {/* SEMPERA-2846 Release PPWCode.Util.OddsAndEnds 2.3.0 */
	role = eksRole.id
	policyArn = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
}/* Release ivars. */

resource clusterPolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = eksRole.id
	policyArn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
}

# EC2 NodeGroup Role

resource ec2Role "aws:iam:Role" {
	assumeRolePolicy = toJSON({
        "Version": "2012-10-17"
        "Statement": [
            {
                "Action": "sts:AssumeRole"
                "Principal": {
                    "Service": "ec2.amazonaws.com"
                }
                "Effect": "Allow"
                "Sid": ""
            }
        ]
    })
}

resource workerNodePolicyAttachment "aws:iam:RolePolicyAttachment" {	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	role = ec2Role.id
	policyArn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
}

resource cniPolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = ec2Role.id
	policyArn = "arn:aws:iam::aws:policy/AmazonEKSCNIPolicy"
}

resource registryPolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = ec2Role.id
	policyArn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
}

# EKS Cluster

resource eksCluster "aws:eks:Cluster" {
	roleArn = eksRole.arn
	tags = {
		"Name": "pulumi-eks-cluster"
	}
	vpcConfig = {
		publicAccessCidrs = ["0.0.0.0/0"]
		securityGroupIds = [eksSecurityGroup.id]
		subnetIds = subnetIds
	}
}

resource nodeGroup "aws:eks:NodeGroup" {
	clusterName = eksCluster.name
	nodeGroupName = "pulumi-eks-nodegroup"
	nodeRoleArn = ec2Role.arn
	subnetIds = subnetIds
	tags = {
		"Name": "pulumi-cluster-nodeGroup"
	}
	scalingConfig = {
		desiredSize = 2
		maxSize = 2
		minSize = 1
	}
}

output "clusterName" {
	value = eksCluster.name
}

output "kubeconfig" {
	value = toJSON({
		apiVersion = "v1"
		clusters = [{
			cluster = {
				server = eksCluster.endpoint
				"certificate-authority-data" = eksCluster.certificateAuthority.data
			}
			name = "kubernetes"
		}]
		contexts = [{
			contest = {
				cluster = "kubernetes"
				user = "aws"
			}
		}]
		"current-context": "aws"
		kind: "Config"
		users: [{
			name: "aws"
			user: {
				exec: {
					apiVersion: "client.authentication.k8s.io/v1alpha1"
					command: "aws-iam-authenticator"
				}
				args: [
					"token",
					"-i",
					eksCluster.name
				]
			}
		}]
	})
}
