abigen -sol deployed.sol -pkg contracts -out deployed.go
abigen -sol existing.sol -pkg contracts -out existing.go
abigen -sol existingwithoutABI.sol -pkg contracts -out existingwithoutabi.go


abigen -sol metacoin.sol -pkg metacoin -out metacoin.go
abigen -sol coinspawn.sol -pkg metacoin -out coinspawn.go
abigen -sol coincaller.sol -pkg metacoin -out coincaller.go

coinspawn: 0xb2eb0bb624bd7d5275c0ce3ef0eef99fb77a6b74

caller1: 0xcdbc69a51e4835f5b04d710e7db10ba2db276d48
caller2: 0x7fbcd14bfa9dc7b2187ebd0956b148f2b6736f8b


0x307a235df8ac386aeb7c696aeb3039a7e8a6ce8bde097b41d4c1bba0ba9ab3a3

token: 0xAE6313a252d905cdc0d8e9116fE31696CC832145

1. Re-deploy metacoin contract => Redeploy all contract (if not no event)
