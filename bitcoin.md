# Bitcoin

Bitcoin is an experimental new digital currency that enables instant payments to anyone, anywhere in the world. Bitcoin uses peer-to-peer technology to operate with no central authority: managing transactions and issuing money are carried out collectively by the network. 

[Bitcoin Core](https://github.com/bitcoin/bitcoin) is the name of open source software which enables the use of this currency.

**Introduction**

- A purely peer-to-peer version of electronic cash would allow online payments to be sent directly from one party to another without going through a financial institution.
	- Digital signatures provide part of the solution
		- To prove authenticity: RSA algorithm, etc
	- But the main benefits are lost if a trusted third party is still required to prevent double-spending.
		- Paypal, Citibank, etc.
- They proposed a solution to the double-spending problem using a peer-to-peer network.
	- The network timestamps transactions by hashing them into an ongoing chain of hash-based proof-of-work, forming a record that cannot be changed without redoing the proof-of-work.
	-  The longest chain not only serves as proof of the sequence of events witnessed, but proof that it came from the largest pool of CPU power.
		- Every 10 minutes, the state of the network achieves consensus on a block, and the transactions up to that point are verified and agreed to at this point. As consensus is achieved on a block, it strengthens the security of the previous block, building the security of the network as the network goes along.
			- It takes $400 million worth of computing resources to take over a single network for 10 minutes.
		- As long as a majority of CPU power is controlled by nodes that are not cooperating to attack the network, they'll generate the longest chain and outpace attackers.
- The network itself requires minimal structure. Messages are broadcast on a best effort basis, and nodes can leave and rejoin the network at will, accepting the longest proof-of-work chain as proof of what happened while they were gone.

**Current to Proposed Solution**

- Commerce on the Internet has come to rely almost exclusively on financial institutions serving as trusted third parties to process electronic payments.
	- Existing payment networks were not built for the Internet.
		- While the system works well enough for most transactions, it still suffers from the inherent weaknesses of the trust based model.
	- The Bitcoin protocol will disrupt the financial industry by making it significantly cheaper to transfer ownership worldwide.
- What is needed is an electronic payment system based on cryptographic proof instead of trust, allowing any two willing parties to transact directly with each other without the need for a trusted third party. 
	- The system is secure as long as honest nodes collectively control more CPU power than any cooperating group of attacker nodes.
	- Transactions that are computationally impractical to reverse would protect sellers from fraud, and routine escrow mechanisms could easily be implemented to protect buyers.

**Bitcoin**

- We define an electronic coin as a chain of digital signatures.
	- Digital signatures consist of a public and a private key.
	- If your bank account’s routing number is the public key (enables incoming wires), your PIN is the private key (enables outgoing wires).
	- If your written down physical signature is the public key, the procedural knowledge of your hand’s movements is the private key.
- Each owner transfers the coin to the next by:
	1. Digitally signing a hash of the previous transaction,
	2. Digitally signing public key of the next owner,
	3. Adding these to the end of the coin.
	4. A payee can verify the signatures to verify the chain of ownership.

<img src="http://images.rapgenius.com/47dfc5dde83036b092872542a62290ad.1000x530x1.png" alt="Drawing" style="width: 400px;"/>

- The problem of course is the payee can't verify that one of the owners did not double-spend the coin. We need a way for the payee to know that the previous owners did not sign any earlier transactions.
	- For our purposes, the earliest transaction is the one that counts, so we don't care about later attempts to double-spend.
	- The only way to confirm the absence of a transaction is to be aware of all transactions.
- To accomplish this without a trusted party, transactions must be publicly announced
	- We need a system for participants to agree on a single history of the order in which they were received. 
	- The payee needs proof that at the time of each transaction, the majority of nodes agreed it was the first received.

**Timestamp Server**

- In order to verify that individuals don't double-spend the coin, we need to establish a system. The solution we propose begins with a timestamp server.
-  A timestamp server works by taking a hash of a block of items to be timestamped and widely publishing the hash.
	- Anyone can easily see the entire history of Bitcoins mined, rather than having a single authority, person, or company responsible for keeping watch.
- The timestamp proves that the data must have existed at the time, obviously, in order to get into the hash. 
	- Each timestamp includes the previous timestamp in its hash, forming a chain, with each additional timestamp reinforcing the ones before it.

<img src="http://images.rapgenius.com/75bfa9367368c916cc45e81b568bbbb0.1000x332x1.png" alt="Drawing" style="width: 400px;"/>	

- To implement a distributed timestamp server on a peer-to-peer basis, we will need to use a proof-of-work system similar to Adam Back's Hashcash.
	- Proof-of-work system: 
		- An economic measure to deter denial of service attacks and other service abuses such as spam on a network by requiring some work from the service requester, usually meaning processing time by a computer.
		- Hard for the requester but easy to check for the service provider
		- This idea is also known as a CPU cost function, client puzzle, computational puzzle or CPU pricing function.
		- It is distinct from a CAPTCHA, which is intended for a human to solve quickly, rather than a computer.
			- In this case the computer has to solve it.
	- Hashcash is a proof-of-work algorithm, which has been used as a denial-of-service counter measure technique in a number of systems.
	- A hashcash stamp constitutes a proof-of-work which takes a parameterizable amount of work to compute for the sender. The recipient (and indeed anyone as it is publicly auditable) can verify received hashcash stamps efficiently.
	- The proof-of-work involves scanning for a value that when hashed, such as with SHA-256, the hash begins with a number of zero bits. 
	- The average work required is exponential in the number of zero bits required and can be verified by executing a single hash.
		- The proof-of-work system in Bitcoin is designed so that solving a new block (work) is difficult, but verifying the solution is very easy.
		- This means that mining new Bitcoins is difficult/slow, and proving that this work was done is relatively easy/fast (in terms of computing).
		- If fact, the difficulty or work required to solve new transactions is positively related to the overall network **hashrate**. 
	- This means the more computing power behind the Bitcoin network, the harder it becomes to mine new bitcoins (and vice-versa).

- For our timestamp network, we implement the proof-of-work by incrementing a nonce in the block until a value is found that gives the block's hash the required zero bits.
	- A nonce is an arbitrary number used only once in a cryptographic communication. 
	- It is often a random or pseudo-random number
- Once the CPU effort has been expended to make it satisfy the proof-of-work, the block cannot be changed without redoing the work.
- As later blocks are chained after it, the work to change the block would include redoing all the blocks after it.

<img src="http://images.rapgenius.com/34416a885a331715330987ac346111d8.1000x295x1.png" alt="Drawing" style="width: 400px;"/>	

- The proof-of-work also solves the problem of determining representation in majority decision making.
	- Proof-of-work is essentially one-CPU-one-vote. 
		- If the majority were based on one-IP-address-one-vote, it could be subverted by anyone able to allocate many IPs. So proof-of-work is beneficial here.
	- The majority decision is represented by the longest chain, which has the greatest proof-of-work effort invested in it.
	- If a majority of CPU power is controlled by honest nodes, the honest chain will grow the fastest and outpace any competing chains.
		- To modify a past block, an attacker would have to redo the proof-of-work of the block and all blocks after it and then catch up with and surpass the work of the honest nodes. 
		- To compensate for increasing hardware speed and varying interest in running nodes over time, the proof-of-work difficulty is determined by a moving average targeting an average number of blocks per hour. 
			- If they're generated too fast, the difficulty increases.

**Blockchain**

- The Blockchain enables a network of distributed nodes to achieve consensus — to agree on the common state of the network — by each proving that they had participated in that network through the proof-of-work system.
	- Take a Sudoku … it’s actually hard to figure out which numbers fit into the squares, [but] it’s very easy to check if the Sudoku’s right … easy to verify … hard to solve …. The proof-of-work which is used in Bitcoin is exactly the same, easy to verify very hard to solve.
	- https://www.youtube.com/watch?v=JP9-lAYngi4#t=87

**Network**

- A new block is generated every 10 minutes, and it is generally prudent for parties to wait for at least 6 new blocks in order for the transaction to be considered complete. 10 minutes is the time it takes for nodes (miners) in the network to “find a proof-of-work” (solve the mathematical puzzle created by SHA-256). 

- When the puzzle is solved by a miner and as long as the majority of the network agrees that the transactions in the block are valid, the block is added to the network. Off-blockchain transactions within a Bitcoin network are instantaneous.

- Steps to run the network:

	1. New transactions are broadcast to all nodes.
	2. Each node collects new transactions into a block.
	3. Each node works on finding a difficult proof-of-work for its block.
	4. When a node finds a proof-of-work, it broadcasts the block to all nodes.
	5. Nodes accept the block only if all transactions in it are valid and not already spent.
	6. Nodes express their acceptance of the block by working on creating the next block in the chain, using the hash of the accepted block as the previous hash.

- Nodes always consider the longest chain to be the correct one and will keep working on extending it.
- If two nodes broadcast different versions of the next block simultaneously, some nodes may receive one or the other first. 
	- In that case, they work on the first one they received, but save the other branch in case it becomes longer. The tie will be broken when the next proof of-work is found and one branch becomes longer; the nodes that were working on the other branch will then switch to the longer one.

- New transaction broadcasts do not necessarily need to reach all nodes. As long as they reach many nodes, they will get into a block before long. 
	- Block broadcasts are also tolerant of dropped messages. If a node does not receive a block, it will request it when it receives the next block and realizes it missed one.

**Blockchain fork**

- A blockchain fork occurred on March 12, 2013 due to a bug in Bitcoin-Qt version 0.7 that caused different nodes in the network to believe that different versions of the blockchain were correct. Just as Satoshi described in the paper, the longer chain ended up winning.
- [Read more](http://www.reddit.com/r/Bitcoin/comments/1a51xx/now_that_its_over_the_blockchain_fork_explained/)

**Incentive**



**Sources**

- https://bitcoin.org/bitcoin.pdf
- http://news.rapgenius.com/Satoshi-nakamoto-bitcoin-a-peer-to-peer-electronic-cash-system-annotated
- http://en.wikipedia.org/wiki/Proof-of-work_system
- http://en.wikipedia.org/wiki/Cryptographic_nonce

**Further Reading**

- http://www.reddit.com/r/Bitcoin/comments/18kt6y/psa_to_new_users_due_to_reddit_gold_announcement/
- http://bitcoinbook.info
- https://bitcoin.org/en/how-it-works
- http://www.righto.com/2014/02/bitcoins-hard-way-using-raw-bitcoin.html
- http://www.michaelnielsen.org/ddi/how-the-bitcoin-protocol-actually-works/
- http://www.youtube.com/channel/UCgo7FCCPuylVk4luP3JAgVw/videos
- https://github.com/bitcoin/bitcoin
- http://www.reddit.com/r/Bitcoin/comments/1a51xx/now_that_its_over_the_blockchain_fork_explained/