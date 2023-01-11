# digests terraform provider

## How to use

- Prepare provider binary

    ```shell
    # build plugin binary
    go build -o terraform-provider-digests .
    # if you are using windows, don't forget extension ".exe"
    
    # create plugin directory
    mkdir "${TF_PLUGIN_DIR}/local/10hin/digests/0.0.1/${OS}_${ARCH}"
    # -> ${TF_PLGIN_DIR} : local plugin mirror
    # -> ${OS} : "linux", "windows", etc.
    # -> ${ARCH} : "amd64", etc.
    
    # copy provider binary
    cp terraform-provider-digests "${TF_PLUGIN_DIR}/local/10hin/digests/0.0.1/${OS}_${ARCH}/"
    ```

- Configure terraform CLI
    - create `.terraformrc`/`terraform.rc` file according to [Document](https://developer.hashicorp.com/terraform/cli/config/config-file)
    - configure provider installation method (cf. [doc](https://developer.hashicorp.com/terraform/cli/config/config-file#provider-installation))

- Move to `example/` directory
- Execute `terraform init`/`terraform validate`/`terraform plan`/`terraform apply`
