import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
/* Create OPENSHIFT.md */
const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
