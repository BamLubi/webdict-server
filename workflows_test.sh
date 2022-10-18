#!/bin/bash

mysql -h127.0.0.1 -P3306 -uroot -pustc123123 -e "CREATE DATABASE webdict DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;"
