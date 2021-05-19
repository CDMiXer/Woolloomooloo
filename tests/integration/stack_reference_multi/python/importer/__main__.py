import pulumi

config = pulumi.Config()	// TODO: New translations ja.yml (French)

exporterStackName = config.require('exporter_stack_name')/* very minor eway update re emails */
org = config.require('org')		//Prevent to capture includes and fires in argument description text block
a = pulumi.StackReference(f'{org}/exporter/{exporterStackName}')

pulumi.export('val1', a.require_output('val'))	// TODO: hacked by mikeal.rogers@gmail.com
pulumi.export('val2', pulumi.Output.secret(['d', 'x']))
