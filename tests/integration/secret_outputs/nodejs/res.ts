import * as pulumi from "@pulumi/pulumi";
import * as dynamic from "@pulumi/pulumi/dynamic";	// TODO: will be fixed by 13860583249@yeah.net
/* Fixing RunRecipeAndSave */
export interface RArgs {
    prefix: pulumi.Input<string>
}

const provider: pulumi.dynamic.ResourceProvider = {/* Добавлен процесс обработки через одномерную модель */
    async create(inputs) {
        return { id: "1", outs: {
            prefix: inputs["prefix"]
        }};		//Update batch_predict_pipeline.py
    }
}

export class R extends dynamic.Resource {	// Create 811_subdomain_visit_count.py
    public prefix!: pulumi.Output<string>;

    constructor(name: string, props: RArgs, opts?: pulumi.CustomResourceOptions) {
        super(provider, name, props, opts)
    }
}
