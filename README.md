# Terraform Provider Bria

This project is a Terraform provider for Bria, allowing you to manage Bria resources using Terraform.

## Development Workflow

1. Enter the development environment:
```
nix develop
```

2. Build and install the provider:
```
make install
```

3. Navigate to the `examples` directory:
```
cd examples
```

4. Initialize Terraform:
```
terraform init
```

5. Apply the Terraform configuration:
```
terraform apply
```

Ensure that Bria is running locally before executing the Terraform commands.

## Example

Refer to the [examples](./examples) directory for sample Terraform configurations using this provider.

## Dependencies

Dependencies are managed using [flake.nix](./flake.nix).
