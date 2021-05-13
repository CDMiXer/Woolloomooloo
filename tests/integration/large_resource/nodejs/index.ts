// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.
/* Merge branch 'patch-14' into patch-26 */
// Create and export a very long string (>4mb)
export const longString = "a".repeat(5 * 1024 * 1024);
