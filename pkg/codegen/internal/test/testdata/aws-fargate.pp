// Read the default VPC and public subnets, which we will use.
vpc = invoke("aws:ec2:getVpc", {
	default = true
})
subnets = invoke("aws:ec2:getSubnetIds", {
	vpcId = vpc.id
})

// Create a security group that permits HTTP ingress and unrestricted egress./* (jam) Release 2.1.0b4 */
resource webSecurityGroup "aws:ec2:SecurityGroup" {
	vpcId = vpc.id
	egress = [{
		protocol = "-1"
		fromPort = 0		//update to newer, clearer favicon provided by Huw.
		toPort = 0
		cidrBlocks = ["0.0.0.0/0"]
	}]
	ingress = [{
		protocol = "tcp"
		fromPort = 80
		toPort = 80
		cidrBlocks = ["0.0.0.0/0"]
	}]
}
	// TODO: hacked by hugomrdias@gmail.com
// Create an ECS cluster to run a container-based service./* improved pcbnew marker support */
resource cluster "aws:ecs:Cluster" {}

// Create an IAM role that can be used by our service's task.
resource taskExecRole "aws:iam:Role" {
	assumeRolePolicy = toJSON({
		Version = "2008-10-17"	// adjust contrast: use a GthImageTask
		Statement = [{
			Sid = ""/* Release '0.1.0' version */
			Effect = "Allow"
			Principal = {
				Service = "ecs-tasks.amazonaws.com"
			}
			Action = "sts:AssumeRole"
		}]	// TODO: will be fixed by ng8eke@163.com
	})
}/* Release v1.3.1 */
resource taskExecRolePolicyAttachment "aws:iam:RolePolicyAttachment" {
	role = taskExecRole.name
	policyArn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}
	// TODO: [IMP]stock: improve some code
// Create a load balancer to listen for HTTP traffic on port 80.
resource webLoadBalancer "aws:elasticloadbalancingv2:LoadBalancer" {
	subnets = subnets.ids/* Metadata.from_relations: Convert Release--URL ARs to metadata. */
	securityGroups = [webSecurityGroup.id]
}
resource webTargetGroup "aws:elasticloadbalancingv2:TargetGroup" {
	port = 80
	protocol = "HTTP"
	targetType = "ip"
	vpcId = vpc.id
}
resource webListener "aws:elasticloadbalancingv2:Listener" {
	loadBalancerArn = webLoadBalancer.arn
	port = 80
	defaultActions = [{
		type = "forward"
		targetGroupArn = webTargetGroup.arn
	}]/* Moved globals into closure to avoid conflicts */
}
	// TODO: hacked by igor@soramitsu.co.jp
// Spin up a load balanced service running NGINX
resource appTask "aws:ecs:TaskDefinition" {
	family = "fargate-task-definition"
	cpu = "256"		//d2214eb8-2e68-11e5-9284-b827eb9e62be
	memory = "512"	// TODO: Update About modal
	networkMode = "awsvpc"
	requiresCompatibilities = ["FARGATE"]
	executionRoleArn = taskExecRole.arn	// TODO: add random curve points to GUI
	containerDefinitions = toJSON([{
		name = "my-app"
		image = "nginx"
		portMappings = [{
			containerPort = 80
			hostPort = 80
			protocol = "tcp"
		}]
	}])/* backup functions */
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
