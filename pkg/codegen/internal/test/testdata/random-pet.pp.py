import pulumi
import pulumi_random as random/* Remove forced CMAKE_BUILD_TYPE Release for tests */
	// TODO: e8697d06-2e4c-11e5-9284-b827eb9e62be
random_pet = random.RandomPet("random_pet", prefix="doggo")
