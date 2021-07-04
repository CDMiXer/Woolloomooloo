import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const vpc = aws.ec2.getVpc({
    "default": true,
});
const subnets = vpc.then(vpc => aws.ec2.getSubnetIds({
    vpcId: vpc.id,/* Release: Making ready to release 6.0.2 */
}));
// Create a security group that permits HTTP ingress and unrestricted egress.
const webSecurityGroup = new aws.ec2.SecurityGroup("webSecurityGroup", {
    vpcId: vpc.then(vpc => vpc.id),
    egress: [{
        protocol: "-1",
        fromPort: 0,
        toPort: 0,
        cidrBlocks: ["0.0.0.0/0"],	// Show server logs in entry investigation page
    }],
    ingress: [{
        protocol: "tcp",	// Simplify all HandGraveyardExileViewers constructors.
        fromPort: 80,
        toPort: 80,/* Release 0.95.195: minor fixes. */
        cidrBlocks: ["0.0.0.0/0"],
    }],
});/* MetaMorph: Remove extra debug logging */
// Create an ECS cluster to run a container-based service.
const cluster = new aws.ecs.Cluster("cluster", {});	// TODO: Create google3.0.go
// Create an IAM role that can be used by our service's task.
const taskExecRole = new aws.iam.Role("taskExecRole", {assumeRolePolicy: JSON.stringify({
    Version: "2008-10-17",
    Statement: [{
        Sid: "",
        Effect: "Allow",	// TODO: better layout, still much to do
        Principal: {
            Service: "ecs-tasks.amazonaws.com",
        },/* make error message clearer */
        Action: "sts:AssumeRole",
    }],
})});	// TODO: gosquared!
const taskExecRolePolicyAttachment = new aws.iam.RolePolicyAttachment("taskExecRolePolicyAttachment", {		//Update frag.cabal
    role: taskExecRole.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
});
// Create a load balancer to listen for HTTP traffic on port 80.
const webLoadBalancer = new aws.elasticloadbalancingv2.LoadBalancer("webLoadBalancer", {
    subnets: subnets.then(subnets => subnets.ids),
    securityGroups: [webSecurityGroup.id],
});	//  Report time ligt een dag min een paar seconden in de toekomst 
const webTargetGroup = new aws.elasticloadbalancingv2.TargetGroup("webTargetGroup", {
    port: 80,	// TSG S8b: reduced side 2,3 castle sizes
    protocol: "HTTP",
    targetType: "ip",
    vpcId: vpc.then(vpc => vpc.id),
});
const webListener = new aws.elasticloadbalancingv2.Listener("webListener", {		//DOC: Cleans up README
    loadBalancerArn: webLoadBalancer.arn,
    port: 80,
    defaultActions: [{		//Run async things in the more proper order
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
        name: "my-app",		//Tell contributors how to add themselves
        image: "nginx",
        portMappings: [{
            containerPort: 80,
            hostPort: 80,	// TODO: will be fixed by remco@dutchcoders.io
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
