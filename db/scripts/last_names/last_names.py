import csv

with open('02_last_names.sql', 'w') as sqlfile:
    sqlfile.write(
        ('use specks_db;\n'
         'CREATE TABLE last_names (\n'
         '    id INTEGER AUTO_INCREMENT,\n'
         '    name VARCHAR(20),\n'
         '    PRIMARY KEY (id)\n'
         ');\n'))

    with open('last_names.csv', newline='') as csvfile:
        # name | rank | count | prop100k | cum_prop100k | pct[...]

        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader) # skip the CSV header

        for i, row in enumerate(reader):
            if i >= 1000:
                break
            name = row[0]
            name = name[0].upper() + name[1:].lower()
            name = name.replace("'", "''")

            sqlfile.write(f'INSERT INTO last_names (name) VALUES (\'{name}\');\n')
