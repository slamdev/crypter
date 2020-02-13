# crypter

CLI to encrypt secrets in config files, so they can be stored safely in git.

## Encrypt

```shell script
DEC_SECRET='very-secret-password'
PK=$(cat public-key.pem)
docker run --rm slamdev/crypter crypter encrypt -s "${PK}" -v "${DEC_SECRET}"
```
outputs: `{cipher}base64-data...=={cipher}`

This can be placed in a config file:
```yaml
pg_host: localhost
pg_user: postgres
pg_password: '{cipher}base64-data...=={cipher}'
```

## Decrypt

```shell script
ENC_SECRET=$(cat config.yaml)
PK=$(cat private-key.pem)
docker run --rm slamdev/crypter crypter decrypt -s "${PK}" -v "${ENC_SECRET}"
```
outputs:
```yaml
pg_host: localhost
pg_user: postgres
pg_password: 'very-secret-password'
```
