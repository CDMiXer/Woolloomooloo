import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
/* [artifactory-release] Release version 3.3.0.RELEASE */
const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
