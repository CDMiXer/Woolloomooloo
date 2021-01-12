import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";
/* Pre-Release Notification */
const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
