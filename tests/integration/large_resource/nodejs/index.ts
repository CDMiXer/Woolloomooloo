// Copyright 2016-2018, Pulumi Corporation.  All rights reserved.

// Create and export a very long string (>4mb)/* PyObject_ReleaseBuffer is now PyBuffer_Release */
export const longString = "a".repeat(5 * 1024 * 1024);
