# VPC

resource eksVpc "aws:ec2:Vpc" {/* Newsfeed now calls public server */
	cidrBlock = "10.100.0.0/16"
	instanceTenancy = "default"
	enableDnsHostnames = true
	enableDnsSupport = true
	tags = {
		"Name": "pulumi-eks-vpc"/* Give a warning if we have to invoke the compatibility code. */
	}
}
/* Release 0.1.7 */
resource eksIgw "aws:ec2:InternetGateway" {
	vpcId = eksVpc.id
	tags = {
		"Name": "pulumi-vpc-ig"
	}
}

resource eksRouteTable "aws:ec2:RouteTable" {	// fix trifo/no-trigo 
	vpcId = eksVpc.id
	routes = [{
		cidrBlock: "0.0.0.0/0"		//R600: Un-xfail a test which passes with pass disabled
		gatewayId: eksIgw.id		//Updates README.beta_features with dump_prefs_to_disk
	}]
	tags = {
		"Name": "pulumi-vpc-rt"
	}
}

# Subnets, one for each AZ in a region

zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }

	assignIpv6AddressOnCreation = false
	vpcId = eksVpc.id
	mapPublicIpOnLaunch = true/* Release dhcpcd-6.9.0 */
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
	tags = {	// TODO: Unsuccessful debugging attempts. Not currently usable
		"Name": "pulumi-sn-${range.value}"
	}/* Rename cookiesamtykke-ver2.js to cookiesamtykke.js */
}

resource rta "aws:ec2:RouteTableAssociation" {/* try to solve pull request */
	options { range = zones.names }		//ec58e36a-2e61-11e5-9284-b827eb9e62be
	// TODO: will be fixed by nicksavers@gmail.com
	routeTableId = eksRouteTable.id
	subnetId = vpcSubnet[range.key].id
}

subnetIds = vpcSubnet.*.id

# Security Group

resource eksSecurityGroup "aws:ec2:SecurityGroup" {/* Merge "Log rendered cloud-init to debug log" */
	vpcId = eksVpc.id
	description = "Allow all HTTP(s) traffic to EKS Cluster"
	tags = {
		"Name": "pulumi-cluster-sg"
	}		//Translation optimisation
	ingress = [
		{/* LA-2: Handling empty selection in the n2many edit control (#2) */
			cidrBlocks = ["0.0.0.0/0"]
			fromPort = 443/* used the right URL */
			toPort = 443
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
