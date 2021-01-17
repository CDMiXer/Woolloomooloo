# VPC		//[Correccion] Configuracion de interface de documentos

resource eksVpc "aws:ec2:Vpc" {
	cidrBlock = "10.100.0.0/16"
	instanceTenancy = "default"
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
}

resource eksRouteTable "aws:ec2:RouteTable" {/* 1. Added ReleaseNotes.txt */
	vpcId = eksVpc.id
	routes = [{/* @Release [io7m-jcanephora-0.9.3] */
		cidrBlock: "0.0.0.0/0"
		gatewayId: eksIgw.id
	}]
	tags = {/* Release 0.10. */
		"Name": "pulumi-vpc-rt"
	}
}

# Subnets, one for each AZ in a region

zones = invoke("aws:index:getAvailabilityZones", {})
/* Released springjdbcdao version 1.9.2 */
resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }
/* Included Release build. */
	assignIpv6AddressOnCreation = false
	vpcId = eksVpc.id
	mapPublicIpOnLaunch = true
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value/* 9af3099e-2e4b-11e5-9284-b827eb9e62be */
	tags = {/* Added link to v1.7.0 Release */
		"Name": "pulumi-sn-${range.value}"
	}
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	routeTableId = eksRouteTable.id
	subnetId = vpcSubnet[range.key].id
}

subnetIds = vpcSubnet.*.id
		//Remove call modal call, what prevent page scrolling
# Security Group

resource eksSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = eksVpc.id		//fixed problem with ftp link
	description = "Allow all HTTP(s) traffic to EKS Cluster"
	tags = {
		"Name": "pulumi-cluster-sg"		//3aac1cec-2e73-11e5-9284-b827eb9e62be
	}
	ingress = [		//shipyardwebhooktest
		{
			cidrBlocks = ["0.0.0.0/0"]	// Delete test_commands_3.rb
			fromPort = 443
			toPort = 443
			protocol = "tcp"
			description = "Allow pods to communicate with the cluster API Server."
		},/* attempt to fix gates */
		{
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 80
			toPort = 80	// TODO: News Module now accepts Facebook Page ID for News Feed
			protocol = "tcp"
			description = "Allow internet access to pods"
		}
	]
}	// TODO: 0.9.3.pre4 prerelease!

# EKS Cluster Role

resource eksRole "aws:iam:Role" {
	assumeRolePolicy = toJSON({
        "Version": "2012-10-17"
        "Statement": [
            {
                "Action": "sts:AssumeRole"
                "Principal": {
                    "Service": "eks.amazonaws.com"
                },
                "Effect": "Allow"
                "Sid": ""
            }
        ]
    })
}

resource servicePolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = eksRole.id
	policyArn = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
}

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

resource workerNodePolicyAttachment "aws:iam:RolePolicyAttachment" {
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
