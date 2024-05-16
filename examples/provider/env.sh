# These values are used to help get the SAML token from the Users IDP IE:(Okta or Entra)
# For more Information check out https://github.com/dell/terraform-provider-apex/blob/main/scripts/saml_script/README.md

# Path to Script directory
export APEX_SAML_TOKEN_SCRIPT_DIR="./scripts/saml_script"
# Python SAML script to run when we need to get a new SAML token.
export APEX_SAML_TOKEN_SCRIPT_NAME="saml_get_okta_token.py"
# The file you output your saml token too
export APEX_SAML_TOKEN_OUTPUT_FILE_NAME="\tmp\samlPipe"
# The DI Token URL (User should not have to change this)
export APEX_DI_TOKEN_URL="https://www.dell.com/di/v3/fp/oauth/token"
# The Apex Users Access Key (Can be retrived through the Dell Premier Apex UI)
export APEX_TOKEN_ACCESS_KEY="aaa-bbb-ccc"
# The Apex Users Token Secret (Can be retrived through the Dell Premier Apex UI)
export APEX_TOKEN_SECRET="random"