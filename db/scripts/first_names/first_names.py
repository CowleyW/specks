import csv

with open('01_first_names.sql', 'w') as sqlfile:
    sqlfile.write(
        ('use specks_db;\n'
         'CREATE TABLE first_names (\n'
         '    id INTEGER AUTO_INCREMENT,\n'
         '    gender TINYINT(1),\n'
         '    name VARCHAR(20),\n'
         '    PRIMARY KEY (id)\n'
         ');\n'))

    with open('baby_names.csv', newline='') as csvfile:
        # Year of Birth | Gender | Ethnicity | First Name | Count | Rank

        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader) # skip the CSV header

        for i, row in enumerate(reader):
            name = row[3]
            name = name[0].upper() + name[1:].lower()
            name = name.replace("'", "''")

            gender = 0 if row[1] == 'MALE' else 1
            sqlfile.write(f'INSERT INTO first_names (gender, name) VALUES ({gender}, \'{name}\');\n')
