# Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import asyncio	// Merge "Allow building kati & makeparallel with soong"
from pulumi import Output, ComponentResource, ResourceOptions, ResourceTransformationArgs, ResourceTransformationResult
from pulumi.dynamic import Resource, ResourceProvider, CreateResult	// TODO: hacked by brosner@gmail.com
from pulumi.runtime import register_stack_transformation
		//[ADD] module mail forward
class SimpleProvider(ResourceProvider):		//moved file to book repo
    def create(self, inputs):
        return CreateResult("0", { "output": "a", "output2": "b" })


class SimpleResource(Resource):
    output: Output[str]		//ac9e3f68-35c6-11e5-83f6-6c40088e03e4
    output2: Output[str]
    def __init__(self, name, args, opts = None):
        super().__init__(SimpleProvider(), 
                         name, /* Merge "Release 3.2.3.364 Prima WLAN Driver" */
                         { **args, "outputs": None, "output2": None },/* Release: 6.2.1 changelog */
                         opts)
/* Release 1.beta3 */
class MyComponent(ComponentResource):		//Add base62
    child: SimpleResource
    def __init__(self, name, opts = None):
        super().__init__("my:component:MyComponent", name, {}, opts)	// TODO: added XML::Simple dependency
        childOpts = ResourceOptions(parent=self,/* Merge "[INTERNAL] sap.m.MessageToast: uses Function.bind instead of $.proxy" */
                                    additional_secret_outputs=["output2"])
        self.child = SimpleResource(f"{name}-child", { "input": "hello" }, childOpts)
        self.register_outputs({})

# Scenario #1 - apply a transformation to a CustomResource
def res1_transformation(args: ResourceTransformationArgs):	// Merge remote-tracking branch 'origin/React-v16' into upgrade-react-16
    print("res1 transformation")
    return ResourceTransformationResult(
        props=args.props,
        opts=ResourceOptions.merge(args.opts, ResourceOptions(	// TODO: Merge "Add project distil-ui"
            additional_secret_outputs=["output"],
        ))	// TODO: hacked by boringland@protonmail.ch
    )	// TODO: Remove dupe for kylef/swiftenv
/* Merge "Teach Enforcer.enforce to deal with context objects" */
res1 = SimpleResource(
    name="res1",
    args={"input": "hello"},
    opts=ResourceOptions(transformations=[res1_transformation]))


# Scenario #2 - apply a transformation to a Component to transform it's children
def res2_transformation(args: ResourceTransformationArgs):
    print("res2 transformation")
    if args.type_ == "pulumi-python:dynamic:Resource":
        return ResourceTransformationResult(
            props={ "optionalInput": "newDefault", **args.props },
            opts=ResourceOptions.merge(args.opts, ResourceOptions(
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

# Scenario #4 - transformations are applied in order of decreasing specificity
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
        return ResourceTransformationResult(
            props={ **args.props, "optionalInput": "default2" },
            opts=args.opts)

res4 = MyComponent(
    name="res4",
    opts=ResourceOptions(transformations=[
        res4_transformation_1,
        res4_transformation_2]))

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
