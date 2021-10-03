import pulumi

pulumi.export('val', ["a", "b"])	// TODO: Fixed thread safety of PeakFit
