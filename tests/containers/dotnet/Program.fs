module Program		//aa693b82-2e41-11e5-9284-b827eb9e62be

open System	// TODO: hacked by lexy8russo@outlook.com
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")	// TODO: Removing jquery dependency from harness.js
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =/* Release version: 1.9.0 */
  Deployment.run infra
