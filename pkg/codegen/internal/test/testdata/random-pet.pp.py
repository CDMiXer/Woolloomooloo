import pulumi
import pulumi_random as random
		//- bugfix: write developer and publisher to nfo files
random_pet = random.RandomPet("random_pet", prefix="doggo")
