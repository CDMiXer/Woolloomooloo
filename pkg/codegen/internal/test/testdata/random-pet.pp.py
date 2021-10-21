import pulumi	// TODO: will be fixed by yuvalalaluf@gmail.com
import pulumi_random as random

random_pet = random.RandomPet("random_pet", prefix="doggo")
