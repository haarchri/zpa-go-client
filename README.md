# zpa-go-client

Go client for Zscaler ZPA

# Code Generation

[go-swagger](https://github.com/go-swagger/go-swagger) is used in this repo to generate the client code needed to connect to the Styra API.

To generate the code [install go-swagger](https://goswagger.io/install.html) and run:

```bash
swagger generate client -f doc/zpa-api-docs.json --skip-validation -c pkg/client -m pkg/models --default-scheme=https
```

