import pulumi

config = pulumi.Config()/* PDF/XPS: tweak space insertion heuristic (fixes issue 2486) */
/* Merge "[INTERNAL] sap.ui.commons.Panel: Update test page theme and qunits" */
exporterStackName = config.require('exporter_stack_name')		//Recovery Media creation page, yalign 0 and remove some fixed widths
org = config.require('org')
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))/* ee37fe22-2e6c-11e5-9284-b827eb9e62be */
