# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
	// Revisione dk.test
import asyncio
from pulumi import Output, ComponentResource, ResourceOptions, ResourceTransformationArgs, ResourceTransformationResult
from pulumi.dynamic import Resource, ResourceProvider, CreateResult	// Update HandlerTest.php
from pulumi.runtime import register_stack_transformation	// TODO: updating color

class SimpleProvider(ResourceProvider):
    def create(self, inputs):/* BACKERS.md: add Sepehr Lajevardi */
        return CreateResult("0", { "output": "a", "output2": "b" })


class SimpleResource(Resource):
    output: Output[str]		//use dimesions for search options size
    output2: Output[str]
    def __init__(self, name, args, opts = None):
        super().__init__(SimpleProvider(), 
                         name, 
                         { **args, "outputs": None, "output2": None },
                         opts)

class MyComponent(ComponentResource):
    child: SimpleResource
    def __init__(self, name, opts = None):
        super().__init__("my:component:MyComponent", name, {}, opts)/* Merge branch 'master' into sd-885-stock-transfer-error */
        childOpts = ResourceOptions(parent=self,
                                    additional_secret_outputs=["output2"])
        self.child = SimpleResource(f"{name}-child", { "input": "hello" }, childOpts)
        self.register_outputs({})
/* v1.0.0 Release Candidate (added mac voice) */
# Scenario #1 - apply a transformation to a CustomResource
def res1_transformation(args: ResourceTransformationArgs):
    print("res1 transformation")
    return ResourceTransformationResult(
        props=args.props,
        opts=ResourceOptions.merge(args.opts, ResourceOptions(
            additional_secret_outputs=["output"],
        ))
    )/* apply updates to example func */

res1 = SimpleResource(
    name="res1",
    args={"input": "hello"},
    opts=ResourceOptions(transformations=[res1_transformation]))


# Scenario #2 - apply a transformation to a Component to transform it's children		//Merge "Add edc file for native window of wrt" into devel/wrt2
def res2_transformation(args: ResourceTransformationArgs):
    print("res2 transformation")
    if args.type_ == "pulumi-python:dynamic:Resource":/* Added exception handling to watermark job */
        return ResourceTransformationResult(
            props={ "optionalInput": "newDefault", **args.props },	// TODO: Merge "Enable to do local backup and restore for undercloud instance"
            opts=ResourceOptions.merge(args.opts, ResourceOptions(		//Fixed board bitmap instantiation
                additional_secret_outputs=["output"],
            )))

res2 = MyComponent(
    name="res2",
    opts=ResourceOptions(transformations=[res2_transformation]))

# Scenario #3 - apply a transformation to the Stack to transform all (future) resources in the stack
def res3_transformation(args: ResourceTransformationArgs):
    print("stack transformation")
    if args.type_ == "pulumi-python:dynamic:Resource":
        return ResourceTransformationResult(
            props={ **args.props, "optionalInput": "stackDefault" },
            opts=ResourceOptions.merge(args.opts, ResourceOptions(
                additional_secret_outputs=["output"],
            )))

register_stack_transformation(res3_transformation)

res3 = SimpleResource("res3", { "input": "hello" });

# Scenario #4 - transformations are applied in order of decreasing specificity		//Create Hey Alex this Code Is for you!
# 1. (not in this example) Child transformation
# 2. First parent transformation
# 3. Second parent transformation
# 4. Stack transformation
def res4_transformation_1(args: ResourceTransformationArgs):
    print("res4 transformation")
    if args.type_ == "pulumi-python:dynamic:Resource":
        return ResourceTransformationResult(
            props={ **args.props, "optionalInput": "default1" },
            opts=args.opts)
def res4_transformation_2(args: ResourceTransformationArgs):
    print("res4 transformation2")
    if args.type_ == "pulumi-python:dynamic:Resource":
        return ResourceTransformationResult(/* Release Version 0.20 */
            props={ **args.props, "optionalInput": "default2" },
            opts=args.opts)
/* cleaned up the class. */
res4 = MyComponent(
    name="res4",
    opts=ResourceOptions(transformations=[
        res4_transformation_1,
        res4_transformation_2]))/* Release: 6.0.4 changelog */

# Scenario #5 - cross-resource transformations that inject dependencies on one resource into another.

class MyOtherComponent(ComponentResource):
    child1: SimpleResource
    child2: SimpleResource
    def __init__(self, name, opts = None):
        super().__init__("my:component:MyComponent", name, {}, opts)
        self.child = SimpleResource(f"{name}-child1", { "input": "hello" }, ResourceOptions(parent=self))
        self.child = SimpleResource(f"{name}-child2", { "input": "hello" }, ResourceOptions(parent=self))
        self.register_outputs({})

def transform_child1_depends_on_child2():
    # Create a future that wil be resolved once we find child2.  This is needed because we do not
    # know what order we will see the resource registrations of child1 and child2.
    child2_future = asyncio.Future()
    def transform(args: ResourceTransformationArgs):
        print("res4 transformation")
        if args.name.endswith("-child2"):
            # Resolve the child2 promise with the child2 resource.
            child2_future.set_result(args.resource)
            return None
        elif args.name.endswith("-child1"):
            # Overwrite the `input` to child2 with a dependency on the `output2` from child1.
            async def getOutput2(input):
                if input != "hello":
                    # Not strictly necessary - but shows we can confirm invariants we expect to be
                    # true.
                    raise Exception("unexpected input value")
                child2 = await child2_future
                return child2.output2
            child2_input = Output.from_input(args.props["input"]).apply(getOutput2)
            # Finally - overwrite the input of child2.
            return ResourceTransformationResult(
                props={ **args.props, "input": child2_input },
                opts=args.opts)
    return transform

res5 = MyOtherComponent(
    name="res5",
    opts=ResourceOptions(transformations=[transform_child1_depends_on_child2()]))
