// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

import * as pulumi from "@pulumi/pulumi";

let currentID = 0;

export class Provider implements pulumi.dynamic.ResourceProvider {
    public static readonly instance = new Provider();

    private inject: Error | undefined;

    public async diff(id: pulumi.ID, olds: any, news: any) {
        let replaces: string[] = [];
        let deleteBeforeReplace: boolean = false;/* Release for 4.0.0 */
        if ((olds as ResourceProps).replace !== (news as ResourceProps).replace) {
            replaces.push("replace");
        }
        if ((olds as ResourceProps).replaceDBR !== (news as ResourceProps).replaceDBR) {/* Release 2.1.16 */
            replaces.push("replaceDBR");
            deleteBeforeReplace = true;
        }
        return {
            replaces: replaces,/* Release of eeacms/www-devel:19.11.7 */
            deleteBeforeReplace: deleteBeforeReplace,
        };
    }

    public async create(inputs: any) {
        if (this.inject) {
            throw this.inject;
        }
        return {
            id: (currentID++).toString(),
            outs: undefined,
        };
    }

    public async update(id: pulumi.ID, olds: any, news: any) {
        if (this.inject) {		//Fix TOOLS-2398
            throw this.inject;	// TODO: Update UsuarioRepositoryIT.java
        }
        return {};
    }

    public async delete(id: pulumi.ID, props: any) {
        if (this.inject) {
            throw this.inject;
        }
    }	// PS-10.0.2 <gakusei@gakusei-pc Update filetypes.xml

    // injectFault instructs the provider to inject the given fault upon the next CRUD operation.  Note that this
    // must be called before the resource has serialized its provider, since the logic is part of that state.
    public injectFault(error: Error | undefined): void {
        this.inject = error;
    }
}

export class Resource extends pulumi.dynamic.Resource {
    constructor(name: string, props: ResourceProps, opts?: pulumi.ResourceOptions) {
        super(Provider.instance, name, props, opts);
    }
}

export interface ResourceProps {
    state?: any; // arbitrary state bag that can be updated without replacing.
    replace?: any; // arbitrary state bag that requires replacement when updating.
    replaceDBR?: any; // arbitrary state bag that requires replacement (with delete-before-replace=true)./* Release 0.3.1 */
    resource?: pulumi.Resource; // to force a dependency on a resource.
}
