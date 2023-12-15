# How to use TLS/SSL in golang

`openssl` is a actually cool. The following are some of the things you can do using `openssl`
- create self-singed certificates.
- encrypt and decrypt files.
- checksum and hashing.
- SSL/TLS connection testing.

The following is how the certificates should be organised.

```text
-----BEGIN CERTIFICATE-----
(Your Root CA Certificate)
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
(Your Intermediate CA Certificate)
-----END CERTIFICATE-----

-----BEGIN CERTIFICATE-----
(Your Server Certificate)
-----END CERTIFICATE-----
```
