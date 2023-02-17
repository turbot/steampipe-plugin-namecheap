connection "namecheap" {
  plugin = "namecheap"

  # Username / API User is required for requests. Required.
  # username = "janedoe"

  # A specific API User can also be defined. Optional, by default, this will be
  # set to the `username`.
  api_user = "janedoe"

  # API key for requests. Required.
  # See instructions at https://www.namecheap.com/support/api/intro/.
  # api_key = "33d0d62a6a163083ba7b3bab31bd6612"

  # Namecheap requires your IP address to be granted access on the API, and to
  # be explicitly defined in the request. Required.
  # client_ip = "1.2.3.4"
}
