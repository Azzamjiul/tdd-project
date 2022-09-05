package stocks

type Portofolio []Money

func (p Portofolio) Add(money Money) Portofolio {
	p = append(p, money)
	return p
}

func (p Portofolio) Evaluate(currency string) Money {
	total := 0.0

	for _, m := range p {
		total = total + m.amount
	}

	return Money{amount: total, currency: "USD"}
}
