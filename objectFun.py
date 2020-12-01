class Person:
    def __init__(self):
        self.__name = "hhhhh"  #私有属性

    def getInfo(self):
        print("this is myInfo!")

    @classmethod
    def getName(cls):
        #返回私有属性
        #return self.__name
        pass

    #私有方法
    def __updateName(self, newValue):
        self.__name = newValue

    def updateNames(self, newValue):
        #修改私有属性
        self.__name = newValue

class B(Person):
    pass

one = Person()
print('修改之前', one.getName())
#one.__name无法进行调用
#one.__updateName()无法进行调用

Person.getName()
#two.__name__="aaaa"
#print("应该是增加属性了", two.__name__)

two = B()
#B不能继承Person的私有方法
#two.__updateName()也不能调用
#two.__name也不能调用
two.updateNames("wwwwww")
print("子类继承后修改之后", two.getName())
#C:\Users\jierui303\AppData\Local\Programs\Python\Python36

#print(one.colorName)
#one.colorName = "I's color"
#print(one.colorName)
#one.newName = "hello world"
#print(one.newName)