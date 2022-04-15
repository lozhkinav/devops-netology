# Домашнее задание к занятию "08.04 Создание собственных modules"

## Основная часть

Наша цель - написать собственный module, который мы можем использовать в своей role, через playbook. Всё это должно быть собрано в виде collection и отправлено в наш репозиторий.

1. - 3. 
```
Файл создан: my_new_module.py, заполнен и отредактирован
```
4. Проверьте module на исполняемость локально.
```
alex@upc:~/devops-projects/myrepo/8_ansible/ansible_rep/ansible $ python -m ansible.modules.my_new_module input.json

{"invocation": {"module_args": {"content": "some data \nmulti line", "path": "/tmp/test.txt"}}, "message": "file was written", "changed": true, "original_message": "some data \nmulti line"}

alex@upc:~/devops-projects/myrepo/8_ansible/ansible_rep/ansible $ cat /tmp/test.txt
some data 
multi line
alex@upc:~/devops-projects/myrepo/8_ansible/ansible_rep/ansible 
```
5. Напишите single task playbook и используйте module в нём.
6. Проверьте через playbook на идемпотентность.
```
alex@upc:~/devops-projects/myrepo/8_ansible/ansible_rep/ansible $ ansible-playbook test_pb.yml 
[WARNING]: You are running the development version of Ansible. You should only run
Ansible from "devel" if you are modifying the Ansible engine, or trying out features
under development. This is a rapidly changing source of code and can become unstable at
any point.
[WARNING]: provided hosts list is empty, only localhost is available. Note that the
implicit localhost does not match 'all'
[WARNING]: ansible.utils.display.initialize_locale has not been called, this may result
in incorrectly calculated text widths that can cause Display to print incorrect line
lengths

PLAY [test my module] *******************************************************************

TASK [Gathering Facts] ******************************************************************
ok: [localhost]

TASK [run my module] ********************************************************************
changed: [localhost]

TASK [dump test_out] ********************************************************************
ok: [localhost] => {
    "msg": {
        "changed": true,
        "failed": false,
        "message": "file was written",
        "original_message": "some new content"
    }
}

PLAY RECAP ******************************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


Проверка на идемпотентность
alex@upc:~/devops-projects/myrepo/8_ansible/ansible_rep/ansible $ ansible-playbook test_pb.yml 
[WARNING]: You are running the development version of Ansible. You should only run Ansible from "devel"
if you are modifying the Ansible engine, or trying out features under development. This is a rapidly
changing source of code and can become unstable at any point.
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost
does not match 'all'
[WARNING]: ansible.utils.display.initialize_locale has not been called, this may result in incorrectly
calculated text widths that can cause Display to print incorrect line lengths

PLAY [test my module] **********************************************************************************

TASK [Gathering Facts] *********************************************************************************
ok: [localhost]

TASK [run my module] ***********************************************************************************
ok: [localhost]

TASK [dump test_out] ***********************************************************************************
ok: [localhost] => {
    "msg": {
        "changed": false,
        "failed": false,
        "message": "file was written",
        "original_message": "some new content"
    }
}

PLAY RECAP *********************************************************************************************
localhost                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```
7. Выйдите из виртуального окружения.
```
alex@upc:~/devops-projects/myrepo/8_ansible/ansible_rep/ansible $  deactivate
```
8. Инициализируйте новую collection: `ansible-galaxy collection init my_own_namespace.yandex_cloud_elk`
9. В данную collection перенесите свой module в соответствующую директорию.
10. Single task playbook преобразуйте в single task role и перенесите в collection. У role должны быть default всех параметров module
11. Создайте playbook для использования этой role.
```
Роль создана, запущена + идемпотентность
alex@upc:~/devops-projects/myrepo/8_ansible/1 $ ansible-galaxy collection init my_netology.my_collection
- Collection my_netology.my_collection was created successfully

alex@upc:~/devops-projects/myrepo/8_ansible/1 $ ansible-playbook site2.yml
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost
does not match 'all'

PLAY [localhost] ***************************************************************************************

TASK [Gathering Facts] *********************************************************************************
ok: [localhost]

TASK [myrole : run my module] **************************************************************************
changed: [localhost]

PLAY RECAP *********************************************************************************************
localhost                  : ok=2    changed=1    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


alex@upc:~/devops-projects/myrepo/8_ansible/1 $ ansible-playbook site2.yml
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost
does not match 'all'

PLAY [localhost] ***************************************************************************************

TASK [Gathering Facts] *********************************************************************************
ok: [localhost]

TASK [myrole : run my module] **************************************************************************
ok: [localhost]

PLAY RECAP *********************************************************************************************
localhost                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   

alex@upc:~/devops-projects/myrepo/8_ansible/1 $ 
```
12. Заполните всю документацию по collection, выложите в свой репозиторий, поставьте тег `1.0.0` на этот коммит.
13. Создайте .tar.gz этой collection: `ansible-galaxy collection build` в корневой директории collection.
```
alex@upc:~/devops-projects/myrepo/8_ansible/my_collection/my_netology/my_collection $ ansible-galaxy collection build
Created collection for my_netology.my_collection at /home/alex/devops-projects/myrepo/8_ansible/my_collection/my_netology/my_collection/my_netology-my_collection-1.0.0.tar.gz
```
14. Создайте ещё одну директорию любого наименования, перенесите туда single task playbook и архив c collection.
15. Установите collection из локального архива: `ansible-galaxy collection install <archivename>.tar.gz`
16. Запустите playbook, убедитесь, что он работает.
```
alex@upc:~/devops-projects/myrepo/8_ansible/1 $ ansible-playbook site.yml
[WARNING]: provided hosts list is empty, only localhost is available. Note that the implicit localhost
does not match 'all'

PLAY [test my module] **********************************************************************************

TASK [Gathering Facts] *********************************************************************************
ok: [localhost]

TASK [run my module] ***********************************************************************************
ok: [localhost]

TASK [dump test_out] ***********************************************************************************
ok: [localhost] => {
    "msg": {
        "changed": false,
        "failed": false,
        "message": "file exists",
        "original_message": "some new content"
    }
}

PLAY RECAP *********************************************************************************************
localhost                  : ok=3    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   


alex@upc:~/devops-projects/myrepo/8_ansible/1 $ tree
.
├── collections
│   └── ansible_collections
│       └── my_netology
│           └── my_collection
│               ├── docs
│               ├── FILES.json
│               ├── MANIFEST.json
│               ├── plugins
│               │   ├── modules
│               │   │   └── my_new_module.py
│               │   └── README.md
│               ├── README.md
│               └── roles
│                   ├── defaults
│                   │   └── main.yml
│                   ├── handlers
│                   │   └── main.yml
│                   ├── meta
│                   │   └── main.yml
│                   ├── tasks
│                   │   └── main.yml
│                   └── vars
│                       └── main.yml
├── group_vars
│   └── all
│       └── vars.yml
├── inventory
│   └── prod.yml
├── my_netology-my_collection-1.0.0.tar.gz
├── site2.yml
└── site.yml

alex@upc:~/devops-projects/myrepo/8_ansible/1 $ cat site.yml 
---
  - name: test my module
    hosts: localhost
    tasks:
    - name  : run my module
      my_netology.my_collection.my_new_module:
        path: "/tmp/file.txt"
        content: "some new content"
      register: test_out
    - name: dump test_out
      debug:
        msg: "{{ test_out }}"  
```
17. В ответ необходимо прислать ссылку на репозиторий с collection
```
https://github.com/lozhkinav/my_collection
```