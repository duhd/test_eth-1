Running

### Preparing

+ go run add_peers.go
+ go run deploy_wallet.go
### Running webserver & listening server

+ go run web_server.go
+ go run block_subscribe.go
### Running client to test







###   Support functions

abigen -sol VNDWallet.sol -pkg contracts -out vndwallet.go


contract: 0x495303d2faceafa41b20aaab3f6aeccf797d7528

#0.Initial network
go run add_peers.go "http://103.126.156.17:8501,http://103.126.156.18:8502"

# 1. Deploy contract
go run deploy_wallet.go

# 2. View Eth balance
go run view_eth_balance.go  "http://103.126.156.17:8501"  "0xffbcd481c1330e180879b4d2b9b50642eea43c02"

# 3. View Token balance
go run view_token_account.go  "http://103.126.156.17:8501"   "0x495303d2faceafa41b20aaab3f6aeccf797d7528" "0xffbcd481c1330e180879b4d2b9b50642eea43c02"

# 4. Transfer token
go run transfer_token.go ./keystore/UTC--2019-03-18T03-35-27.638151000Z--ffbcd481c1330e180879b4d2b9b50642eea43c02  "http://103.126.156.17:8501" "123456" "0x495303d2faceafa41b20aaab3f6aeccf797d7528" "0xa17a7a153c8d873a1df803c74e0664c13726f5e8" 1 "Chuyen tien"

# 5. View balance of received address
go run view_token_account.go  "http://103.126.156.17:8501"   "0x495303d2faceafa41b20aaab3f6aeccf797d7528" "0xa17a7a153c8d873a1df803c74e0664c13726f5e8"

# 6. Get event of transaction
go run event_read.go  "http://103.126.156.17:8501"  "0x495303d2faceafa41b20aaab3f6aeccf797d7528"  "0xffbcd481c1330e180879b4d2b9b50642eea43c02" "0xa17a7a153c8d873a1df803c74e0664c13726f5e8"

#7. Get transaction state
go run get_transaction_state.go  "http://103.126.156.17:8501"  "0xb0f5632557e0efc5031ab3c1dad137c855ff355b50f7422aa384c4cc559445f4"

#8. Subcribe to transfer event
go run event_subcribe_transfer.go  "ws://103.126.156.17:8541"  "0x495303d2faceafa41b20aaab3f6aeccf797d7528" "0xffbcd481c1330e180879b4d2b9b50642eea43c02,0xa17a7a153c8d873a1df803c74e0664c13726f5e8" "0xffbcd481c1330e180879b4d2b9b50642eea43c02,0xa17a7a153c8d873a1df803c74e0664c13726f5e8"


go run get_transaction_state.go  "http://103.126.156.17:8501"  "0xf22f2cf8ee86297e1fcd1d7e2947f44d482d937ea148a07435d2a7b4d1b8494c"

go run block_subscribe.go  "ws://103.126.156.17:8541"  
