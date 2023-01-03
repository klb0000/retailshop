import csv
import json


rows = []
brands = dict()
with open('product_data.csv') as fr:
    csv_r =csv.DictReader(fr)
    for row in csv_r:
        pid = row['uniq_id']
        brand = row['brand']
        if len(brand) == 0: continue
        brands[pid] = brand
        
with open('b.json', 'w') as fw:
        json.dump(brands, fw)
#     fw.writelines('uniq_id,brand\n')
#     for row in rows:
#         fw.writelines(row.strip(' ')+ '\n')

# data = json.(brands)
# print(data)

for k in brands:
    print(k, brands[k])
# print(len(brands))