import os
import zipfile
import requests

TARGET_DIR = '../raw_data'

JESTER_URL = 'https://eigentaste.berkeley.edu/dataset/'
JESTER_TEXTS = 'jester_dataset_1_joke_texts.zip'
JESTER_RATINGS_1 = 'jester_dataset_1_1.zip'
JESTER_RATINGS_2 = 'jester_dataset_1_2.zip'
JESTER_RATINGS_3 = 'jester_dataset_1_3.zip'


def fetch_jester_data():
    """
    Fetch Jester dataset and save it to the target directory.
    """
    if not os.path.exists(TARGET_DIR):
        os.makedirs(TARGET_DIR)

    # Fetch joke texts
    response = requests.get(JESTER_URL + JESTER_TEXTS)
    with open(os.path.join(TARGET_DIR, JESTER_TEXTS), 'wb') as f:
        f.write(response.content)

    # Fetch ratings
    for i in range(1, 4):
        response = requests.get(JESTER_URL + f'jester_dataset_1_{i}.zip')
        with open(os.path.join(TARGET_DIR, f'jester_dataset_1_{i}.zip'), 'wb') as f:
            f.write(response.content)

    # Unzip files
    with zipfile.ZipFile(os.path.join(TARGET_DIR, JESTER_TEXTS), 'r') as zip_ref:
        zip_ref.extractall(TARGET_DIR)

    for i in range(1, 4):
        with zipfile.ZipFile(os.path.join(TARGET_DIR, f'jester_dataset_1_{i}.zip'), 'r') as zip_ref:
            zip_ref.extractall(TARGET_DIR)

if __name__ == '__main__':
    print('Fetching Jester dataset...')
    fetch_jester_data()