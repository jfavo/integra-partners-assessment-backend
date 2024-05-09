#!/bin/bash

# Move to sqitch directory
cd ./sqitch

# Configure sqitch settings
sqitch config --user engine.pg.client /usr/bin/psql
sqitch config --user user.name $(echo $(git config --get user.name))
sqitch config --user user.email $(echo $(git config --get user.email))

# Local DB target to test deployment/verify/reverts
sqitch target add $POSTGRES_DB db:pg://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOSTNAME:$POSTGRES_PORT/$POSTGRES_DB
sqitch engine add pg $POSTGRES_DB