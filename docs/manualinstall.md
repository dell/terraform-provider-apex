## Setting Up Provider Environment

1. Clone the repository
2. Enter the repository directory
3. Set environment variables for both host url and token
    ```shell
    export TF_VAR_JWT_TOKEN=ey12345...
    export TF_VAR_HOST=url              
    ```
4. For building the provider and automatically moving it to proper directory, run 
    ```shell
    make build-provider
    ```
    - For example this will move it to `~/.terraform.d/plugins/<your_plugin_path>/<version>/<OS_ARCH>`
5. Run 
    ```shell
    cd examples/                  // cd into examples folder containing .tf file
    ```

6. Terraform as desired!

7. Activate terraform logs!
   
   Terraform has detailed logs that can be enabled as shown below.
   
   To turn on logs, run
    ```shell
    export TF_LOG=<level>
    ```
    Most logs are on level "INFO", but here is [a list of terraform log types](https://www.terraform.io/internals/debugging) that vary in terms of verbosity and output types. For now, "INFO" level should be fine.

# Using Provider (no Development)
1. If you're not developing, simply download the desired provider binary and copy it to a directory within your TF plugin directory like so:
    ```shell
    mkdir -p ~/.terraform.d/plugins/<your_plugin_path>/<version>/<OS_ARCH>
    cp <plugin_binary> ~/.terraform.d/plugins/<your_plugin_path>/<version>/<OS_ARCH>
    ```
2. Now you can automatically import the functionality of this provider by adding the provider to the "require" clause at the start of your `.tf` file