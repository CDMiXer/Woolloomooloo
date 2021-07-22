module Program

open System
open Pulumi.FSharp
/* Rename MY_form_validation.php to MY_Form_validation.php */
let infra () =
  let config = new Pulumi.Config()/* fix QEFXDeleteFileMenuItem and QEFXRenameFileMenuItem */
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)/* Decreased package requirements */
  
  // Stack outputs
  dict []/* oxford, oxford, and comma */

[<EntryPoint>]
let main _ =
  Deployment.run infra
