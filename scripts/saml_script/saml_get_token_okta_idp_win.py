# Copyright © 2023 Dell Inc. or its subsidiaries.

# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”),to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so

# THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE

import re
from urllib.parse import urlparse
import requests
import json
from urllib.parse import urlparse
import base64
from dotenv import load_dotenv
from bs4 import BeautifulSoup
import sys
import getpass
import requests
from configparser import ConfigParser
import base64
import xml.etree.ElementTree as ET
from bs4 import BeautifulSoup
from os.path import expanduser
# from urlparse import urlparse, urlunparse
from requests_ntlm import HttpNtlmAuth
import logging 
import win32pipe
import os

logger=logging.getLogger() 
logger.setLevel(logging.DEBUG)

# options
OKTA_LOGIN_ENDPOINT = 'okta-login-endpoint'
OKTA_SAML_ENDPOINT = 'okta-saml-endpoint'

# receving the smal and okta url's from config json
if len (sys.argv) >=3:
 DEFAULT_LOGIN_ENDPOINT = sys.argv[1]
 DEFAULT_SAML_ENDPOINT = sys.argv[2]
 DI_URL = sys.argv[3]
 

USERNAME = 'IDP_USERNAME'
PASSWORD = 'IDP_PASSWORD'

class OktaSamlResponse:

    def __init__(self, options: dict = None):
        self.options = options or {}
        self.saml_response = None

        self.login_endpoint = self.options.get(
            OKTA_LOGIN_ENDPOINT) or DEFAULT_LOGIN_ENDPOINT
        self.saml_endpoint = self.options.get(
            OKTA_SAML_ENDPOINT) or DEFAULT_SAML_ENDPOINT

        load_dotenv()

    def value(self):
        if self.saml_response is None:
           self.__get_new()

        return self.saml_response

    def __get_new(self):
        
        session = requests.Session()
        try:
            response = requests.get(DEFAULT_LOGIN_ENDPOINT)
            response.raise_for_status() 
             # Check for HTTP errors
            # Process the response if successful
        except requests.exceptions.RequestException as e:
            print(f"invalid url: {e}")
            sys.exit(0)
        # get the sign-in page
        sign_in_page = session.get(self.login_endpoint, verify=None)
        # Username and password should be substituted with right values
        username = "username"
        password = "password"

        if username is None or password is None:
            raise Exception("Environment variables " + USERNAME + " and/or "
                            + PASSWORD + " are missing.")

        sign_in_payload = self.__parse_sign_in_page(sign_in_page, username,
                                                    password)

        # find the action of the form to submit the form
        
        action = self.__find_sign_in_page_action(sign_in_page)
        if not action:
            print("Exiting as no form action found... ")
            return
        parsed_url = urlparse(self.login_endpoint)
        idp_base_url = parsed_url.scheme + "://" + parsed_url.netloc
        form_submit_url = idp_base_url + action

        # submit the sign-in form
        sign_in_page_response = session.post(
                form_submit_url, data=sign_in_payload)
    
        authn_url = idp_base_url + '/api/v1/authn'
        authn_response = session.post(authn_url,
                                      json={'username': username,
                                            'password': password})

        #validate user credentials
        session_token =self.validate_session_token(authn_response)

        # get the session id (sid)
        sessions_url = idp_base_url + '/api/v1/sessions'
        sessions_response = session.post(sessions_url,
                                         json={"sessionToken": session_token})

        sessions_json = json.loads(sessions_response.text)
        sid = sessions_json['id']

        # get the SAML response
        saml_page_response = session.get(self.saml_endpoint,
                                         cookies={'sid': sid,
                                                  "DT": session.cookies.get(
                                                      "DT"),
                                                  "JSESSIONID": session.cookies.get(
                                                      "JSESSIONID")})

        saml_response =self.__find_saml_response(saml_page_response)
        print("==================")
        print(saml_response)
        #print()
        #print("Assertion is:", saml_response)
        #print()
        print("==================")
        #print()
     


       # print("==== Retreiving DI Bearer token ===")
       # print()
        pipeName = "/tmp/samlPipe"
        try:
            pipe1 = win32pipe.CreateNamedPipe(pipeName, win32pipe.PIPE_ACCESS_DUPLEX,
                              win32pipe.PIPE_TYPE_MESSAGE | win32pipe.PIPE_READMODE_MESSAGE,
                              1, 65536, 65536, 0, None)
            print(f"Named pipe '{pipeName}' created successfully!")
        except Exception as e:
            print(f"Error creating named pipe: {e}")
        # Write data to the pipe
   
        # Open the FIFO pipe for writing
        with open(pipeName, "w") as pipe:
        # Write data to the pipe
    
            pipe.write(saml_response)
            pipe.close()
    @staticmethod
    def validate_session_token(authn_response):
      # get the session token
        status_ok =200
        authn_json = json.loads(authn_response.text)
        if authn_response.status_code == status_ok:  
            session_token = authn_json['sessionToken']
        else:
            print("ERROR: Invalid Credentials/ No assertion found in the response... ")
            sys.exit()
        return session_token

    @staticmethod
    def __parse_sign_in_page(sign_in_page_response, username, password):
        # Parse html of sign-in page and return payload to be submitted
        form_soup = BeautifulSoup(sign_in_page_response.text,
                                  features='html.parser')
        payload = {}
        for input_tag in form_soup.find_all(re.compile('(INPUT|input)')):
            name = input_tag.get('name', '')
            value = input_tag.get('value', '')
            if "user" in name.lower():
                # Make an educated guess that this is correct field for username
                payload[name] = username
            elif "email" in name.lower():
                # Some IdPs also label the username field as 'email'
                payload[name] = username
            elif "pass" in name.lower():
                # Make an educated guess that this is correct field for password
                payload[name] = password
            else:
                # Populate the parameter with existing value
                # (picks up hidden fields in the login form)
                payload[name] = value
        return payload

    @staticmethod
    def __find_sign_in_page_action(form_response):
        form_soup = BeautifulSoup(form_response.text, features='html.parser')
        for input_tag in form_soup.find_all(re.compile('(FORM|form)')):
            action = input_tag.get('action')
            if not action:
                print("ERROR: no action found on sign in page")
            return action if action else None

    @staticmethod
    def __find_saml_response(response):
        soup = BeautifulSoup(response.text, features='html.parser')
        saml_response = None

        # Look for the SAMLResponse attribute of the input tag
        # (determined by analyzing the debug print lines above)
        for input_tag in soup.find_all('input'):
            if input_tag.get('name') == 'SAMLResponse':
                saml_response = input_tag.get('value')
        return saml_response

oktaSamlResponse = OktaSamlResponse()
oktaSamlResponse.value()
