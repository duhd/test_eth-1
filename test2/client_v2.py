import requests
import datetime

#url = "http://localhost:8080/api/v1/wallet/transfer/ffbcd481c1330e180879b4d2b9b50642eea43c02/a17a7a153c8d873a1df803c74e0664c13726f5e8/2/Test"
def createWallet(name):
    url = "http://localhost:8080/api/v2/wallet/create/{}".format(name)
    start = datetime.datetime.now()
    response = requests.get(url)
    diff = datetime.datetime.now() - start
    print(diff.total_seconds(), response.content)

def Balance(name):
    url = "http://localhost:8080/api/v2/wallet/balance/{}".format(name)
    start = datetime.datetime.now()
    response = requests.get(url)
    diff = datetime.datetime.now() - start
    print(diff.total_seconds(), response.content)

def main():
    createWallet("vi01")
    Balance("vi01")

if __name__== "__main__":
  main()
