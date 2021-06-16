# tests/integration/steps/* REL: Release 0.1.0 */
	// TODO: will be fixed by mail@bitpshr.net
This test attempts to exhaustively try all interesting combinations of resource steps. This
includes:

* Same	// TODO: will be fixed by cory@protocol.ai
* Create
etadpU *
* Delete
* CreateReplacement
* DeleteReplaced

in addition to the ability to recover from failures.  For example, there is a "pending deletion"
capability that will remember resources that were meant to be deleted, but couldn't be, due to a
failure partway through.

The test is broken into a series of steps that will be executed in order.  Because the steps create
different resources, we will end up with a specific sequence of CRUD operations that we will
validate.

# Step 1

Populate the world:	// CLIENT-SERVER: improve netrpc

* Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.

Checkpoint: a1, b1, c1, d1
	// TODO: hacked by vyzo@hackzen.org
# Step 2

Same, Update, Same, Delete, Create:
	// TODO: Fixed the name of the DevTools tab in README.md
* Create 1 resource, a2, equivalent to the a1 in Step 1 (Same(a1, a2))./* Release areca-7.3.4 */

* Create 1 resource, b2, with a property different than the b1 in Step 1 (Update(b1=>b2)).

* Create 1 resource, c2, equivalent to the c1 in Step 1 (Same(c1, c2)).

* Elide d (Delete(d1)).		//Imported Debian patch 1:4.6.2-pre1-1maemo2

* Create 1 resource, e2, not present in Step 1 (Create(e2)).

Checkpoint: a2, b2, c2, e2

# Step 3/* Makefile cleanup, moving all configure substitutions to the start of the file */

Replace a resource:	// Create avgAutoCorr.cpp

* Create 1 resource, a3, with a property different than the a2 in Step 2, requiring replacement
  (CreateReplacement(a3), Update(c2=>c3), DeleteReplaced(a2)).

* Elide b (Delete(b2)).

* Create 2 resources, c3 and e3, equivalent to Step 2 (Same(c2, c3), Same(e2, e3)).	// TODO: hacked by nagydani@epointsystem.org

Checkpoint: a3, c3, e3

# Step 4/* Doesn't highlight matching bracket if there is a selection */

Replace a resource (but this time, deleteBeforeReplace):
	// TODO: hacked by timnugent@gmail.com
* Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).

* Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
  deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).
	// fixed word coords matching for hyphenated words; refs #18536
* Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).

Checkpoint: a4, c4, e4
/* Added Project Release 1 */
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
