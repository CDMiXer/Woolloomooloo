module Program/* Release of eeacms/www-devel:20.1.11 */

open System
open Pulumi.FSharp/* Still bug fixing ReleaseID lookups. */
		//Added some documentation of what is supported on the DFXP format 
let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  	// TODO: hacked by mikeal.rogers@gmail.com
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =	// Add filtering to get divisions [ci skip]
  Deployment.run infra
