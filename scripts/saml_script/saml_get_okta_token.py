# Copyright © 2023 Dell Inc. or its subsidiaries.

# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”),to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so

# THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE

import sys
import subprocess
import xml.etree.ElementTree as ET
import json

# CONSTANTS
# SSL certificate verification: Whether or not strict certificate
# verification is done, False should only be used for dev/test
# when sssl verification is true, either the certificate has to be from a public CA and present in store
# in case of self signed certs, path to the cert (pem format) has to be given to the sslverification variable
# sslverification = 'apexiam-cert.pem'
# read config json to load the idp and saml ur's
config_file_path = "config-okta.json"
try:
    with open(config_file_path, encoding="utf8") as config_file:
        config_data = json.load(config_file)
    oktaloginendpoint = config_data["okta"]["login_endpoint"]
    oktasamlendpoint = config_data["okta"]["saml_endpoint"]
    oktauser = config_data["okta"]["username"]
    oktapassword = config_data["okta"]["password"]
    oktasamlpipe = config_data["okta"]["saml_pipe"]
    diurl = config_data["Exchange_URL"]
except FileNotFoundError:
    print("Config file not found.")
    exit(0)
except KeyError:
    print("key not found in the config file. Please update the url in config file.")
    exit(0)
    
def save_config(file_path, config_data):
            with open(file_path, 'w') as config_file:
                json.dump(config_data, config_file, indent=4)
#provide chocie to user to choose
def let_user_pick(options):
    print("Please choose:")

    for idx, element in enumerate(options):
        print("{}) {}".format(idx + 1, element))

    i = input("Enter number: ")
    print()
    try:
        if 0 < int(i) <= len(options):
            return int(i) - 1
    except:
        pass
    return None
options = ["okta",'exit']
res = 0 #let_user_pick(options)
if options[res] in config_data:
    idp_config = config_data[options[res]]
    if "okta"  in options[res].lower():
       print(f"The below are example parameters vaules for okta saml_endpoint and login_endpoint:")
       print(f"Current configuration for the url's:")
       print()
       print(f"IDP URL: {idp_config['login_endpoint']}")
       print()
       print(f"SAML URL: {idp_config['saml_endpoint']}")
       print()
    edit_choice = "no" #input("Do you want to edit this IDP's configuration? (yes/no): ").strip()
    if edit_choice.lower() == 'yes':
        if "okta"  in options[res].lower():
            print()
            idp_config['login_endpoint'] = input("Enter New IDP URL: ").strip()
            print()
            edit_choice = input("Do you want to edit this SAML configuration? (yes/no): ").strip()
            print()
            if edit_choice.lower() == 'yes':
                idp_config['saml_endpoint'] = input("Enter New SAML URL: ").strip()
            else:
                print("none")
        save_config(config_file_path, config_data)
        print()
        print("Configuration updated successfully, re-run the script after url update.")
        sys.exit(0)
    else:
            def handler():
                """main handler"""
                if "okta" in options[res].lower():
                    # saml_get_token_okta_idp.py should be updated to either saml_get_token_okta_idp_lin or saml_get_token_okta_idp_win
                    subprocess.run(["python", "saml_get_token_okta_idp.py",oktaloginendpoint,oktasamlendpoint,diurl, oktauser, oktapassword, oktasamlpipe])
                    quit()   

    handler()
