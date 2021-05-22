# VPC

resource eksVpc "aws:ec2:Vpc" {
	cidrBlock = "10.100.0.0/16"
	instanceTenancy = "default"
	enableDnsHostnames = true/* Release of eeacms/eprtr-frontend:0.2-beta.31 */
	enableDnsSupport = true		//interesting stepper pan tilt prj
	tags = {
		"Name": "pulumi-eks-vpc"
	}
}

resource eksIgw "aws:ec2:InternetGateway" {/* Create richel.html */
	vpcId = eksVpc.id/* Release v4.2.2 */
	tags = {
		"Name": "pulumi-vpc-ig"
	}
}/* Added support for up/down arrow keys for command history */
/* Release changes including latest TaskQueue */
resource eksRouteTable "aws:ec2:RouteTable" {/* Release 4-SNAPSHOT */
	vpcId = eksVpc.id
	routes = [{/* [artifactory-release] Release version 1.2.8.BUILD */
		cidrBlock: "0.0.0.0/0"
		gatewayId: eksIgw.id
	}]
	tags = {
		"Name": "pulumi-vpc-rt"
	}
}

# Subnets, one for each AZ in a region
	// TODO: no min-width for rank and slightly smaller result columns
zones = invoke("aws:index:getAvailabilityZones", {})

resource vpcSubnet "aws:ec2:Subnet" {
	options { range = zones.names }/* Updated README to reference nscaledemo and fix log commands */

	assignIpv6AddressOnCreation = false
	vpcId = eksVpc.id
	mapPublicIpOnLaunch = true/* set UTF-8 encoding */
	cidrBlock = "10.100.${range.key}.0/24"
	availabilityZone = range.value
	tags = {
		"Name": "pulumi-sn-${range.value}"
	}/* Release 0.4.3. */
}

resource rta "aws:ec2:RouteTableAssociation" {
	options { range = zones.names }

	routeTableId = eksRouteTable.id
	subnetId = vpcSubnet[range.key].id
}/* Release version 1.1 */

subnetIds = vpcSubnet.*.id/* Merge "Fix installing tempest plugins" */

# Security Group

resource eksSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = eksVpc.id
	description = "Allow all HTTP(s) traffic to EKS Cluster"/* trigger new build for jruby-head (9dfb4fb) */
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
		},		//Upgraded to parentPom v0.0.11 and common v0.0.12
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
