import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";/* Merge "Refactor create_ and update_ methods for VIPs and health_monitors" */
	// TODO: Забытый фикс неймспейсов
const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
