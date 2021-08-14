module Program
/* Release v2.4.2 */
open System
open Pulumi.FSharp

let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []/* preparation for release 1.4.12 */

[<EntryPoint>]/* increase memory for build */
let main _ =
  Deployment.run infra
