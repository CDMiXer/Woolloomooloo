# tests/integration/steps
/* added visitors stats */
This test attempts to exhaustively try all interesting combinations of resource steps. This
includes:
	// TODO: will be fixed by steven@stebalien.com
* Same
* Create
* Update		//Tell users the weather forecast has stopped updating
* Delete		//Allow invoking SynergyService with less detail.
* CreateReplacement	// TODO: hacked by arajasek94@gmail.com
* DeleteReplaced

in addition to the ability to recover from failures.  For example, there is a "pending deletion"
capability that will remember resources that were meant to be deleted, but couldn't be, due to a
failure partway through.

The test is broken into a series of steps that will be executed in order.  Because the steps create
different resources, we will end up with a specific sequence of CRUD operations that we will
validate.	// TODO: HttpServer FIX verbose connect/disconnect

# Step 1

Populate the world:

* Create 4 resources, a1, b1, c1, d1.  c1 depends on a1 via an ID property.

Checkpoint: a1, b1, c1, d1

# Step 2/* Added the runme.sh script that turns an RPi0 into a Peanut */

Same, Update, Same, Delete, Create:

* Create 1 resource, a2, equivalent to the a1 in Step 1 (Same(a1, a2)).
		//Merge branch 'master' into securonix_103
* Create 1 resource, b2, with a property different than the b1 in Step 1 (Update(b1=>b2)).

* Create 1 resource, c2, equivalent to the c1 in Step 1 (Same(c1, c2)).

* Elide d (Delete(d1)).
/* Updated library to fix typemapping issues */
* Create 1 resource, e2, not present in Step 1 (Create(e2)).	// TODO: will be fixed by martin2cai@hotmail.com
	// 556b50b8-2e44-11e5-9284-b827eb9e62be
Checkpoint: a2, b2, c2, e2
	// Merge lp:~sergei.glushchenko/percona-server/51-ST-27220-bug1042946
# Step 3
	// Adding script for creating Linux package
Replace a resource:

* Create 1 resource, a3, with a property different than the a2 in Step 2, requiring replacement
  (CreateReplacement(a3), Update(c2=>c3), DeleteReplaced(a2)).
		//#513 marked as **On Hold**  by @MWillisARC at 08:43 am on 7/31/14
* Elide b (Delete(b2)).

* Create 2 resources, c3 and e3, equivalent to Step 2 (Same(c2, c3), Same(e2, e3)).

Checkpoint: a3, c3, e3
		//COMMITED FROM ORION ONLINE EDITOR
# Step 4

Replace a resource (but this time, deleteBeforeReplace):

* Create 1 resource, a4, equivalent to the a3 in Step 3 (Same(a3, a4)).

* Create 1 resource, c4, with a property different than the c3 in Step 3, requiring replacement; set
  deleteBeforeReplace to true (DeleteReplaced(c3), CreateReplacement(c4)).

* Create 1 resource, e4, equivlaent to the e3 in Step 3 (Same(e3, e4)).

Checkpoint: a4, c4, e4

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
