import csv

with open('03_colors.sql', 'w') as sqlfile:
    sqlfile.write(
        ('use specks_db;\n'
         'CREATE TABLE colors (\n'
         '    id INTEGER AUTO_INCREMENT,\n'
         '    name VARCHAR(30),\n'
         '    PRIMARY KEY (id)\n'
         ');\n'))

    with open('wikipedia_color_names.csv', newline='') as csvfile:
        # color name | RGB / HSL Data

        reader = csv.reader(csvfile, delimiter=',', quotechar='"')
        next(reader) # skip the CSV header

        for i, row in enumerate(reader):
            if i >= 1000:
                break
            name = row[0]
            name = name.replace("'", "''")

            if len(name) > 30:
                continue

            sqlfile.write(f'INSERT INTO colors (name) VALUES (\'{name}\');\n')
