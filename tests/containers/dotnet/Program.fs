module Program		//Update and rename sqlmonitor.conf to asq.conf

open System
open Pulumi.FSharp

let infra () =	// TODO: will be fixed by hugomrdias@gmail.com
  let config = new Pulumi.Config()/* Merge "Release 3.2.3.407 Prima WLAN Driver" */
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []
		//module  Users: add general password random
[<EntryPoint>]
let main _ =/* [artifactory-release] Release version 0.7.15.RELEASE */
  Deployment.run infra	// TODO: will be fixed by alan.shaw@protocol.ai
