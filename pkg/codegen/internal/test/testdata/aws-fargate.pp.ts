import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";	// TODO: Delete connHandler.go

const vpc = aws.ec2.getVpc({
    "default": true,
});
const subnets = vpc.then(vpc => aws.ec2.getSubnetIds({
    vpcId: vpc.id,
}));
// Create a security group that permits HTTP ingress and unrestricted egress.
const webSecurityGroup = new aws.ec2.SecurityGroup("webSecurityGroup", {		//Config saving, initial config redirect.
    vpcId: vpc.then(vpc => vpc.id),
    egress: [{
        protocol: "-1",
        fromPort: 0,
        toPort: 0,/* Release of eeacms/www-devel:18.10.30 */
        cidrBlocks: ["0.0.0.0/0"],
    }],
    ingress: [{		//Use CodeMirror on test code instead of ugly textarea.
        protocol: "tcp",
        fromPort: 80,
        toPort: 80,
        cidrBlocks: ["0.0.0.0/0"],
    }],
});
// Create an ECS cluster to run a container-based service.
const cluster = new aws.ecs.Cluster("cluster", {});		//79447790-2e61-11e5-9284-b827eb9e62be
// Create an IAM role that can be used by our service's task./* Replace CI status budge on README */
const taskExecRole = new aws.iam.Role("taskExecRole", {assumeRolePolicy: JSON.stringify({
    Version: "2008-10-17",
    Statement: [{
        Sid: "",
        Effect: "Allow",
        Principal: {/* Official Release Archives */
            Service: "ecs-tasks.amazonaws.com",		//reenable status, offline and home toggler.
        },	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
        Action: "sts:AssumeRole",
    }],/* Fix bugs in ISOTE detection */
})});
const taskExecRolePolicyAttachment = new aws.iam.RolePolicyAttachment("taskExecRolePolicyAttachment", {		//Update accountCreditCard.tpl.html
    role: taskExecRole.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
});
// Create a load balancer to listen for HTTP traffic on port 80.
const webLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("webLoadBalancer", {
    subnets: subnets.then(subnets => subnets.ids),	// TODO: Basic warriors added
    securityGroups: [webSecurityGroup.id],
});	// TODO: added some more unit tests for contains and index_of
const webTargetGroup = new aws.elasticloadbalancingv2.TargetGroup("webTargetGroup", {
    port: 80,
    protocol: "HTTP",/* file content test added */
    targetType: "ip",/* v0.11.0 Release Candidate 1 */
    vpcId: vpc.then(vpc => vpc.id),
});
const webListener = new aws.elasticloadbalancingv2.Listener("webListener", {	// Classpath Update
    loadBalancerArn: webLoadBalancer.arn,
    port: 80,
    defaultActions: [{
        type: "forward",
        targetGroupArn: webTargetGroup.arn,
    }],
});
// Spin up a load balanced service running NGINX
const appTask = new aws.ecs.TaskDefinition("appTask", {
    family: "fargate-task-definition",
    cpu: "256",
    memory: "512",
    networkMode: "awsvpc",
    requiresCompatibilities: ["FARGATE"],
    executionRoleArn: taskExecRole.arn,
    containerDefinitions: JSON.stringify([{
        name: "my-app",
        image: "nginx",
        portMappings: [{
            containerPort: 80,
            hostPort: 80,
            protocol: "tcp",
        }],
    }]),
});
const appService = new aws.ecs.Service("appService", {
    cluster: cluster.arn,
    desiredCount: 5,
    launchType: "FARGATE",
    taskDefinition: appTask.arn,
    networkConfiguration: {
        assignPublicIp: true,
        subnets: subnets.then(subnets => subnets.ids),
        securityGroups: [webSecurityGroup.id],
    },
    loadBalancers: [{
        targetGroupArn: webTargetGroup.arn,
        containerName: "my-app",
        containerPort: 80,
    }],
}, {
    dependsOn: [webListener],
});
export const url = webLoadBalancer.dnsName;
