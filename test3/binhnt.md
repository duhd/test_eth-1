abigen -sol deployed.sol -pkg contracts -out deployed.go
abigen -sol existing.sol -pkg contracts -out existing.go
abigen -sol existingwithoutABI.sol -pkg contracts -out existingwithoutabi.go


abigen -sol metacoin.sol -pkg metacoin -out metacoin.go
abigen -sol coinspawn.sol -pkg metacoin -out coinspawn.go
abigen -sol coincaller.sol -pkg metacoin -out coincaller.go

coinspawn: 0x47562495779f98c048460514becf5e1a2d217c9f

caller1: 0xa341d306c6c90d19f7be11d4d347bd4206854b5c
caller2: 0x27af2a27850e185ef3ccab0c62fdea25ace9a8f9


token: 0xAE6313a252d905cdc0d8e9116fE31696CC832145

1. Re-deploy metacoin contract => Redeploy all contract (if not no event)
