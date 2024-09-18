#!/bin/bash

# Record trade
./minifab invoke -o bank -n playground-hlf -p '"RecordTrade","TRD1","T","5","25.50","2024-09-18T17:55:05.000Z","PROCESSED"'

# Record payment
./minifab invoke -o investment -n playground-hlf -p '"RecordPayment","PAY1","User1","User2","10.05","2024-09-18T17:57:42.000Z","PROCESSED"'

# Record settlement
./minifab invoke -o clearing -n playground-hlf -p '"HandleSettlement","SETL1","TRD1","PAY1","10.05","2024-09-18T17:57:42.000Z","PROCESSED"'

# Query Trade
./minifab invoke -o bank -n playground-hlf -p '"QueryTrades","TRD1"'

# Query Payment
./minifab invoke -o investment -n playground-hlf -p '"QueryPayments","PAY1"'

# Query Settlement
./minifab invoke -o clearing -n playground-hlf -p '"QuerySettlements","SETL1"'
