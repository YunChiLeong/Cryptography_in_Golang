[![Open in Visual Studio Code](https://classroom.github.com/assets/open-in-vscode-f059dc9a6f8d3a56e377f745f24479a46679e63a5d9fe6f495e02850cd0d8118.svg)](https://classroom.github.com/online_ide?assignment_repo_id=6259274&assignment_repo_type=AssignmentRepo)
# HW 3 - Certificate Signing Requests

The goal for this assignment is for you to write a Go program that can generate a valid Certificate Signing Request (CSR) and then use the CSR to retrieve a certificate from a Certificate authority I created.

## CSR Requirements

Your certificate signing request must conform to the following requirements.

- The `Signature Algorithm` field _must_ be set to object identifier for `ed25519`. There is a constant in the x509 package that will help with this. This means the public key you generate for the certificate must also be an `ed25519` key. Make sure to also save the private key, even though it isn't part of the CSR.
- The `Subject` of your CSR must have the `CN`, `O`, and `OU` fields set. I don't care what you set them to necessarily, just make sure they are appropriate.
- Your CSR _must_ include your email address in the `Subject Alternative Name` extension. You don't need to include any DNS names or IP addresses as SANs.

## CA Interface

Here are some basic instructions on how to interact with the hosted CA that you'll need to submit your CSR to.

The CA will be hosted on `https://crypt.invariant.dev` and exposes a simple http api. It listens on a single path, `/sign`. In order to get your CSR accepted and a certificate returned, the format of your request must meet the following requirements.

- Your request must be to the `/sign` path. I.e. `https://crypt.invariant.dev/sign`.
- Your request _must_ be a `POST` request.
- Your request _must_ contain a header called `Csr_Auth_Code`. The value of this header must be a code that I will give to each of you individually. This is how the CA will determine whether or not it should grant your request for a certificate.
- Your CSR _must_ be included in the body of the http request as a binary file. 

The result of the http request will either be an error message in plain text describing what error occured, or it will be the signed certificate generated from your CSR. The `Content-Type` and `Content-Disposition` headers will be set appropriately. 

In addition to the requirements above, you _must_ perform this request in Go code. Feel free to test with `curl` or `Postman` or whatever else, but you need to submit code that actually performs the request. This can be a separate Go program or the same one that generates the CSR, just document whichever you choose to do.

## Connect as mTLS Client

Once you have your signed certificate, you will test that it works by writing some more Go code that makes an HTTP GET request to a webserver at `https://mtls.invariant.dev` and is configured to present the certificate you just obtained. You can do this by creating your own `http.Client` instead of using the default one; you will need to read the documentation on how to set up the client to present a certificate. The server will simply give you back a response that says `Success!` in text if the client certificate validation was successful. Once you've gotten the success message, you're done with the assignment.