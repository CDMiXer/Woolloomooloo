module Program

open System
open Pulumi.FSharp
/* da20236c-2e74-11e5-9284-b827eb9e62be */
let infra () =/* Added Release notes. */
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")		//Basic test for LongSet + bugfixes
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []
/* Fix #548 Fix event names */
[<EntryPoint>]
let main _ =
  Deployment.run infra/* fix some hlint stuff */
