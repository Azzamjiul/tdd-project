import unittest
import functools
import operator

class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency

    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)

    def divide(self, multiplier):
        return Money(self.amount / multiplier, self.currency)

    def __eq__(self, other):
        return self.amount == other.amount and self.currency == other.currency

class Portofolio:
    def __init__(self):
        self.moneys = []

    def add(self, *moneys):
        self.moneys.extend(moneys)
    
    def evaluate(self, currency):
        total = functools.reduce(operator.add, map(lambda m : m.amount, self.moneys), 0)
        return Money(total, currency)

class TestMoney(unittest.TestCase):
    def testMultiplication(self):
        fiveDollars = Money(5, "USD")
        tenDollars = Money(10, "USD")
        self.assertEqual(fiveDollars.times(2), tenDollars)

    def testMultiplicationInEuros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = tenEuros.times(2)
        twentyEuros.__eq__(tenEuros.times(2))

    def testDivision(self):
        originalMoney = Money(4002, "KRW")
        expectedMoneyAfterDivision = Money(1000.5, "KRW")
        expectedMoneyAfterDivision.__eq__(originalMoney.divide(4))

    def testAddition(self):
        fiveDollars = Money(5, "USD")
        tenDollars = Money(10, "USD")
        fifteenDollars = Money(15, "USD")
        portofolio = Portofolio()
        portofolio.add(fiveDollars, tenDollars)
        self.assertEqual(fifteenDollars, portofolio.evaluate("USD"))

if __name__ == '__main__':
    unittest.main()