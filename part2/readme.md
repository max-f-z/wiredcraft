## Part 2

Make sure you have vagrant installed and centos/7 box ready.

run
```
vagrant up
```

this should start a centos/7 vagrant box and automatically provision the new vm using

the provided playbook & role.

The steps is really strightforward and intuitive. 

- Fisrt disable selinux & install prerequisites.

- Install & config linux

- Install & config redis

- Install & config app

- resart services

- setup crontab to backup redis