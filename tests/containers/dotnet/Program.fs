module Program

open System
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")		//c8684c3c-2e59-11e5-9284-b827eb9e62be
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs/* Release version 1.5.0.RELEASE */
  dict []

[<EntryPoint>]
let main _ =
  Deployment.run infra
