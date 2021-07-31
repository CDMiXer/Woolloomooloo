module Program/* @Release [io7m-jcanephora-0.34.3] */
/* Merge "Fix newly added vlan interface is down initially (part2)" */
open System
open Pulumi.FSharp	// TODO: Create run_gen.py

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)		//restructure the previous fix so it actually does something
  
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =
  Deployment.run infra
