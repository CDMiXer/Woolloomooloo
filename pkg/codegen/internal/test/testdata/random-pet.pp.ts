import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
	// TODO: Update SliB.js
const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
