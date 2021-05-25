import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";/* Merge "Release Notes 6.0 -- Other issues" */

const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
