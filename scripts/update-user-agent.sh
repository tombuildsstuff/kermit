#!/bin/bash

# Since this SDK is no longer present within `Azure/azure-sdk-for-go`, using the user agent/version number from that
# is misleading - it also means we unintentionally import the Azure SDK for Go.
# Instead let's update the user agent to use a local version with the version number at release time

if [[ "$OSTYPE" == "darwin"* ]]; then
  find ../../sdk -type f -print0 | xargs -0 sed -i '' -e "s/Azure\/azure-sdk-for-go\/version/tombuildsstuff\/kermit\/version/" -e "s/Azure-SDK-For-Go\//tombuildsstuff\/kermit\//"
else
  find ../../sdk -type f -print0 | xargs -0 sed -i -e "s/Azure\/azure-sdk-for-go\/version/tombuildsstuff\/kermit\/version/" -e "s/Azure-SDK-For-Go\//tombuildsstuff\/kermit\//"
fi

