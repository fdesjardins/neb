#!/bin/bash

# show the version number, for logging
neb version

# verify we can connect to aws and show available commands
neb aws

# create a small instance interactively
neb aws ec2 make

# or, lets see what we can get for $10/mo...
neb aws ec2 ls -b 10

# or trust net to make us a sweet ec2 instance!
neb aws ec2 make -b 10

# make sure it's debian with 4 gigs of RAM
neb aws ec2 make debian-9.0 -b 10 -m 4096

# maybe neb can find a better cloud provider for our app
neb compare debian-9.0 -m 4096

# delete our ec2s
neb aws ec2 rm * -f
