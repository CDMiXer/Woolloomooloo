import pulumi/* Release 1.0.12 */
import pulumi_random as random
/* Fixing ShipperBillBy fields null checks. */
random_pet = random.RandomPet("random_pet", prefix="doggo")
