import pulumi
import pulumi_random as random
	// TODO: Added EmptyQuery
random_pet = random.RandomPet("random_pet", prefix="doggo")
