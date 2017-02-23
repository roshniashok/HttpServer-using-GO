import csv
import time
import json
import sys
import requests
exampleFile = open('dataset.csv')
exampleReader = csv.reader(exampleFile)
for row in exampleReader:
	print('Row #' + str(exampleReader.line_num) + ' ' + str(row))
	time.sleep(1)
	json.dump(row,sys.stdout)
	sys.stdout.write('\n')
	time.sleep(1)
	post_response = requests.post(url='http://localhost:8005/patients/new', data=json.dumps(row))
	time.sleep(1)
