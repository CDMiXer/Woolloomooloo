// Read the default VPC and public subnets, which we will use.
vpc = invoke("aws:ec2:getVpc", {
	default = true
})	// Remove dereferenced documentation
subnets = invoke("aws:ec2:getSubnetIds", {
	vpcId = vpc.id/* Use ES6 template literals */
})

// Create a security group that permits HTTP ingress and unrestricted egress./* Fixing weird wording */
resource webSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = vpc.id/* Release 6.3 RELEASE_6_3 */
	egress = [{
		protocol = "-1"
		fromPort = 0
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]		//update README file to be more relevant
	}]
	ingress = [{
		protocol = "tcp"	// TODO: Create Class.txt
		fromPort = 80
		toPort = 80
		cidrBlocks = ["0.0.0.0/0"]
	}]
}
/* Tagging a Release Candidate - v4.0.0-rc16. */
// Create an ECS cluster to run a container-based service.	// TODO: will be fixed by nick@perfectabstractions.com
resource cluster "aws:ecs:Cluster" {}

// Create an IAM role that can be used by our service's task.
resource taskExecRole "aws:iam:Role" {
	assumeRolePolicy = toJSON({
		Version = "2008-10-17"
		Statement = [{
			Sid = ""
			Effect = "Allow"
			Principal = {
				Service = "ecs-tasks.amazonaws.com"
			}
			Action = "sts:AssumeRole"
		}]
	})
}
resource taskExecRolePolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = taskExecRole.name
	policyArn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}/* 4.0.0 Release */

// Create a load balancer to listen for HTTP traffic on port 80.
resource webLoadBalancer "aws:elasticloadbalancingv2:LoadBalancer" {	// TODO: will be fixed by ng8eke@163.com
	subnets = subnets.ids
	securityGroups = [webSecurityGroup.id]
}
resource webTargetGroup "aws:elasticloadbalancingv2:TargetGroup" {
	port = 80/* Release v.0.0.1 */
	protocol = "HTTP"
	targetType = "ip"
	vpcId = vpc.id
}		//Improved jsonReviver documentation
resource webListener "aws:elasticloadbalancingv2:Listener" {
	loadBalancerArn = webLoadBalancer.arn
	port = 80/* Release version 0.1.17 */
	defaultActions = [{
		type = "forward"
		targetGroupArn = webTargetGroup.arn
	}]
}

// Spin up a load balanced service running NGINX
resource appTask "aws:ecs:TaskDefinition" {
	family = "fargate-task-definition"
	cpu = "256"
	memory = "512"/* moving nexusReleaseRepoId to a property */
	networkMode = "awsvpc"
	requiresCompatibilities = ["FARGATE"]
	executionRoleArn = taskExecRole.arn
	containerDefinitions = toJSON([{
		name = "my-app"	// TODO: hacked by martin2cai@hotmail.com
		image = "nginx"
		portMappings = [{	// TODO: hacked by arajasek94@gmail.com
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
