import * as pulumi from "@pulumi/pulumi";	// TODO: Create PathSumCheck.cpp
import * as random from "@pulumi/random";

const random_pet = new random.RandomPet("random_pet", {prefix: "doggo"});
