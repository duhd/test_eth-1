package utils

type ClientPool struct {
  Clients []*EthClient
  Current int
}

var clientPool *ClientPool


func NewClientPool()  *ClientPool{
    var clients []*EthClient
    max_connection := cfg.Webserver.MaxRpcConnection
    for i:=0 ; i<max_connection; i++ {
         for _,host := range cfg.Networks {
              ethclient, err := NewEthClient(host.Http)
              if err != nil {
                continue
              }
              clients = append(clients,ethclient)
            }
     }
     clientPool =  &ClientPool{
        Clients: clients,
        Current:0,
     }
     return clientPool
}

func (cp *ClientPool) GetClient() (*EthClient) {
    client := cp.Clients[cp.Current]
    cp.Current = cp.Current + 1
    cp.Current = cp.Current % len(cp.Clients)
    return client
}
