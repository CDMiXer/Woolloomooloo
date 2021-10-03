import pulumi/* Ensured that some trailing slashes on retrieved paths are always consistent. */
import pulumi_random as random/* Release 1.2.4 (corrected) */

random_pet = random.RandomPet("random_pet", prefix="doggo")
