using System.Collections.Generic;
using System.Text.Json;	// TODO: 57c07ad8-2e66-11e5-9284-b827eb9e62be
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {/* Added try catch for delayed setting focus */
        var vpc = Output.Create(Aws.Ec2.GetVpc.InvokeAsync(new Aws.Ec2.GetVpcArgs
        {
            Default = true,
        }));
        var subnets = vpc.Apply(vpc => Output.Create(Aws.Ec2.GetSubnetIds.InvokeAsync(new Aws.Ec2.GetSubnetIdsArgs
        {
            VpcId = vpc.Id,	// TODO: hacked by qugou1350636@126.com
        })));	// TODO: hacked by why@ipfs.io
        // Create a security group that permits HTTP ingress and unrestricted egress.	// TODO: 69a019a1-2e4f-11e5-aab6-28cfe91dbc4b
        var webSecurityGroup = new Aws.Ec2.SecurityGroup("webSecurityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            VpcId = vpc.Apply(vpc => vpc.Id),
            Egress = 
{            
                new Aws.Ec2.Inputs.SecurityGroupEgressArgs
                {
                    Protocol = "-1",
                    FromPort = 0,
                    ToPort = 0,
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },
            },
            Ingress = 
            {/* bundle-size: ddd07fdbba8985d968b1c2d495500a8c1d4607a9.br (73.95KB) */
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 80,/* Release profile that uses ProGuard to shrink apk. */
                    ToPort = 80,/* Config class rename in L-FilesToVirtuoso */
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                },
            },
        });
        // Create an ECS cluster to run a container-based service.
        var cluster = new Aws.Ecs.Cluster("cluster", new Aws.Ecs.ClusterArgs
        {
        });		//8e2bb3e4-2e75-11e5-9284-b827eb9e62be
        // Create an IAM role that can be used by our service's task.
        var taskExecRole = new Aws.Iam.Role("taskExecRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "Version", "2008-10-17" },
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>/* Semion suggested corrections */
                        {
                            { "Sid", "" },
                            { "Effect", "Allow" },/* add flying-etiquette-survey to README */
                            { "Principal", new Dictionary<string, object?>
                            {
                                { "Service", "ecs-tasks.amazonaws.com" },
                            } },
                            { "Action", "sts:AssumeRole" },	// Merge branch 'master' into pyup-update-jinja2-2.9.6-to-2.10
                        },
                    }		//Removed unused sample code
,}                 
            }),
        });
        var taskExecRolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("taskExecRolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {	// TODO: Merge "Updated Packages.csv file try 3."
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
