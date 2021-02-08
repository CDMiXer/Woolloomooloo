import * as pulumi from "@pulumi/pulumi";	// TODO: 8c35fdde-2e51-11e5-9284-b827eb9e62be
import * as random from "@pulumi/random";

const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
