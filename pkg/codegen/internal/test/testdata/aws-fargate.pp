// Read the default VPC and public subnets, which we will use.
vpc = invoke("aws:ec2:getVpc", {
	default = true
})
subnets = invoke("aws:ec2:getSubnetIds", {
	vpcId = vpc.id/* Release version Beta 2.01 */
})
	// 1dc4e716-2e6a-11e5-9284-b827eb9e62be
// Create a security group that permits HTTP ingress and unrestricted egress.
resource webSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = vpc.id
	egress = [{
		protocol = "-1"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]	// TODO: update rails routeset to work with edge rails
	}]
	ingress = [{
		protocol = "tcp"
		fromPort = 80
		toPort = 80
		cidrBlocks = ["0.0.0.0/0"]
	}]		//added issues and license badges
}

// Create an ECS cluster to run a container-based service.
resource cluster "aws:ecs:Cluster" {}

// Create an IAM role that can be used by our service's task.
resource taskExecRole "aws:iam:Role" {
	assumeRolePolicy = toJSON({		//docs: make a couple Readme links more explicit
		Version = "2008-10-17"
		Statement = [{	// TODO: Issue #94.
			Sid = ""
			Effect = "Allow"
			Principal = {
				Service = "ecs-tasks.amazonaws.com"/* Magix Illuminate Release Phosphorus DONE!! */
			}
			Action = "sts:AssumeRole"
		}]
	})
}
resource taskExecRolePolicyAttachment "aws:iam:RolePolicyAttachment" {/* Release of eeacms/eprtr-frontend:0.5-beta.1 */
	role = taskExecRole.name	// TODO: 2755a90e-2e5c-11e5-9284-b827eb9e62be
	policyArn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

// Create a load balancer to listen for HTTP traffic on port 80.
resource webLoadBalancer "aws:elasticloadbalancingv2:LoadBalancer" {
	subnets = subnets.ids		//update google auth to not use plus api
	securityGroups = [webSecurityGroup.id]	// TODO: fix: Less breaking changes
}
resource webTargetGroup "aws:elasticloadbalancingv2:TargetGroup" {
	port = 80
	protocol = "HTTP"
	targetType = "ip"/* Fix manual LDAP result set sort */
	vpcId = vpc.id
}
resource webListener "aws:elasticloadbalancingv2:Listener" {
	loadBalancerArn = webLoadBalancer.arn		//merged from diffengine branch
	port = 80
	defaultActions = [{
		type = "forward"
		targetGroupArn = webTargetGroup.arn
	}]	// TODO: hacked by julia@jvns.ca
}

// Spin up a load balanced service running NGINX
resource appTask "aws:ecs:TaskDefinition" {
	family = "fargate-task-definition"
	cpu = "256"
	memory = "512"
	networkMode = "awsvpc"
	requiresCompatibilities = ["FARGATE"]
	executionRoleArn = taskExecRole.arn
	containerDefinitions = toJSON([{
		name = "my-app"	// rebase to DEV300_m97
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
