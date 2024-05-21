# Dell sample scripts
Sample script in Python to get SAML token from Okta or Entra. SAML token can be written to a named pipe (file). 
In this example, `/tmp/samlPipe` needs to be created first.
The Terraform provider will read the SAML token from that pipe (file) and use it to get Dell Identity token.
One of these two sample scripts is needed to be setup in order to authenitcate with Apex.
**After the steps below are set. The user must set the enviorment variables. like so for windows and linux respectively**
[windows_enviorment_variables](https://github.com/dell/terraform-provider-apex/blob/main/examples/provider/env.bat)
[linux_enviorment_variables](https://github.com/dell/terraform-provider-apex/blob/main/examples/provider/env.sh)

# Getting Started Okta
Install the Python prerequisites using the requirements.txt file `pip install -r requirements.txt`
Copy either Windows (saml_get_token_okta_idp_win) or Linux (saml_get_token_okta_idp_lin) file to saml_get_token_okta_idp.py as per the platform (line 86 of saml_get_okta_token). 

Update config-okta.json (Set the saml_endpoint, login_endpoint, username, password, saml_pipe according to the directions in the config-okta.json)
For more information on using the script check https://developer.dell.com/apis/83acf3bf-becf-41d2-aa1c-b3deffc5e549/versions/1.0.0/docs/Auth.md

# Setup your Python development environment for Okta

1. Clone the repo https://github.com/dell/terraform-provider-apex as normal
2. Script is tested with Python 3.
3. The script is ready to run after 
```console
pip install -r requirements.txt
```
4. Then, run 
```console
python saml_get_okta_token.py         # outputs token and debug messages
```

# Getting Started with Microsoft Entra

1. Clone the repo https://github.com/dell/terraform-provider-apex as normal
2. Install the Python prerequisites using the requirements.txt file in `/scripts/saml_scripts` Command: `pip install -r requirements.txt`
3. Update the config-entra.json (Set the saml_endpoint, username, password, saml_pipe, assertion_consumer_service_url, and entra_app_client_secret according to the directions in the config-entra.json)
4. Then, run
```
python saml_get_entra_token.py
```

Note: If 'python' is not found create a soft link for python to python3