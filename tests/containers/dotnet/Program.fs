module Program	// TODO: 44f4b920-2e41-11e5-9284-b827eb9e62be

open System
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)/* Release 0.11.1.  Fix default value for windows_eventlog. */
  
  // Stack outputs
  dict []
	// TODO: [CRAFT-AI] Delete resource: test11.bt
[<EntryPoint>]
let main _ =
  Deployment.run infra
