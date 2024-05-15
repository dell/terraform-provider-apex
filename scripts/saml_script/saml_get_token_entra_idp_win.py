# Copyright © 2023 Dell Inc. or its subsidiaries.

# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the “Software”),to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so

# THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE

import re
import time
from urllib.parse import quote
import zlib
import requests
import base64
from dotenv import load_dotenv
import sys
import requests
import base64
import xml.etree.ElementTree as ET
# from urlparse import urlparse, urlunparse
import logging 
import win32pipe
from selenium import webdriver
from selenium.webdriver.chrome.options import Options
from selenium.webdriver.common.by import By
from datetime import datetime
from urllib.parse import unquote

logger=logging.getLogger() 
logger.setLevel(logging.DEBUG)

# receving the smal and okta url's from config json
if len (sys.argv) >=3:
 DEFAULT_LOGIN_ENDPOINT = sys.argv[1]
 DEFAULT_SAML_ENDPOINT = sys.argv[2]
 USERNAME = sys.argv[3]
 PASSWORD = sys.argv[4]
 SAML_PIPE = sys.argv[5]
 ASSERTION_CONSUMER_URL = sys.argv[6]
 ENTRA_APP_CLIENT_SECRET = sys.argv[7]
 DI_URL = sys.argv[3]

class OktaSamlResponse:

    def __init__(self, options: dict = None):
        self.options = options or {}
        self.saml_response = None

        self.login_endpoint = DEFAULT_LOGIN_ENDPOINT
        self.saml_endpoint = DEFAULT_SAML_ENDPOINT
        self.username = USERNAME
        self.password = PASSWORD
        self.saml_pipe = SAML_PIPE
        self.assertion_consumer_service_url = ASSERTION_CONSUMER_URL
        self.entra_app_client_secret = ENTRA_APP_CLIENT_SECRET

        load_dotenv()

    def value(self):
        if self.saml_response is None:
           self.__get_new()

        return self.saml_response

    def __get_new(self):
     
        try:
            response = requests.get(DEFAULT_LOGIN_ENDPOINT)
            response.raise_for_status() 
             # Check for HTTP errors
            # Process the response if successful
        except requests.exceptions.RequestException as e:
            print(f"invalid url: {e}")
            sys.exit(0)

        if self.username is None or self.password is None:
            raise Exception("Environment variables " + USERNAME + " and/or "
                            + PASSWORD + " are missing.")

        
        options = Options()
        options.add_argument('--headless')
        options.add_argument('--disable-gpu')
        options.set_capability('goog:loggingPrefs', {'performance': 'ALL'})
        driver = webdriver.Chrome(options=options)
        driver.get(self.login_endpoint)
        time.sleep(2)
        inputEl = driver.find_element(By.XPATH,'//input[@type="email"]')
        inputEl.send_keys(self.username)

        inputNextButton = driver.find_element(By.ID,'idSIButton9')
        inputNextButton.click()

        time.sleep(1)

        inputElPass = driver.find_element(By.XPATH,'//input[@type="password"]')
        inputElPass.send_keys(self.password)

        inputNextButtonSignin = driver.find_element(By.ID,'idSIButton9')
        inputNextButtonSignin.click()

        time.sleep(1)

        inputNextButtonStaySignedin = driver.find_element(By.ID,'idSIButton9')
        inputNextButtonStaySignedin.click()

        time.sleep(2)
        redirect = self._generate_saml_request_redirect_binding()
        driver.get(redirect + "&RelayState=https://www.dell.com/account")

        time.sleep(3)
        logs = driver.get_log("performance")
        saml_response = self._extract_value(str(logs))
        # print("Extracted value Decoded:", decode_base64(
        #     saml_response
        # ))

        driver.quit()

        pipeName = self.saml_pipe
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

    def _extract_value(self, input_string):
        pattern = r'SAMLResponse=(.*?)"'
        match = re.search(pattern, input_string)
        if match:
            value = match.group(1)
            value = unquote(value)
            parts = value.split("&")
            return parts[0]
        else:
            return None

    def _generate_saml_request(self):
        authn_request = ET.Element(
            "samlp:AuthnRequest",
            {
                "xmlns:samlp": "urn:oasis:names:tc:SAML:2.0:protocol",
                "xmlns": "urn:oasis:names:tc:SAML:2.0:metadata",
                "ID": self.entra_app_client_secret,
                "Version": "2.0",
                "IssueInstant": datetime.utcnow().strftime("%Y-%m-%dT%H:%M:%SZ"),
                "AssertionConsumerServiceURL": self.assertion_consumer_service_url,
                "ForceAuthn": "false",
                "IsPassive": "false",
            },
        )
        authn_request = self._add_issuer(authn_request)

        request_xml = ET.tostring(authn_request, encoding="unicode")
        return request_xml

    def _add_issuer(self, authn_request):
        issuer_element = ET.SubElement(
            authn_request, "Issuer", {"xmlns": "urn:oasis:names:tc:SAML:2.0:assertion"}
        )
        issuer_element.text = "https://idhub.dell.com"
        print(authn_request)
        return authn_request

    def _generate_saml_request_redirect_binding(
        self
    ):
        saml_request = self._generate_saml_request()
        packed_xml = self._pack_saml_request(saml_request)
        saml_request_param = f"SAMLRequest={packed_xml}"

        return f"{self.saml_endpoint}?{saml_request_param}"

    def _pack_saml_request(self, request_xml):
        request_deflated = zlib.compress(request_xml.encode("utf-8"))[2:-4]
        saml_request_encoded = base64.b64encode(request_deflated).decode("utf-8")
        return quote(saml_request_encoded)


def decode_base64(encoded_str):
    decoded_bytes = base64.b64decode(encoded_str)
    decoded_str = decoded_bytes.decode("utf-8")
    return decoded_str

oktaSamlResponse = OktaSamlResponse()
oktaSamlResponse.value()
