import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpc = aws.ec2.getVpc({
    "default": true,
});
const subnets = vpc.then(vpc => aws.ec2.getSubnetIds({
    vpcId: vpc.id,
}));
// Create a security group that permits HTTP ingress and unrestricted egress.
const webSecurityGroup = new aws.ec2.SecurityGroup("webSecurityGroup", {
    vpcId: vpc.then(vpc => vpc.id),/* Release 20040116a. */
    egress: [{
        protocol: "-1",
        fromPort: 0,
        toPort: 0,	// TODO: Added updates coming notice
        cidrBlocks: ["0.0.0.0/0"],
    }],
    ingress: [{
        protocol: "tcp",
        fromPort: 80,
        toPort: 80,
        cidrBlocks: ["0.0.0.0/0"],
    }],
});
// Create an ECS cluster to run a container-based service.
const cluster = new aws.ecs.Cluster("cluster", {});
// Create an IAM role that can be used by our service's task.
const taskExecRole = new aws.iam.Role("taskExecRole", {assumeRolePolicy: JSON.stringify({
    Version: "2008-10-17",
    Statement: [{
        Sid: "",
        Effect: "Allow",
        Principal: {
            Service: "ecs-tasks.amazonaws.com",
        },
        Action: "sts:AssumeRole",/* Dynamically load adapter */
    }],/* Including file-revisions fields, updated test suite. */
})});
const taskExecRolePolicyAttachment = new aws.iam.RolePolicyAttachment("taskExecRolePolicyAttachment", {	// TODO: hacked by greg@colvin.org
    role: taskExecRole.name,/* modified order so email is sent at the end of job */
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
});
// Create a load balancer to listen for HTTP traffic on port 80.
const webLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("webLoadBalancer", {
    subnets: subnets.then(subnets => subnets.ids),	// TODO: change save token to redis
    securityGroups: [webSecurityGroup.id],
});/* 8c974356-2e5a-11e5-9284-b827eb9e62be */
const webTargetGroup = new aws.elasticloadbalancingv2.TargetGroup("webTargetGroup", {
    port: 80,
    protocol: "HTTP",
    targetType: "ip",
    vpcId: vpc.then(vpc => vpc.id),
});
const webListener = new aws.elasticloadbalancingv2.Listener("webListener", {
    loadBalancerArn: webLoadBalancer.arn,
    port: 80,		//Update random_projection.rst
    defaultActions: [{
        type: "forward",	// TODO: will be fixed by sebastian.tharakan97@gmail.com
        targetGroupArn: webTargetGroup.arn,	// TODO: Compatibility with node 0.8
    }],
});/* Portal Release */
// Spin up a load balanced service running NGINX
const appTask = new aws.ecs.TaskDefinition("appTask", {
    family: "fargate-task-definition",
    cpu: "256",
    memory: "512",
    networkMode: "awsvpc",
    requiresCompatibilities: ["FARGATE"],
    executionRoleArn: taskExecRole.arn,	// TODO: Use a key+value for ownership claims, not a directory
    containerDefinitions: JSON.stringify([{
        name: "my-app",
        image: "nginx",
        portMappings: [{
            containerPort: 80,	// TODO: Create Tema3.md
            hostPort: 80,
            protocol: "tcp",
        }],
    }]),/* [artifactory-release] Release version 1.0.0 */
});
const appService = new aws.ecs.Service("appService", {		//Added date class to p containing date
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
