# Use the official Ubuntu 20.04 as base
# Add group and user and switch to it
FROM ubuntu:20.04
RUN apt-get update && apt-get upgrade -y
RUN groupadd -g 1000 basicuser && useradd -r -u 1000 -g basicuser basicuser
USER basicuser