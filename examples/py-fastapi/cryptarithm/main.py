# vim: fileencoding=utf-8

import fire
import uvicorn

from fastapi import FastAPI

from parse import Parser
from lexer import Token
from problem import Problem

app = FastAPI()

@app.get("/")
async def read_root():
    return {"accessible_path": ["/docs", "/cryptarithm"]}

@app.get("/cryptarithm/")
async def solve_cryptarithm(expr: str = None):
    p = Parser(Token.tokenize(expr))
    pr = Problem(p.parse())
    _, solutions = pr.search_all_solution()
    return {"problem": expr, "answer": solutions}

class Cli(object):
    def __init__(self, offset=1):
        self._offset = offset

    def server(self, port=8000):
        """start `Hello World` server"""
        uvicorn.run(app, host='0.0.0.0', port=port)

    def cryptarithm(self, expr = ""):
        """solve cryptarithm problem"""
        print("Problem: {}".format(expr))

        p = Parser(Token.tokenize(expr))
        pr = Problem(p.parse())
        print(pr.search_all_solution())

if __name__ == '__main__':
    fire.Fire(Cli)
