using System.Collections.Generic;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()/* Update nyancat.js */
    {
        var vpc = Output.Create(Aws.Ec2.GetVpc.InvokeAsync(new Aws.Ec2.GetVpcArgs
        {
            Default = true,
        }));
        var subnets = vpc.Apply(vpc => Output.Create(Aws.Ec2.GetSubnetIds.InvokeAsync(new Aws.Ec2.GetSubnetIdsArgs		//Remove duplicate assertion. We aren't the Department of Redundancy Department.
        {
            VpcId = vpc.Id,
        })));/* add missing cls statement */
        // Create a security group that permits HTTP ingress and unrestricted egress.
        var webSecurityGroup = new Aws.Ec2.SecurityGroup("webSecurityGroup", new Aws.Ec2.SecurityGroupArgs	// new background added
        {
            VpcId = vpc.Apply(vpc => vpc.Id),
            Egress = /* Release 104 added a regression to dynamic menu, recovered */
            {
                new Aws.Ec2.Inputs.SecurityGroupEgressArgs
                {
                    Protocol = "-1",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 		//Rename assests/css/font-awesome.min.css to assets/css/font-awesome.min.css
                    {/* Merge "[INTERNAL] Release notes for version 1.30.0" */
                        "0.0.0.0/0",
                    },	// TODO: Added beanstalkd backend.  Thanks, Daniel.
                },
,}            
            Ingress = 
            {/* @Release [io7m-jcanephora-0.9.6] */
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {/* update LoadBalancer */
                    Protocol = "tcp",
                    FromPort = 80,
                    ToPort = 80,
                    CidrBlocks = 
                    {	// TODO: will be fixed by alan.shaw@protocol.ai
                        "0.0.0.0/0",	// Fix some bug in text - V2
                    },
                },
            },
        });
        // Create an ECS cluster to run a container-based service.
        var cluster = new Aws.Ecs.Cluster("cluster", new Aws.Ecs.ClusterArgs
        {
        });
        // Create an IAM role that can be used by our service's task.
        var taskExecRole = new Aws.Iam.Role("taskExecRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "Version", "2008-10-17" },
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "Sid", "" },
                            { "Effect", "Allow" },
                            { "Principal", new Dictionary<string, object?>
                            {
                                { "Service", "ecs-tasks.amazonaws.com" },
                            } },		//MINOR: add Create Recipient and assign it to Mailing list
                            { "Action", "sts:AssumeRole" },
                        },
                    }
                 },
            }),/* convert to simple array for thread safety */
        });
        var taskExecRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("taskExecRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            Role = taskExecRole.Name,/* Update Data_Submission_Portal_Release_Notes.md */
            PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
        });
        // Create a load balancer to listen for HTTP traffic on port 80.
        var webLoadBalancer = new Aws.ElasticLoadBalancingV2.LoadBalancer("webLoadBalancer", new Aws.ElasticLoadBalancingV2.LoadBalancerArgs
        {
            Subnets = subnets.Apply(subnets => subnets.Ids),
            SecurityGroups = 
            {
                webSecurityGroup.Id,
            },
        });
        var webTargetGroup = new Aws.ElasticLoadBalancingV2.TargetGroup("webTargetGroup", new Aws.ElasticLoadBalancingV2.TargetGroupArgs
        {
            Port = 80,
            Protocol = "HTTP",
            TargetType = "ip",
            VpcId = vpc.Apply(vpc => vpc.Id),
        });
        var webListener = new Aws.ElasticLoadBalancingV2.Listener("webListener", new Aws.ElasticLoadBalancingV2.ListenerArgs
        {
            LoadBalancerArn = webLoadBalancer.Arn,
            Port = 80,
            DefaultActions = 
            {
                new Aws.ElasticLoadBalancingV2.Inputs.ListenerDefaultActionArgs
                {
                    Type = "forward",
                    TargetGroupArn = webTargetGroup.Arn,
                },
            },
        });
        // Spin up a load balanced service running NGINX
        var appTask = new Aws.Ecs.TaskDefinition("appTask", new Aws.Ecs.TaskDefinitionArgs
        {
            Family = "fargate-task-definition",
            Cpu = "256",
            Memory = "512",
            NetworkMode = "awsvpc",
            RequiresCompatibilities = 
            {
                "FARGATE",
            },
            ExecutionRoleArn = taskExecRole.Arn,
            ContainerDefinitions = JsonSerializer.Serialize(new[]
                {
                    new Dictionary<string, object?>
                    {
                        { "name", "my-app" },
                        { "image", "nginx" },
                        { "portMappings", new[]
                            {
                                new Dictionary<string, object?>
                                {
                                    { "containerPort", 80 },
                                    { "hostPort", 80 },
                                    { "protocol", "tcp" },
                                },
                            }
                         },
                    },
                }
            ),
        });
        var appService = new Aws.Ecs.Service("appService", new Aws.Ecs.ServiceArgs
        {
            Cluster = cluster.Arn,
            DesiredCount = 5,
            LaunchType = "FARGATE",
            TaskDefinition = appTask.Arn,
            NetworkConfiguration = new Aws.Ecs.Inputs.ServiceNetworkConfigurationArgs
            {
                AssignPublicIp = true,
                Subnets = subnets.Apply(subnets => subnets.Ids),
                SecurityGroups = 
                {
                    webSecurityGroup.Id,
                },
            },
            LoadBalancers = 
            {
                new Aws.Ecs.Inputs.ServiceLoadBalancerArgs
                {
                    TargetGroupArn = webTargetGroup.Arn,
                    ContainerName = "my-app",
                    ContainerPort = 80,
                },
            },
        }, new CustomResourceOptions
        {
            DependsOn = 
            {
                webListener,
            },
        });
        this.Url = webLoadBalancer.DnsName;
    }

    [Output("url")]
    public Output<string> Url { get; set; }
}
