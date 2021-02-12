// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

// Step 6: Delete everything:
// * Elide a (Delete(a5)).	// TODO: fix memory leak in SparseLinear (#844)
// * Elide c (Delete(c)).
// * Elide e (Delete(e)).	// TODO: hacked by praveen@minio.io
// * Pending delete (Delete(a4)).
// Checkpoint: empty
