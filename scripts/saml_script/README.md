# Dell sample script 
Sample script in Python to get SAML token from Okta. SAML token can be written to a named pipe (file). 
In this example, /tmp/samlPipe needs to be created first.
Provider will read the SAML token from that pipe and use it to get Dell Identity token.


# Getting Started
Install prerequisites to run Python script in saml_script directory. 
Copy either Windows (saml_get_token_okta_idp_win) or Linux (saml_get_token_okta_idp_lin) file to saml_get_token_okta_idp.py as per the platform. 

Update userid/password in saml_get_token_okta_idp.py (line 79, 80)
Update config.json in saml_script
For more information on using the script check https://developer.dell.com/apis/83acf3bf-becf-41d2-aa1c-b3deffc5e549/versions/1.0.0/docs/Auth.md

You will need to setup environmental variables provided in examples/provider folder for the provider.

# Setup your Python development environment

1. Clone the repo as normal
2. Script is tested with Python 3.
3. The script is ready to run after 
```console
pip install -r requirements.txt
```
4. Then, run 
```console
python saml_get_token.py         # outputs token and debug messages
```

