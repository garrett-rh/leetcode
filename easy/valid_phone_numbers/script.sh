#!/bin/bash
grep -Eo "^\([[:digit:]]{3}\)[[:space:]][[:digit:]]{3}-[[:digit:]]{4}$|^[[:digit:]]{3}-[[:digit:]]{3}-[[:digit:]]{4}$" file.txt
