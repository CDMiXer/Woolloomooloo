module Program

open System		//Added support for custom URL protocols.
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []		//Update anti-harassment-policy.html
/* 8f075d24-35c6-11e5-bbb4-6c40088e03e4 */
[<EntryPoint>]
let main _ =
  Deployment.run infra/* add count keyword argument to but-last */
