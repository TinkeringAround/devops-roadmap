#!/usr/bin/env bash

# Reading in user name and password
echo "What is your intended username?"
read -r USERNAME
echo "What is your password?"
read -r PASSWORD

# Creating the user and setting the password
echo "Creating user $USERNAME"
useradd -m "$USERNAME"
chpasswd <<<"$USERNAME":"$PASSWORD"
echo "$USERNAME user account being created."
