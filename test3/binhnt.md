abigen -sol deployed.sol -pkg contracts -out deployed.go
abigen -sol existing.sol -pkg contracts -out existing.go
abigen -sol existingwithoutABI.sol -pkg contracts -out existingwithoutabi.go


abigen -sol metacoin.sol -pkg metacoin -out metacoin.go
abigen -sol coinspawn.sol -pkg metacoin -out coinspawn.go
abigen -sol coincaller.sol -pkg metacoin -out coincaller.go

coinspawn: 0xb2eb0bb624bd7d5275c0ce3ef0eef99fb77a6b74

caller1: 0x495303d2faceafa41b20aaab3f6aeccf797d7528
caller2: 0xf1a79b3123eb75a7cf1230601383f8da484a55a6


token: 0xAE6313a252d905cdc0d8e9116fE31696CC832145

1. Re-deploy metacoin contract => Redeploy all contract (if not no event)
