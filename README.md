# Specks

Specks is a web application for generating mock data in common data formats.

## Self-hosting

1. Install [docker compose](https://docs.docker.com/compose/).
2. Create a folder `secrets/` and add files `db_pw` and `db_root_pw` containing your MySQL server's passwords.
3. Run `./build.sh` to run the development build.
4. Run `./build.sh --prod` to run the production build. (Does not work yet)

## License

[MIT](https://choosealicense.com/licenses/mit/)
