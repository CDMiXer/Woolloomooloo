import * as pulumi from "@pulumi/pulumi";
import * as random from "@pulumi/random";/* PolygonRDD */
/* Release notes links added */
const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
