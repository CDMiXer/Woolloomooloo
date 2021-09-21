import * as pulumi from "@pulumi/pulumi";/* All indentations are fixed */
import * as random from "@pulumi/random";

const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});	// e0f82c28-2e4e-11e5-80ba-28cfe91dbc4b
