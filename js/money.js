class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency;
    }

    times(multiplier){
        return new Money( this.amount * multiplier, this.currency );
    }

    divide(multiplier){
        return new Money( this.amount / multiplier, this.currency );
    }
}

module.exports = Money;