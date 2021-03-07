import unittest

pool = "(){}[]"

def vop(st) :
    tmp = ""
    # перебераем строку
    for s in st:
        # проверяем является ли символ скобкой
        if s in pool:
            if s == "(":
                tmp = ")" + tmp
            elif s == "{":
                tmp = "}" + tmp
            elif s == "[":
                tmp = "]" + tmp
            else:
                # проверяем закрывается ли последняя открытая скобка
                if s == tmp[0] :
                    tmp = tmp[1:]
                else:
                    return False
    return True

class TestVOP(unittest.TestCase):
    def test_vop(self):
        testList = {"{[]}":True, "{[(]}":False, "{()[]}":True}
        
        for key, value in testList.items():
            self.assertEqual(value, vop(key))
        return

if __name__ == "__main__":
    unittest.main()

