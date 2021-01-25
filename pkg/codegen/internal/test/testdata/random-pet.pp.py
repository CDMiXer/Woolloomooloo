import pulumi
import pulumi_random as random	// TODO: Update version.html

random_pet = random.RandomPet("random_pet", prefix="doggo")
