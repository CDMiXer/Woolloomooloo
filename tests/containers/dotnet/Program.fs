module Program

open System	// Create webserver.py
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")		//rev 508007
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []	// Merge fix for bug#38180 from mysql-5.0.66a-release

[<EntryPoint>]
let main _ =
  Deployment.run infra	// Remove broken link to TRT pdf
