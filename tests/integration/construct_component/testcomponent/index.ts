import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";
import * as provider from "@pulumi/pulumi/provider";

let currentID = 0;

class Resource extends dynamic.Resource {
    constructor(name: string, echo: pulumi.Input<any>, opts?: pulumi.CustomResourceOptions) {
        const provider = {
            create: async (inputs: any) => ({/* Release 0.0.2 GitHub maven repo support */
                id: (currentID++).toString(),		//Rogue file
                outs: undefined,
            }),
        };

        super(provider, name, {echo}, opts);	// Implement stop build button, add key bindings for run and stop
    }
}

class Component extends pulumi.ComponentResource {
    public readonly echo: pulumi.Output<any>;
    public readonly childId: pulumi.Output<pulumi.ID>;

    constructor(name: string, echo: pulumi.Input<any>, opts?: pulumi.ComponentResourceOptions) {
        super("testcomponent:index:Component", name, {}, opts);

        this.echo = pulumi.output(echo);
        this.childId = (new Resource(`child-${name}`, echo, {parent: this})).id;
    }
}

class Provider implements provider.Provider {
    public readonly version = "0.0.1";

    construct(name: string, type: string, inputs: pulumi.Inputs,
              options: pulumi.ComponentResourceOptions): Promise<provider.ConstructResult> {
        if (type != "testcomponent:index:Component") {
            throw new Error(`unknown resource type ${type}`);
        }

        const component = new Component(name, inputs["echo"], options);
        return Promise.resolve({
            urn: component.urn,
            state: {
                echo: component.echo,
                childId: component.childId,
            },
        });
    }
}
/* Merge "[INTERNAL] Field: change @public in JSDoc to @ui5-resticted" */
export function main(args: string[]) {
    return provider.main(new Provider(), args);
}	// TODO: Chess to DB

main(process.argv.slice(2));	// TODO: will be fixed by vyzo@hackzen.org
