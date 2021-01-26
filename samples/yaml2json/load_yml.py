import yaml
with open('input.yml') as file:
  obj = yaml.safe_load(file)
  print(obj)
