import urllib3
#urllib3.disable_warnings(urllib3.exceptions.InsecureRequestWarning)

http = urllib3.PoolManager(cert_reqs='CERT_NONE')
#url = 'https://expired.badssl.com/'
url = 'https://wrong.host.badssl.com/'
response = http.request('GET', url)

print(f"Status Code: {response.status}")
print(f"Response Body: {response.data[:500]}")
