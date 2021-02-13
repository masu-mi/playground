# coding: utf-8

class MultipleContext():
    def __init__(self, *args):
        self.opts = args
    def __enter__(self):
        for opt in self.opts:
            opt.__enter__()
    def __exit__(self, *args):
        for opt in self.opts:
            opt.__exit__(args)


import re

__yes = re.compile('[y](es?)?', flags=re.IGNORECASE)

def pass_checkpoint(msg):
    choice = input("{} [y/N]: ".format(msg))
    print(choice)
    if __yes.fullmatch(choice):
        return True
    return False
