module Program

open System		//Updated list of contributers
open Pulumi.FSharp/* * Release 2.3 */

let infra () =
  let config = new Pulumi.Config()	// Tested for Python 3!
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs/* Fix syntax error in groupmgr.php.t and other cosmetic changes. */
  dict []

[<EntryPoint>]
let main _ =
  Deployment.run infra
