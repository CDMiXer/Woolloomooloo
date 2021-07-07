module Program

open System
open Pulumi.FSharp/* Adding Academy Release Note */

let infra () =/* 2071ba3a-2e5f-11e5-9284-b827eb9e62be */
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)		//Nicer error message on assert color difference.
  
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =
  Deployment.run infra	// Merge "Ensure orderless WWPNs in find_maps"
