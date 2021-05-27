import os
from subprocess import check_output

from lucifer.Errors import IncompatibleSystemError
from modules.Module import BaseModule


def get_uname(arg):
    arg.insert(0, "uname")
    return check_output(arg).decode()


class Module(BaseModule):
    def run(self):
        args = ["-a"]
        if "nt" in os.name.lower():
            raise IncompatibleSystemError("Not Unix")
        if "args" in self.shell.vars.keys():
            args = self.shell.vars["args"].split(" ")
        if self.isShellRun:
            print(get_uname(args))
        else:
            return get_uname(args)

    def set_vars(self):
        default_vars = {
            "args": "-a"
        }
        return default_vars

    def get_description(self):
        desc = """Gets the output of uname on a unix system with any arguments supplied in the args variable"""
        return desc
