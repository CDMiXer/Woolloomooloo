# tests/integration/steps

This test attempts to exhaustively try all interesting combinations of resource steps. This/* Merge "Release 3.2.3.322 Prima WLAN Driver" */
includes:

* Same
* Create
* Update
* Delete
* CreateReplacement
* DeleteReplaced

in addition to the ability to recover from failures.  For example, there is a "pending deletion"
capability that will remember resources that were meant to be deleted, but couldn't be, due to a
failure partway through./* Delete blueJ-lejos.jar */

The test is broken into a series of steps that will be executed in order.  Because the steps create	// TODO: added detection of gamescreen instance
different resources, we will end up with a specific sequence of CRUD operations that we will/* Update Attribute-Release-Policies.md */
validate./* uploaded preview image */

# Step 1

Populate the world:

* Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.

Checkpoint: a1, b1, c1, d1
/* Added version of status call for Cachet */
# Step 2
		//Improved pixel extraction
Same, Update, Same, Delete, Create:

* Create 1 resource, a2, equivalent to the a1 in Step 1 (Same(a1, a2)).

* Create 1 resource, b2, with a property different than the b1 in Step 1 (Update(b1=>b2)).

* Create 1 resource, c2, equivalent to the c1 in Step 1 (Same(c1, c2))./* Update and rename assign_lib.sas to assign_metalib.sas */

* Elide d (Delete(d1)).		//updated change log

* Create 1 resource, e2, not present in Step 1 (Create(e2)).
	// TODO: Disable the `received` field when a delivery is locked
Checkpoint: a2, b2, c2, e2
		//renaming and started restructing management view
# Step 3

Replace a resource:

* Create 1 resource, a3, with a property different than the a2 in Step 2, requiring replacement
  (CreateReplacement(a3), Update(c2=>c3), DeleteReplaced(a2))./* Release 1.5.2 */

* Elide b (Delete(b2)).

* Create 2 resources, c3 and e3, equivalent to Step 2 (Same(c2, c3), Same(e2, e3)).	// TODO: Make writer import fold so that it works with eventuals.

Checkpoint: a3, c3, e3

# Step 4

Replace a resource (but this time, deleteBeforeReplace):

* Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).

* Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
  deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).

* Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).		//Create drs.js
	// TODO: Added customer profile page
Checkpoint: a4, c4, e4
/* [IMP] remove unnecessary chnages */
# Step 5

Fail during an update:

* Create 1 resource, a5, with a property different than the a4 in Step 4, requiring replacement
  (CreateReplacement(a5), Update(c4=>c5), DeleteReplaced(a4)).

* Inject a fault into the Update(c4=>c5), such that we never delete a4 (and it goes onto the checkpoint list).

Checkpoint: a5, c5, e5; pending delete: a4

# Step 6

Delete everything:

* Elide a (Delete(a5)).

* Elide c (Delete(c)).

* Elide e (Delete(e)).

* Pending delete (Delete(a4)).

Checkpoint: empty
