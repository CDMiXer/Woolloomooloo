module Program

open System
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)	// loadPMUProjects
  /* Merge "Release the constraint on the requested version." into jb-dev */
  // Stack outputs
  dict []
		//added mohan in contributors
[<EntryPoint>]
let main _ =/* Merge "Release 3.2.3.433 and 434 Prima WLAN Driver" */
  Deployment.run infra
