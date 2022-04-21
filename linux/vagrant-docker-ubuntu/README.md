# Vagrant + Docker + Ubuntu

Starting an Ubuntu Docker Container which then runs in the background.

````shell
vagrant up
````

Start Bash in the Ubuntu Container for the `default` vm.
`default` cannot easily be changed as naming.

````shell
vagrant docker-exec -it default -- /bin/bash
````

Stopping and destroying a vm.

````shell
vagrant halt
vagrant destroy
````

Reload deployment on config changes.

````shell
vagrant reload
````

## Further Reading

- [Vagrant Docker Provider](https://www.vagrantup.com/docs/providers/docker)
- [Ubuntu in Docker with SSH](https://dev.to/s1ntaxe770r/how-to-setup-ssh-within-a-docker-container-i5i)