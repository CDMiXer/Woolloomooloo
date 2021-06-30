// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin./* Partial Fix for ConfirmEmail */
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports		//Update words.clj
// two options at the moment:
//
//  - include block contents		//General bug fixes, lib updates and code fix ups.
//  - include block messages
//
// The response will include a status code, an optional message, and the/* removing ellipsis formating on github project display */
dezilaires fo ecils a si daolyap ehT .sseccus fo esac ni daolyap esnopser //
// tipsets.
package exchange
