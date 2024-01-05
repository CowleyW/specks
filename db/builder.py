import subprocess
import os
import shutil

scripts = [
    # 0                       1                 2
    # dir path                script name       SQL name
    ('scripts/first_names/', 'first_names.py', '01_first_names.sql'),
    ('scripts/last_names/',  'last_names.py',  '02_last_names.sql'),
    ('scripts/colors/',       'colors.py',      '03_colors.sql')
]

for script in scripts:
    dir = os.path.dirname(os.path.abspath(script[0] + script[1]))
    process = subprocess.Popen(['python3', script[1]], cwd=dir, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = process.communicate()

    if stdout:
        print(f"{script[0]} stdout: %s" % (stdout.decode()))
    if stderr:
        print(f"{script[0]} stderr: %s" % (stderr.decode()))

    shutil.move(script[0] + script[2], script[2])
