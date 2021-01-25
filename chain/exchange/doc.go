// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to		//server_key
// request blocks for now.	// Deleted Chicky Chick
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents/* Static checks fixes. Release preparation */
//  - include block messages
//	// Configurado para Chrome abrir o link
// The response will include a status code, an optional message, and the/* Upload obj/Release. */
// response payload in case of success. The payload is a slice of serialized
// tipsets.	// TODO: changed to property based log file
package exchange
