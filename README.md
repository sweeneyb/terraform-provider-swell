# Terraform Provider Swell

Run the following command to build the provider

```shell
go build -o terraform-provider-swell.exe
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```

## Resources
### Category
Currently, only 3 fields are supported, and 'active' seems buggy

resource "swell_category" "coffee" {
  name = "coffee items"
  description = "items for coffee"
  active = false
}