set -ex
cat > /etc/postgresql/9.4/main/pg_hba.conf <<EOF
local   all             postgres                                peer

# TYPE  DATABASE        USER            ADDRESS                 METHOD

# "local" is for Unix domain socket connections only
local   all             all                                     peer
# IPv4 local connections:
host    all             all             127.0.0.1/32            trust
# IPv6 local connections:
host    all             all             ::1/128                 trust
EOF


service postgresql start
su postgres -c 'createuser -s root'
createdb vidos
