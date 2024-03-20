require 'net/http'
require 'openssl'
require 'uri'

url = URI.parse('https://expired.badssl.com/')

http = Net::HTTP.new(url.host, url.port)
http.use_ssl = true

#http.verify_mode = OpenSSL::SSL::VERIFY_NONE

begin
  response = http.get(url)
  puts "Response Code: #{response.code}"
  puts "Response Body: #{response.body[0..499]}"
rescue => e
  puts "Failed to make a request: #{e.message}"
end

