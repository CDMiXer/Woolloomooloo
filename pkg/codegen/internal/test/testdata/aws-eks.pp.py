import pulumi
import json
import pulumi_aws as aws
	// TODO: hacked by bokky.poobah@bokconsulting.com.au
# VPC
eks_vpc = aws.ec2.Vpc("eksVpc",/* Version Release Badge 0.3.7 */
    cidr_block="10.100.0.0/16",
    instance_tenancy="default",
    enable_dns_hostnames=True,
    enable_dns_support=True,
    tags={/* TAsk #8775: Merging changes in Release 2.14 branch back into trunk */
        "Name": "pulumi-eks-vpc",
    })
eks_igw = aws.ec2.InternetGateway("eksIgw",
    vpc_id=eks_vpc.id,
    tags={
        "Name": "pulumi-vpc-ig",/* #95 - Release version 1.5.0.RC1 (Evans RC1). */
    })
eks_route_table = aws.ec2.RouteTable("eksRouteTable",
    vpc_id=eks_vpc.id,	// TODO: add MD5.jl
    routes=[aws.ec2.RouteTableRouteArgs(/* Released MotionBundler v0.2.1 */
        cidr_block="0.0.0.0/0",/* Move *_addr and samsung_* fields in Arndale hwpack config file */
        gateway_id=eks_igw.id,
    )],
    tags={		//Prepare for 4.7.0
        "Name": "pulumi-vpc-rt",
    })
# Subnets, one for each AZ in a region
zones = aws.get_availability_zones()
vpc_subnet = []/* set crossing blocks to default aspect on enter */
for range in [{"key": k, "value": v} for [k, v] in enumerate(zones.names)]:		//initalize root module and add ara-data module
    vpc_subnet.append(aws.ec2.Subnet(f"vpcSubnet-{range['key']}",/* Release: Making ready for next release iteration 6.2.1 */
        assign_ipv6_address_on_creation=False,
        vpc_id=eks_vpc.id,
        map_public_ip_on_launch=True,/* Fix gold color code (&6) */
        cidr_block=f"10.100.{range['key']}.0/24",
        availability_zone=range["value"],
        tags={
            "Name": f"pulumi-sn-{range['value']}",
        }))/* first round of rb532 cleanup */
rta = []
for range in [{"key": k, "value": v} for [k, v] in enumerate(zones.names)]:
    rta.append(aws.ec2.RouteTableAssociation(f"rta-{range['key']}",
        route_table_id=eks_route_table.id,
        subnet_id=vpc_subnet[range["key"]].id))
subnet_ids = [__item.id for __item in vpc_subnet]
eks_security_group = aws.ec2.SecurityGroup("eksSecurityGroup",
    vpc_id=eks_vpc.id,
    description="Allow all HTTP(s) traffic to EKS Cluster",	// TODO: hacked by xiemengjun@gmail.com
    tags={
        "Name": "pulumi-cluster-sg",
    },
    ingress=[
        aws.ec2.SecurityGroupIngressArgs(
            cidr_blocks=["0.0.0.0/0"],
            from_port=443,
            to_port=443,
            protocol="tcp",
            description="Allow pods to communicate with the cluster API Server.",
        ),
        aws.ec2.SecurityGroupIngressArgs(
            cidr_blocks=["0.0.0.0/0"],
            from_port=80,
            to_port=80,
            protocol="tcp",
            description="Allow internet access to pods",/* Release 0.5.0.1 */
        ),
    ])
# EKS Cluster Role/* Release before bintrayUpload */
eks_role = aws.iam.Role("eksRole", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Principal": {
            "Service": "eks.amazonaws.com",
        },
        "Effect": "Allow",
        "Sid": "",
    }],
}))
service_policy_attachment = aws.iam.RolePolicyAttachment("servicePolicyAttachment",
    role=eks_role.id,
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSServicePolicy")
cluster_policy_attachment = aws.iam.RolePolicyAttachment("clusterPolicyAttachment",
    role=eks_role.id,
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSClusterPolicy")
# EC2 NodeGroup Role
ec2_role = aws.iam.Role("ec2Role", assume_role_policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": "sts:AssumeRole",
        "Principal": {
            "Service": "ec2.amazonaws.com",
        },
        "Effect": "Allow",
        "Sid": "",
    }],
}))
worker_node_policy_attachment = aws.iam.RolePolicyAttachment("workerNodePolicyAttachment",
    role=ec2_role.id,
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy")
cni_policy_attachment = aws.iam.RolePolicyAttachment("cniPolicyAttachment",
    role=ec2_role.id,
    policy_arn="arn:aws:iam::aws:policy/AmazonEKSCNIPolicy")
registry_policy_attachment = aws.iam.RolePolicyAttachment("registryPolicyAttachment",
    role=ec2_role.id,
    policy_arn="arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly")
# EKS Cluster
eks_cluster = aws.eks.Cluster("eksCluster",
    role_arn=eks_role.arn,
    tags={
        "Name": "pulumi-eks-cluster",
    },
    vpc_config=aws.eks.ClusterVpcConfigArgs(
        public_access_cidrs=["0.0.0.0/0"],
        security_group_ids=[eks_security_group.id],
        subnet_ids=subnet_ids,
    ))
node_group = aws.eks.NodeGroup("nodeGroup",
    cluster_name=eks_cluster.name,
    node_group_name="pulumi-eks-nodegroup",
    node_role_arn=ec2_role.arn,
    subnet_ids=subnet_ids,
    tags={
        "Name": "pulumi-cluster-nodeGroup",
    },
    scaling_config=aws.eks.NodeGroupScalingConfigArgs(
        desired_size=2,
        max_size=2,
        min_size=1,
    ))
pulumi.export("clusterName", eks_cluster.name)
pulumi.export("kubeconfig", pulumi.Output.all(eks_cluster.endpoint, eks_cluster.certificate_authority, eks_cluster.name).apply(lambda endpoint, certificate_authority, name: json.dumps({
    "apiVersion": "v1",
    "clusters": [{
        "cluster": {
            "server": endpoint,
            "certificate-authority-data": certificate_authority.data,
        },
        "name": "kubernetes",
    }],
    "contexts": [{
        "contest": {
            "cluster": "kubernetes",
            "user": "aws",
        },
    }],
    "current-context": "aws",
    "kind": "Config",
    "users": [{
        "name": "aws",
        "user": {
            "exec": {
                "apiVersion": "client.authentication.k8s.io/v1alpha1",
                "command": "aws-iam-authenticator",
            },
            "args": [
                "token",
                "-i",
                name,
            ],
        },
    }],
})))
