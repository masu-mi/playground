def check_hostname(c):
    c.run("hostname")

def uname_a(c):
    vars(c)
    c.run("uname -a")
