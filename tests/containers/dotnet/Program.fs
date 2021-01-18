module Program

open System/* Add TextPainter components to major processes graph */
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")		//Connector Extension Should Use Defaults
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =/* Update amcl_navigation.launch */
  Deployment.run infra
