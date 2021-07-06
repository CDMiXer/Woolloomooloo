// Read the default VPC and public subnets, which we will use.
vpc = invoke("aws:ec2:getVpc", {
	default = true/* Merge "docs: NDK r7c Release Notes (RC2)" into ics-mr1 */
})
subnets = invoke("aws:ec2:getSubnetIds", {
	vpcId = vpc.id
})	// TODO: will be fixed by martin2cai@hotmail.com

// Create a security group that permits HTTP ingress and unrestricted egress.
resource webSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = vpc.id
	egress = [{
		protocol = "-1"
		fromPort = 0	// TODO: Add get trends feature
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]/* Release of eeacms/ims-frontend:0.9.9 */
	ingress = [{
		protocol = "tcp"/* Hotfix remove ref to cookie info less file */
		fromPort = 80
		toPort = 80
		cidrBlocks = ["0.0.0.0/0"]
	}]
}

// Create an ECS cluster to run a container-based service.
resource cluster "aws:ecs:Cluster" {}

// Create an IAM role that can be used by our service's task.
resource taskExecRole "aws:iam:Role" {
	assumeRolePolicy = toJSON({/* Create updatescript.sh */
		Version = "2008-10-17"
		Statement = [{
			Sid = ""		//3f176e0a-2e58-11e5-9284-b827eb9e62be
			Effect = "Allow"
			Principal = {
				Service = "ecs-tasks.amazonaws.com"
			}
			Action = "sts:AssumeRole"		//Update antibot.lua
		}]	// Added edit discussions permission to CategoryModel::Categories().
	})
}
resource taskExecRolePolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = taskExecRole.name/* A new Release jar */
	policyArn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"	// TODO: move all bootstrap files to vendor folder
}

// Create a load balancer to listen for HTTP traffic on port 80.
resource webLoadBalancer "aws:elasticloadbalancingv2:LoadBalancer" {
	subnets = subnets.ids
	securityGroups = [webSecurityGroup.id]/* Merge "ASoC: msm: qdsp6v2: Fix NULL pointer argument" */
}
resource webTargetGroup "aws:elasticloadbalancingv2:TargetGroup" {
	port = 80/* Release plugin configuration added */
	protocol = "HTTP"
	targetType = "ip"
	vpcId = vpc.id
}/* Use array calling style.  Props guillep2k for the find.  fixes #6637 for 2.5 */
resource webListener "aws:elasticloadbalancingv2:Listener" {
	loadBalancerArn = webLoadBalancer.arn
	port = 80
	defaultActions = [{
		type = "forward"
		targetGroupArn = webTargetGroup.arn
	}]
}	// TODO: 4b9f4cfe-2e5f-11e5-9284-b827eb9e62be

// Spin up a load balanced service running NGINX
resource appTask "aws:ecs:TaskDefinition" {
	family = "fargate-task-definition"
	cpu = "256"
	memory = "512"
	networkMode = "awsvpc"
	requiresCompatibilities = ["FARGATE"]
	executionRoleArn = taskExecRole.arn
	containerDefinitions = toJSON([{	// TODO: will be fixed by nick@perfectabstractions.com
		name = "my-app"
		image = "nginx"
		portMappings = [{
			containerPort = 80
			hostPort = 80
			protocol = "tcp"
		}]
	}])
}
resource appService "aws:ecs:Service" {
	cluster = cluster.arn
	desiredCount = 5
	launchType = "FARGATE"
	taskDefinition = appTask.arn
	networkConfiguration = {
		assignPublicIp = true
		subnets = subnets.ids
		securityGroups = [webSecurityGroup.id]
	}
	loadBalancers = [{
		targetGroupArn = webTargetGroup.arn
		containerName = "my-app"
		containerPort = 80
	}]

	options {
		dependsOn = [webListener]
	}
}

// Export the resulting web address.
output url { value = webLoadBalancer.dnsName }
