import ipaddress
import redis
from tqdm import tqdm

# 设置允许的IP池
ALLOW_IPS = ['127.0.0.1']
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("114.214.160.0/19").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("114.214.192.0/18").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("202.38.64.0/19").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("210.45.64.0/20").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("210.45.112.0/20").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("211.86.144.0/20").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("222.195.64.0/19").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("210.72.22.0/24").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("202.141.160.0/20").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("218.22.21.0/27").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("218.104.71.160/28").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("218.104.71.96/28").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("202.141.176.0/20").hosts()]
ALLOW_IPS += [str(x) for x in ipaddress.ip_network("121.255.0.0/16").hosts()]

for i in range(8):
    ALLOW_IPS.append("202.111.192." + str(int(24+i)))

for i in range(32):
    ALLOW_IPS.append("218.22.22." + str(int(160+i)))

ALLOW_IPS.append("112.32.50.146")

#  写入redis
redis_conn = redis.Redis(host="192.168.10.10", port="6379")
for index in tqdm(range(len(ALLOW_IPS))):
    redis_conn.set(ALLOW_IPS[index], "true")

print("success")