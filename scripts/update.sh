#!/bin/bash

# Author: cyberpunkcoder (github.com/cyberpunkcoder) (cyberpunkcoder@gmail.com)
# Date: 12 1 2020
# Description: Updates the gobot executable by downloading the latest release from github.

URL=$(curl -s https://api.github.com/repos/cyberpunkcoder/gobot/releases/latest | grep browser_download_url | cut -d '"' -f 4)
FILE=$(basename $URL)

rm $FILE
wget $URL
chmod +x $FILE
