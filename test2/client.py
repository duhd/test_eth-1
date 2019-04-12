import requests
import datetime

#url = "http://localhost:8080/api/v1/wallet/transfer/ffbcd481c1330e180879b4d2b9b50642eea43c02/a17a7a153c8d873a1df803c74e0664c13726f5e8/2/Test"
url = "http://localhost:8080/api/v1/wallet/new_account"


def main():
    for i in range(1000):
        start = datetime.datetime.now()
        response = requests.get(url)
        diff = datetime.datetime.now() - start
        print(diff.total_seconds(), response.content)

if __name__== "__main__":
  main()
