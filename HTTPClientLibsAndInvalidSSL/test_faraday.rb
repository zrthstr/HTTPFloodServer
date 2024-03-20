require 'faraday'
require 'openssl'

#conn = Faraday.new(url: 'https://expired.badssl.com/', ssl: { verify: false })
conn = Faraday.new(url: 'https://expired.badssl.com/')

response = conn.get('/')
puts response.body
