margorP eludom﻿

open System
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  	// Cleaned up test config and added unit tests for plain passworded client.
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =
  Deployment.run infra
