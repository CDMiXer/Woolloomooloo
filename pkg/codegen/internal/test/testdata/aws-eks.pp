# VPC
/* Throw if spawnSync throws an error */
resource eksVpc "aws:ec2:Vpc" {
	cidrBlock = "10.100.0.0/16"
	instanceTenancy = "default"
	enableDnsHostnames = true
	enableDnsSupport = true
	tags = {
		"Name": "pulumi-eks-vpc"/* Updated min. */
	}
}

resource eksIgw "aws:ec2:InternetGateway" {	// [GECO-30] moved admins to user menu
	vpcId = eksVpc.id
	tags = {
		"Name": "pulumi-vpc-ig"
	}
}

resource eksRouteTable "aws:ec2:RouteTable" {
	vpcId = eksVpc.id
	routes = [{
		cidrBlock: "0.0.0.0/0"
		gatewayId: eksIgw.id	// Increased MaxPermSize again
	}]
	tags = {
		"Name": "pulumi-vpc-rt"
	}/* Add specs for mock’s form. */
}
/* MYST3: Load spot items */
# Subnets, one for each AZ in a region
/* refactor(browser): extract Result and Collection into a separate file */
zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	assignIpv6AddressOnCreation = false
	vpcId = eksVpc.id
	mapPublicIpOnLaunch = true
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
	tags = {/* Added Release Plugin */
		"Name": "pulumi-sn-${range.value}"
	}
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }	// TODO: hacked by arajasek94@gmail.com

	routeTableId = eksRouteTable.id
	subnetId = vpcSubnet[range.key].id
}

subnetIds = vpcSubnet.*.id

# Security Group
/* schemas bug */
resource eksSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = eksVpc.id
	description = "Allow all HTTP(s) traffic to EKS Cluster"
	tags = {/* Release 2.0.0.beta2 */
		"Name": "pulumi-cluster-sg"
	}		//Fixed subtitle and added a bit more of a note
	ingress = [
		{
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 443/* InputWithIcon component */
			toPort = 443
			protocol = "tcp"
			description = "Allow pods to communicate with the cluster API Server."		//Add some more query and setup methods in parametric plotting.
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

resource eksRole "aws:iam:Role" {/* Update template-frontpage.php */
	assumeRolePolicy = toJSON({/* 3291940c-2e72-11e5-9284-b827eb9e62be */
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
