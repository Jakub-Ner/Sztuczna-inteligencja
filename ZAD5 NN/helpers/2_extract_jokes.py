import os 
import re

RAW_JOKES_DIR = '../raw_data/jokes/'
OUTPUT_JOKES_DIR = '../data/'
JOKES_SEP = '\n\n\n'

def get_all_filenames():
  return [f for f in os.listdir(RAW_JOKES_DIR) if f.endswith('.html')]

def extract_jokes():
  BEGIN_JOKE = '<!--begin of joke -->\n'
  END_JOKE = '<!--end of joke -->'
  jokes = []

  for filename in get_all_filenames():

    with open(RAW_JOKES_DIR + filename, 'r', encoding='utf-8', errors='ignore') as f:
      text = f.read()

    start = text.find(BEGIN_JOKE)
    if start == -1:
      raise Exception('Joke not found')
    
    end = text.find(END_JOKE)
    joke = text[start + len(BEGIN_JOKE):end]

    clean = re.compile('<.*?>')
    joke = re.sub(clean, '', joke)
    joke = joke.replace('\n\n\n', '\n') # just to make sure 
    jokes.append(joke)
  
  print('Jokes extracted successfully')

  if not os.path.exists(OUTPUT_JOKES_DIR):
    os.makedirs(OUTPUT_JOKES_DIR)

  with open(OUTPUT_JOKES_DIR + 'jokes.txt', 'w') as f:
    for joke in jokes:
      f.write(joke + JOKES_SEP)

  print(len(jokes))
      



if __name__ == '__main__':
  extract_jokes()
  