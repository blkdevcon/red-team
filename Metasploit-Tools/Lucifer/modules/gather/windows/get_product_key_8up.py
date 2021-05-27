import winreg as _winreg

from LMI import OS
from LMI.Table import generate_table
from modules.Module import BaseModule


class Module(BaseModule):
    def run(self):
        OS.check_os("windows")
        keys = self.getAllKeys()
        if self.isShellRun:
            print(generate_table([
                ['XP', keys['XP']],
                ['IE', keys['IE']],
                ['WPA', keys['WPA']]
            ],
                title="Product Keys",
                headings=['Type', 'Key']))
            return keys
        return keys

    def set_vars(self):
        return self.default_vars

    def get_description(self):
        self.desc = """This module will retrieve the product key from the windows registry on 8 and up and
 attempts this in three different way: XP, IE and WPA!"""
        return self.desc

    def getAllKeys(self):
        keys = {
            "XP": self.GetXPKey(),
            "IE": self.GetIEKey(),
            "WPA": self.GetWPAKey()
        }
        return keys

    @staticmethod
    def DecodeKey(rpk):
        key = ""
        rpkOffset = 52
        digits = "BCDFGHJKMPQRTVWXY2346789"
        isWin8 = (rpk[66] // 6) & 1
        rpk[66] = (rpk[66] & 0xf7) | ((isWin8 & 2) * 4)

        last = 0
        for i in range(24, -1, -1):
            current = 0
            for j in range(14, -1, -1):
                current *= 256
                current = rpk[j + rpkOffset] + current
                rpk[j + rpkOffset] = current // 24
                current %= 24
                last = current
            key = digits[current] + key
        key = key[1:last + 1] + "N" + key[last + 1:]
        for i in range(5, len(key), 6):
            key = key[:i] + "-" + key[i:]
        return key

    def GetKeyFromRegLoc(self, key, value="DigitalProductID"):
        try:
            key = _winreg.OpenKey(
                _winreg.HKEY_LOCAL_MACHINE, key)

            value, _ = _winreg.QueryValueEx(key, value)
        except Exception as e:
            _ = e  # VOID for IDE niceness
            return "Error"
        return self.DecodeKey(list(value))

    def GetIEKey(self):
        return self.GetKeyFromRegLoc(r"SOFTWARE\Microsoft\Internet Explorer\Registration")

    def GetXPKey(self):
        return self.GetKeyFromRegLoc(r"SOFTWARE\Microsoft\Windows NT\CurrentVersion")

    def GetWPAKey(self):
        return self.GetKeyFromRegLoc(r"SYSTEM\WPA\Key-4F3B2RFXKC9C637882MBM")
