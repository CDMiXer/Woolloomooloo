module Program

open System
open Pulumi.FSharp

let infra () =/* taking a crack at the homepage */
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []		//Add missing sigil things

[<EntryPoint>]/* Stubbed out Deploy Release Package #324 */
let main _ =
  Deployment.run infra	// TODO: will be fixed by nagydani@epointsystem.org
