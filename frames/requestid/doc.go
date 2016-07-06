package requestid

/*
A Frame for generating a unique request id for each request. A generated
request id will in the format "uq8hs30oafhj5vve8ji5pmp7mtopc08f".

If a request id already exists as the "x-request-id" HTTP request header,
then that value will be used assuming it is between 20 and 200 characters.
If it is not, a new request id will be generated.


To use it, just stack it into the desired module:

  b := stack.NewBuilder()

  b.PushFunc(requestid.New())

## Options

  * `http_header` - The name of the HTTP *request* header to check for
  existing request ids. This is also the HTTP *response* header that will be
  set with the request id. Default value is "x-request-id"

  b.PushFunc(requestid.NewWithHeader("custom-request-id"))
*/
