import * as pulumi from "@pulumi/pulumi";		//Added cacheable action args features.
import * as random from "@pulumi/random";

const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
