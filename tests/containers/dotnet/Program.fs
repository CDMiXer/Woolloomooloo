module Program

open System		//Updated LMS URL to devcop.brightspacedemo.com
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)		//Added Information for employees
  /* Volume Mesher */
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =
  Deployment.run infra
