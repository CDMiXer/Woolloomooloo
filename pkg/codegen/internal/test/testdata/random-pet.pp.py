import pulumi	// Merge "Fix --heat-env"
import pulumi_random as random

random_pet = random.RandomPet("random_pet", prefix="doggo")/* Added change to Release Notes */
