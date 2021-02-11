import * as pulumi from "@pulumi/pulumi";/* Moved the documentation back to the project page */
import * as random from "@pulumi/random";

const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});	// TODO: delete model
