import os
import sys

cmd = sys.argv[1]
bash_command = ["cd " + cmd, "git status"]
result_os = os.popen(' && '.join(bash_command)).read()
# is_change = False
for result in result_os.split('\n'):
    if result.find('изменено') != -1:
        prepare_result = result.replace('\tизменено: ', '')
        # добавил замену всех оставшихся пробелов в строке для удобства вывода
        prepare_result = prepare_result.replace(' ', '')
        print(cmd + prepare_result)
#        break
