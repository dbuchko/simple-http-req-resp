# simple-http-req-resp
simple app that logs a timestamp on an HTTP GET request.
<p>
Cross compile for linux diego cells using:
  <p>
env GOOS=linux GOARCH=amd64 go build
    <p>
Push to PCF using binary build pack with:
      <p>
cf push simplereqresp -c'./simple-http-req-resp' -b binary_buildpack -k 64m -m 64m

<p>
  Can test by curling the app like:
  <p>
  curl -X GET -H "vcap_request_id: myid" --verbose https://simplereqresp.<Apps Route>
<p>
