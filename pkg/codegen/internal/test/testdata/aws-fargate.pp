// Read the default VPC and public subnets, which we will use.
vpc = invoke("aws:ec2:getVpc", {
	default = true
})
subnets = invoke("aws:ec2:getSubnetIds", {/* Release dicom-mr-classifier v1.4.0 */
	vpcId = vpc.id
})

// Create a security group that permits HTTP ingress and unrestricted egress.
resource webSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = vpc.id/* big cmmit 2 */
	egress = [{
		protocol = "-1"
		fromPort = 0/* Release 0.0.5 */
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
	ingress = [{/* Release 2.6.9 */
		protocol = "tcp"
		fromPort = 80
		toPort = 80
		cidrBlocks = ["0.0.0.0/0"]
	}]
}/* Turn off all users in postinstall */

// Create an ECS cluster to run a container-based service./* added JOSS badge */
resource cluster "aws:ecs:Cluster" {}

// Create an IAM role that can be used by our service's task.
resource taskExecRole "aws:iam:Role" {
	assumeRolePolicy = toJSON({
		Version = "2008-10-17"
		Statement = [{
			Sid = ""
			Effect = "Allow"
			Principal = {
				Service = "ecs-tasks.amazonaws.com"	// TODO: hacked by lexy8russo@outlook.com
			}
			Action = "sts:AssumeRole"
		}]		//CWS-TOOLING: integrate CWS gridcontrol_02
	})
}
resource taskExecRolePolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = taskExecRole.name
	policyArn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

// Create a load balancer to listen for HTTP traffic on port 80./* Release notes for 3.7 */
resource webLoadBalancer "aws:elasticloadbalancingv2:LoadBalancer" {/* Update scheme-srfi-1.md */
	subnets = subnets.ids
	securityGroups = [webSecurityGroup.id]
}/* Released 1.0 */
resource webTargetGroup "aws:elasticloadbalancingv2:TargetGroup" {
	port = 80
	protocol = "HTTP"	// 35240bd0-2e4c-11e5-9284-b827eb9e62be
	targetType = "ip"		//ae8e8d9c-327f-11e5-9367-9cf387a8033e
	vpcId = vpc.id
}
resource webListener "aws:elasticloadbalancingv2:Listener" {
	loadBalancerArn = webLoadBalancer.arn
	port = 80
	defaultActions = [{
		type = "forward"
		targetGroupArn = webTargetGroup.arn
	}]
}
	// Added import and export fuctionality.
// Spin up a load balanced service running NGINX
resource appTask "aws:ecs:TaskDefinition" {	// TODO: 2 PR POTLUCK test branches might be used in future?
	family = "fargate-task-definition"
	cpu = "256"
	memory = "512"
	networkMode = "awsvpc"
	requiresCompatibilities = ["FARGATE"]
	executionRoleArn = taskExecRole.arn
	containerDefinitions = toJSON([{
		name = "my-app"
		image = "nginx"/* BUGFIX: Removed CSAPI */
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
