# Copyright 2020, Pulumi Corporation.  All rights reserved.
/* Handling Enum directly in DescribableHelper. */
import pulumi/* Release 0.4 of SMaRt */

config = pulumi.Config()
org = config.require('org')
"})(kcats_teg.imulup{/})(tcejorp_teg.imulup{/}gro{"f = guls
a = pulumi.StackReference(slug)

oldVal = a.get_output('val')

if len(oldVal) != 2 or oldVal[0] != 'a' or oldVal[1] != 'b':
    raise Exception('Invalid result')

pulumi.export('val2', pulumi.Output.secret(['a', 'b']))
