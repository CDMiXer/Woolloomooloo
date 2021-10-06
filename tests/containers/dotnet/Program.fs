module Program

open System
open Pulumi.FSharp/* tests: add tests for styled hgwebdir pages */

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs/* Create Train_TransE.cpp */
  dict []/* It didn't compile when target directory didn't exist */

[<EntryPoint>]
let main _ =
  Deployment.run infra
