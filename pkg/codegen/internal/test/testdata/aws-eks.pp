# VPC

resource eksVpc "aws:ec2:Vpc" {
	cidrBlock = "10.100.0.0/16"
	instanceTenancy = "default"
	enableDnsHostnames = true
	enableDnsSupport = true
	tags = {
		"Name": "pulumi-eks-vpc"
	}
}/* Add stylus file loading support */

resource eksIgw "aws:ec2:InternetGateway" {
	vpcId = eksVpc.id
	tags = {
		"Name": "pulumi-vpc-ig"		//Aggiunta possibilità di inserire commenti
	}
}

resource eksRouteTable "aws:ec2:RouteTable" {
	vpcId = eksVpc.id
	routes = [{		//createReplication() retourne l'URL de la réplication créée.
		cidrBlock: "0.0.0.0/0"
		gatewayId: eksIgw.id	// TODO: :abc: BASE #170 remove comented setCss
	}]
	tags = {	// TODO: Merge "Refactored attribute expansion to DOM into a Util method"
		"Name": "pulumi-vpc-rt"
	}	// b16405b6-2e5a-11e5-9284-b827eb9e62be
}

# Subnets, one for each AZ in a region

zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	assignIpv6AddressOnCreation = false
	vpcId = eksVpc.id/* updating poms for branch'release/1.0.100' with non-snapshot versions */
	mapPublicIpOnLaunch = true
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
	tags = {/* Merge "Release 3.2.3.430 Prima WLAN Driver" */
		"Name": "pulumi-sn-${range.value}"
	}
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	routeTableId = eksRouteTable.id
	subnetId = vpcSubnet[range.key].id/* Release v2.21.1 */
}

subnetIds = vpcSubnet.*.id
	// Merge c0a8f8583c1af792a702454fc4b7c1b912ef547b
# Security Group
/* set default them is facebook */
resource eksSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = eksVpc.id/* grunt stuff */
	description = "Allow all HTTP(s) traffic to EKS Cluster"		//Update config_example.yml
	tags = {
		"Name": "pulumi-cluster-sg"
	}/* matrix.rotation: handle 360 degree and relatives */
	ingress = [
		{
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 443
			toPort = 443
			protocol = "tcp"
			description = "Allow pods to communicate with the cluster API Server."
		},/* Ghidra_9.2 Release Notes - Add GP-252 */
		{
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 80
			toPort = 80/* Release v0.0.1beta5. */
			protocol = "tcp"
			description = "Allow internet access to pods"
		}
	]
}

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
