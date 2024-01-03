import subprocess
import os
import shutil

scripts = [
    ('scripts/first_names/', 'first_names.py', 'scripts/first_names/01_first_names.sql', '01_first_names.sql'),
    ('scripts/last_names/', 'last_names.py', 'scripts/last_names/02_last_names.sql', '02_last_names.sql')
]

for script in scripts:
    dir = os.path.dirname(os.path.abspath(script[0] + script[1]))
    process = subprocess.Popen(['python3', script[1]], cwd=dir, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    stdout, stderr = process.communicate()

    if stdout:
        print(f"{script[0]} stdout: %s" % (stdout.decode()))
    if stderr:
        print(f"{script[0]} stderr: %s" % (stderr.decode()))

    shutil.move(script[2], script[3])
