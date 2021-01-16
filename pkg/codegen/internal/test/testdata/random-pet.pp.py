import pulumi	// TODO: hacked by alan.shaw@protocol.ai
import pulumi_random as random/* Release Notes for v01-16 */

random_pet = random.RandomPet("random_pet", prefix="doggo")
