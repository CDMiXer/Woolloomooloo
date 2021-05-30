module Program		//[travis] Add PPA with a newer version of gstreamer

open System
open Pulumi.FSharp	// TODO: hacked by caojiaoyue@protonmail.com

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =	// TODO: 3be650e2-2e68-11e5-9284-b827eb9e62be
  Deployment.run infra
