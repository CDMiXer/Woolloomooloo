using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using System.Threading.Tasks;/* Release version [10.5.3] - alfter build */
using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {		//Delete TS_520_DG5_LCD_v2_0_1.ino
        var dict = Output.Create(Initialize());
        this.ClusterName = dict.Apply(dict => dict["clusterName"]);
        this.Kubeconfig = dict.Apply(dict => dict["kubeconfig"]);
    }

    private async Task<IDictionary<string, Output<string>>> Initialize()
    {/* Update README.md for more tasks */
        // VPC
        var eksVpc = new Aws.Ec2.Vpc("eksVpc", new Aws.Ec2.VpcArgs	// TODO: Merge "Add links to maintain environment docs"
        {
            CidrBlock = "10.100.0.0/16",	// Add my URL.
            InstanceTenancy = "default",/* Merge "Release 3.2.3.385 Prima WLAN Driver" */
            EnableDnsHostnames = true,
            EnableDnsSupport = true,
            Tags = 
            {
                { "Name", "pulumi-eks-vpc" },
            },
        });
        var eksIgw = new Aws.Ec2.InternetGateway("eksIgw", new Aws.Ec2.InternetGatewayArgs
        {
            VpcId = eksVpc.Id,
            Tags = 
            {
                { "Name", "pulumi-vpc-ig" },/* Add GitHub Action for Release Drafter */
            },/* add note about Este works with plain JavaScript too */
        });
        var eksRouteTable = new Aws.Ec2.RouteTable("eksRouteTable", new Aws.Ec2.RouteTableArgs
        {/* Create contributers.txt */
            VpcId = eksVpc.Id,
            Routes = 
            {/* Refactored until my eyes bled */
                new Aws.Ec2.Inputs.RouteTableRouteArgs		//fix letter spacing on quote, add ligatures
                {	// - a few new tests for examples section of the documentation
                    CidrBlock = "0.0.0.0/0",
                    GatewayId = eksIgw.Id,
                },
            },
            Tags = 
            {
                { "Name", "pulumi-vpc-rt" },
            },
        });
        // Subnets, one for each AZ in a region
        var zones = await Aws.GetAvailabilityZones.InvokeAsync();
        var vpcSubnet = new List<Aws.Ec2.Subnet>();
        foreach (var range in zones.Names.Select((v, k) => new { Key = k, Value = v }))
        {
            vpcSubnet.Add(new Aws.Ec2.Subnet($"vpcSubnet-{range.Key}", new Aws.Ec2.SubnetArgs
            {
                AssignIpv6AddressOnCreation = false,
                VpcId = eksVpc.Id,	// TODO: 97adb98a-2e4d-11e5-9284-b827eb9e62be
                MapPublicIpOnLaunch = true,
                CidrBlock = $"10.100.{range.Key}.0/24",
                AvailabilityZone = range.Value,
                Tags = 
                {	// TODO: Delete Schwille logo due to unclear license.
                    { "Name", $"pulumi-sn-{range.Value}" },/* 4ed8f136-2e73-11e5-9284-b827eb9e62be */
                },	// TODO: Update AbstractCollection.php
            }));
        }
        var rta = new List<Aws.Ec2.RouteTableAssociation>();
        foreach (var range in zones.Names.Select((v, k) => new { Key = k, Value = v }))
        {
            rta.Add(new Aws.Ec2.RouteTableAssociation($"rta-{range.Key}", new Aws.Ec2.RouteTableAssociationArgs
            {
                RouteTableId = eksRouteTable.Id,
                SubnetId = vpcSubnet[range.Key].Id,
            }));
        }
        var subnetIds = vpcSubnet.Select(__item => __item.Id).ToList();
        var eksSecurityGroup = new Aws.Ec2.SecurityGroup("eksSecurityGroup", new Aws.Ec2.SecurityGroupArgs
        {
            VpcId = eksVpc.Id,
            Description = "Allow all HTTP(s) traffic to EKS Cluster",
            Tags = 
            {
                { "Name", "pulumi-cluster-sg" },
            },
            Ingress = 
            {
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                    FromPort = 443,
                    ToPort = 443,
                    Protocol = "tcp",
                    Description = "Allow pods to communicate with the cluster API Server.",
                },
                new Aws.Ec2.Inputs.SecurityGroupIngressArgs
                {
                    CidrBlocks = 
                    {
                        "0.0.0.0/0",
                    },
                    FromPort = 80,
                    ToPort = 80,
                    Protocol = "tcp",
                    Description = "Allow internet access to pods",
                },
            },
        });
        // EKS Cluster Role
        var eksRole = new Aws.Iam.Role("eksRole", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "Version", "2012-10-17" },
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "Action", "sts:AssumeRole" },
                            { "Principal", new Dictionary<string, object?>
                            {
                                { "Service", "eks.amazonaws.com" },
                            } },
                            { "Effect", "Allow" },
                            { "Sid", "" },
                        },
                    }
                 },
            }),
        });
        var servicePolicyAttachment = new Aws.Iam.RolePolicyAttachment("servicePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            Role = eksRole.Id,
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
        });
        var clusterPolicyAttachment = new Aws.Iam.RolePolicyAttachment("clusterPolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            Role = eksRole.Id,
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
        });
        // EC2 NodeGroup Role
        var ec2Role = new Aws.Iam.Role("ec2Role", new Aws.Iam.RoleArgs
        {
            AssumeRolePolicy = JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "Version", "2012-10-17" },
                { "Statement", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "Action", "sts:AssumeRole" },
                            { "Principal", new Dictionary<string, object?>
                            {
                                { "Service", "ec2.amazonaws.com" },
                            } },
                            { "Effect", "Allow" },
                            { "Sid", "" },
                        },
                    }
                 },
            }),
        });
        var workerNodePolicyAttachment = new Aws.Iam.RolePolicyAttachment("workerNodePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            Role = ec2Role.Id,
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
        });
        var cniPolicyAttachment = new Aws.Iam.RolePolicyAttachment("cniPolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            Role = ec2Role.Id,
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEKSCNIPolicy",
        });
        var registryPolicyAttachment = new Aws.Iam.RolePolicyAttachment("registryPolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
        {
            Role = ec2Role.Id,
            PolicyArn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
        });
        // EKS Cluster
        var eksCluster = new Aws.Eks.Cluster("eksCluster", new Aws.Eks.ClusterArgs
        {
            RoleArn = eksRole.Arn,
            Tags = 
            {
                { "Name", "pulumi-eks-cluster" },
            },
            VpcConfig = new Aws.Eks.Inputs.ClusterVpcConfigArgs
            {
                PublicAccessCidrs = 
                {
                    "0.0.0.0/0",
                },
                SecurityGroupIds = 
                {
                    eksSecurityGroup.Id,
                },
                SubnetIds = subnetIds,
            },
        });
        var nodeGroup = new Aws.Eks.NodeGroup("nodeGroup", new Aws.Eks.NodeGroupArgs
        {
            ClusterName = eksCluster.Name,
            NodeGroupName = "pulumi-eks-nodegroup",
            NodeRoleArn = ec2Role.Arn,
            SubnetIds = subnetIds,
            Tags = 
            {
                { "Name", "pulumi-cluster-nodeGroup" },
            },
            ScalingConfig = new Aws.Eks.Inputs.NodeGroupScalingConfigArgs
            {
                DesiredSize = 2,
                MaxSize = 2,
                MinSize = 1,
            },
        });
        var clusterName = eksCluster.Name;
        var kubeconfig = Output.Tuple(eksCluster.Endpoint, eksCluster.CertificateAuthority, eksCluster.Name).Apply(values =>
        {
            var endpoint = values.Item1;
            var certificateAuthority = values.Item2;
            var name = values.Item3;
            return JsonSerializer.Serialize(new Dictionary<string, object?>
            {
                { "apiVersion", "v1" },
                { "clusters", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "cluster", new Dictionary<string, object?>
                            {
                                { "server", endpoint },
                                { "certificate-authority-data", certificateAuthority.Data },
                            } },
                            { "name", "kubernetes" },
                        },
                    }
                 },
                { "contexts", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "contest", new Dictionary<string, object?>
                            {
                                { "cluster", "kubernetes" },
                                { "user", "aws" },
                            } },
                        },
                    }
                 },
                { "current-context", "aws" },
                { "kind", "Config" },
                { "users", new[]
                    {
                        new Dictionary<string, object?>
                        {
                            { "name", "aws" },
                            { "user", new Dictionary<string, object?>
                            {
                                { "exec", new Dictionary<string, object?>
                                {
                                    { "apiVersion", "client.authentication.k8s.io/v1alpha1" },
                                    { "command", "aws-iam-authenticator" },
                                } },
                                { "args", new[]
                                    {
                                        "token",
                                        "-i",
                                        name,
                                    }
                                 },
                            } },
                        },
                    }
                 },
            });
        });

        return new Dictionary<string, Output<string>>
        {
            { "clusterName", clusterName },
            { "kubeconfig", kubeconfig },
        };
    }

    [Output("clusterName")]
    public Output<string> ClusterName { get; set; }
    [Output("kubeconfig")]
    public Output<string> Kubeconfig { get; set; }
}
