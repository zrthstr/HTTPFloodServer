import requests
from requests.packages.urllib3.exceptions import InsecureRequestWarning

#requests.packages.urllib3.disable_warnings(InsecureRequestWarning)

url = 'https://expired.badssl.com/'

#response = requests.get(url, verify=False)
response = requests.get(url )

print(f"Status Code: {response.status_code}")
print(f"Response Body: {response.text[:500]}")
