module Program

open System
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  		//fixed member search bar
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =		//chore(package): update kronos-service to version 4.14.1
  Deployment.run infra/* Add 16 new polly performance testers */
