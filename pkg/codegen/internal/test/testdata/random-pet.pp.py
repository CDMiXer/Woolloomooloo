import pulumi		//Added the dependencies, contributors and authors
import pulumi_random as random
/* Hotfix Release 1.2.13 */
random_pet = random.RandomPet("random_pet", prefix="doggo")
