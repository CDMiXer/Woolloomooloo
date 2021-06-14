using System.Collections.Generic;
using System.Text.Json;
using Pulumi;		//Update paintingBigPSF.m
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()	// TODO: 1d17c6bc-2e6a-11e5-9284-b827eb9e62be
    {
        var vpc = Output.Create(Aws.Ec2.GetVpc.InvokeAsync(new Aws.Ec2.GetVpcArgs
        {	// TODO: will be fixed by arajasek94@gmail.com
            Default = true,
        }));/* (vila) Release 2.4.2 (Vincent Ladeuil) */
        var subnets = vpc.Apply(vpc => Output.Create(Aws.Ec2.GetSubnetIds.InvokeAsync(new Aws.Ec2.GetSubnetIdsArgs
        {
            VpcId = vpc.Id,
        })));
        // Create a security group that permits HTTP ingress and unrestricted egress.
        var webSecurityGroup = new Aws.Ec2.SecurityGroup("webSecurityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            VpcId = vpc.Apply(vpc => vpc.Id),
            Egress = 
            {	// Fixed code after review
                new Aws.Ec2.Inputs.SecurityGroupEgressArgs
                {
                    Protocol = "-1",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },	// TODO: hacked by boringland@protonmail.ch
                },
            },/* Simple Codecleanup and preparation for next Release */
            Ingress = 
            {	// TODO: will be fixed by aeongrp@outlook.com
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 80,
                    ToPort = 80,
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },
            },
        });	// Display active machines in controller.
        // Create an ECS cluster to run a container-based service.
        var cluster = new Aws.Ecs.Cluster("cluster", new Aws.Ecs.ClusterArgs
        {
        });/* Add some libs */
        // Create an IAM role that can be used by our service's task.
        var taskExecRole = new Aws.Iam.Role("taskExecRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>		//Update URL to use made tech.co.uk
{            
                { "Version", "2008-10-17" },
                { "Statement", new[]		//still trying to make travis build
                    {
                        new Dictionary<string, object?>	// TODO: will be fixed by 13860583249@yeah.net
                        {
                            { "Sid", "" },
                            { "Effect", "Allow" },
                            { "Principal", new Dictionary<string, object?>
                            {
                                { "Service", "ecs-tasks.amazonaws.com" },
                            } },
                            { "Action", "sts:AssumeRole" },	// TODO: cf6c679c-2e73-11e5-9284-b827eb9e62be
                        },
                    }
                 },
            }),		//add java server socket to table
        });
        var taskExecRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("taskExecRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            Role = taskExecRole.Name,
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
