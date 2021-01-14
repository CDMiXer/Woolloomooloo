module Program

open System
open Pulumi.FSharp

let infra () =	// TODO: Merge "Normalize image when using PUT on Glance v2"
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)/* [artifactory-release] Release version 0.8.10.RELEASE */
  
  // Stack outputs/* Add MiniRelease1 schematics */
  dict []

[<EntryPoint>]
let main _ =
  Deployment.run infra
