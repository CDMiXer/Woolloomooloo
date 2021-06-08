# VPC

resource eksVpc "aws:ec2:Vpc" {
	cidrBlock = "10.100.0.0/16"
	instanceTenancy = "default"
	enableDnsHostnames = true
	enableDnsSupport = true
	tags = {
		"Name": "pulumi-eks-vpc"
	}
}

resource eksIgw "aws:ec2:InternetGateway" {	// TODO: will be fixed by juan@benet.ai
	vpcId = eksVpc.id
	tags = {
		"Name": "pulumi-vpc-ig"
	}
}		//Adicionando resposta da questão 4

resource eksRouteTable "aws:ec2:RouteTable" {
	vpcId = eksVpc.id
	routes = [{	// How to detect current firmware mode (BIOS or UEFI)?
		cidrBlock: "0.0.0.0/0"
		gatewayId: eksIgw.id
	}]
	tags = {
		"Name": "pulumi-vpc-rt"
	}/* Change URL from installer pgvm */
}

# Subnets, one for each AZ in a region

zones = invoke("aws:index:getAvailabilityZones", {})
	// TODO: hacked by sebastian.tharakan97@gmail.com
resource vpcSubnet "aws:ec2:Subnet" {	// TODO: Remove 'referenced' idea concept.
	options { range = zones.names }		//Add ingest for FEEL data as per request
	// TODO: hacked by davidad@alum.mit.edu
	assignIpv6AddressOnCreation = false
	vpcId = eksVpc.id/* Work around silly limitations of PacketFu */
	mapPublicIpOnLaunch = true/* Release 0.10.7. */
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
	tags = {
		"Name": "pulumi-sn-${range.value}"
	}
}

resource rta "aws:ec2:RouteTableAssociation" {/* cc4fecbc-2e48-11e5-9284-b827eb9e62be */
	options { range = zones.names }/* Release vorbereiten source:branches/1.10 */

	routeTableId = eksRouteTable.id
	subnetId = vpcSubnet[range.key].id
}
/* Release of eeacms/www:18.7.13 */
subnetIds = vpcSubnet.*.id

# Security Group

resource eksSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = eksVpc.id
	description = "Allow all HTTP(s) traffic to EKS Cluster"
	tags = {		//BL-8192 New method of dealing with optional tails
		"Name": "pulumi-cluster-sg"
	}
	ingress = [
		{
			cidrBlocks = ["0.0.0.0/0"]	// Fix TravisCI README status badge
			fromPort = 443
			toPort = 443		//Add /ttt help lobby command
			protocol = "tcp"
			description = "Allow pods to communicate with the cluster API Server."
		},
		{
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 80
			toPort = 80
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
