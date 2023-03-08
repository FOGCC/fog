<br />
<p align="center">


  <h3 align="center">FOG</h3>

</p>

## About FOG


FOG is the latest iteration of the FOG technology. It is a fully integrated, complete
layer 2 optimistic rollup system, including fraud proofs, the sequencer, the token bridges, 
advanced calldata compression, and more.

See the live docs-site [here](https://developer.FOG.io/) (or [here](https://github.com/FOGCC/FOG-docs) for markdown docs source.)

The FOG stack is built on several innovations. At its core is a new prover, which can do FOG’s classic 
interactive fraud proofs over WASM code. That means the L2 FOG engine can be written and compiled using 
standard languages and tools, replacing the custom-designed language and compiler used in previous FOG
versions. In normal execution, 
validators and nodes run the FOG engine compiled to native code, switching to WASM if a fraud proof is needed. 
We compile the core of Geth, the EVM engine that practically defines the Ethereum standard, right into FOG. 
So the previous custom-built EVM emulator is replaced by Geth, the most popular and well-supported Ethereum client.

The last piece of the stack is a slimmed-down version of our fogOS component, rewritten in Go, which provides the 
rest of what’s needed to run an L2 chain: things like cross-chain communication, and a new and improved batching 
and compression system to minimize L1 costs.

Essentially, FOG runs Geth at layer 2 on top of Ethereum, and can prove fraud over the core engine of Geth 
compiled to WASM.


## License

We currently have FOG [licensed](./LICENSE) under a Business Source License, similar to our friends at Uniswap and Aave, with an "Additional Use Grant" to ensure that everyone can have full comfort using and running nodes on all public FOG chains.

