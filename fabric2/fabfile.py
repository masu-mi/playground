# coding: utf-8

from fabric2 import Connection, ThreadingGroup, task
from invoke import run

import sys
from tee import StdoutTee, StderrTee

import core
from helper import MultipleContext, pass_checkpoint

@task
def group_lte(c):
    with log_tees():
        tg = ThreadingGroup('test-lte', 'test-enb')
        core.check_hostname(tg)
        for c in tg:
            uname_a(c)

    result = run("cd ../serverspec && bundle exec rake spec:test-lte", pty=True)
    with log_tees():
        print(result.stdout)
        print(result.stderr, file=sys.stderr)
        run("cd ../serverspec && bundle exec rake spec:test-lte")
        if not pass_checkpoint("Do you execute next step?"):
            print("Canceled!!")
            exit(1)
        print("NextStep!!")
        core.uname_a(tg)

@task
def check_hostname(c):
    core.check_hostname(c)

@task
def uname_a(c):
    core.uname_a(c)

def log_tees():
    return MultipleContext(StdoutTee("./log/stdout.log", buff=1),
                 StderrTee("./log/stderr.log", buff=1))

